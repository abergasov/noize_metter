import puppeteer from 'puppeteer';
import dotenv from 'dotenv'
import fs from 'fs';

const mainURL = 'https://192.168.1.12:8081/webvisu/webvisu.htm?CFG_Lang=en&BRLG&NOSP&RLLG';

global.taskID = 0;
global.cpatchaReslove = "";

(async () => {
    dotenv.config()

    const browser = await puppeteer.launch({
        headless: false
    });
    const user = process.env.USER
    const pass = process.env.PASS
    const page = await browser.newPage();
    await page.goto(mainURL);
    //await browser.close()
})();

