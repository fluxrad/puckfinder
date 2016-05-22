export class PuckfinderPage {
  navigateTo() {
    return browser.get('/');
  }

  getParagraphText() {
    return element(by.css('puckfinder-app h1')).getText();
  }
}
