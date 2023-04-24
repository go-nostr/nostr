import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AppearanceSettingsPageComponent } from './appearance-settings-page.component';

describe('AppearanceSettingsPageComponent', () => {
  let component: AppearanceSettingsPageComponent;
  let fixture: ComponentFixture<AppearanceSettingsPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AppearanceSettingsPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AppearanceSettingsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
