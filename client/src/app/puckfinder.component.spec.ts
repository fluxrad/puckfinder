import {
  beforeEachProviders,
  describe,
  expect,
  it,
  inject
} from '@angular/core/testing';
import { PuckfinderAppComponent } from '../app/puckfinder.component';

beforeEachProviders(() => [PuckfinderAppComponent]);

describe('App: Puckfinder', () => {
  it('should create the app',
      inject([PuckfinderAppComponent], (app: PuckfinderAppComponent) => {
    expect(app).toBeTruthy();
  }));

  it('should have as title \'puckfinder works!\'',
      inject([PuckfinderAppComponent], (app: PuckfinderAppComponent) => {
    expect(app.title).toEqual('puckfinder works!');
  }));
});
