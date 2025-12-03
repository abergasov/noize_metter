import puppeteer from 'puppeteer';
import dotenv from 'dotenv'
import fs from 'fs';

const mainURL = 'https://192.168.1.12:8081/webvisu/webvisu.htm?CFG_Lang=en&BRLG&NOSP&RLLG';

global.taskID = 0;
global.cpatchaReslove = "";

(async () => {
    dotenv.config()

    const browser = await puppeteer.launch({
        headless: false,
        ignoreHTTPSErrors: true,
        args: [
            '--ignore-certificate-errors',
            '--allow-insecure-localhost'
        ]
    });

    const page = await browser.newPage();
    await page.evaluateOnNewDocument(() => {
        // capture all text drawn on canvas
        const origFillText = CanvasRenderingContext2D.prototype.fillText;
        CanvasRenderingContext2D.prototype.fillText = function (text, x, y, maxWidth) {
            try {
                window.__canvasTexts = window.__canvasTexts || {};
                const canvas = this.canvas;
                if (canvas) {
                    window.__canvasTexts[canvas.id] = String(text);
                }
            } catch (_) {
                // ignore
            }
            return origFillText.call(this, text, x, y, maxWidth);
        };
    });
    await page.goto(mainURL, {
        waitUntil: 'networkidle0',
    });
    await page.waitForSelector('input[placeholder="Username"]');
    const inputs = await page.$$('input');
    const buttons = await page.$$('button');

    // enter credentials
    await inputs[0].type(process.env.SCRAP_USER);
    await inputs[1].type(process.env.SCRAP_PASS);
    await buttons[1].click();

    // wait for appear element with id cdsRoot
    await page.waitForSelector('#cdsRoot', {
        visible: true,
        timeout: 30000
    });
    console.log("authorized")

    const dumpFile = './canvas_dump.log';
    const loop = async () => {
        try {
            const data = await page.evaluate(() => window.__canvasTexts || {});
            const record = {
                ts: new Date().toISOString(),
                data,
            };
            fs.appendFileSync(dumpFile, JSON.stringify(record) + '\n');
            console.log('dumped:', record);
        } catch (err) {
            console.error('dump error:', err);
        } finally {
            setTimeout(loop, 5000);
        }
    };
    loop();
})();

