import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AccountSettingsPageComponent } from './account-settings-page.component';

describe('AccountSettingsPageComponent', () => {
  let component: AccountSettingsPageComponent;
  let fixture: ComponentFixture<AccountSettingsPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AccountSettingsPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AccountSettingsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
