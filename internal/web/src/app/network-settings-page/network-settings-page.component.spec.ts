import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NetworkSettingsPageComponent } from './network-settings-page.component';

describe('NetworkSettingsPageComponent', () => {
  let component: NetworkSettingsPageComponent;
  let fixture: ComponentFixture<NetworkSettingsPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NetworkSettingsPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NetworkSettingsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
