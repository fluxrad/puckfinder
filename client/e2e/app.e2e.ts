import { PuckfinderPage } from './app.po';

describe('puckfinder App', function() {
  let page: PuckfinderPage;

  beforeEach(() => {
    page = new PuckfinderPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('puckfinder works!');
  });
});
