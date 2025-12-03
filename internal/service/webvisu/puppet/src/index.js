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

    //await browser.close()
})();

