import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BackupSettingsPageComponent } from './backup-settings-page.component';

describe('BackupSettingsPageComponent', () => {
  let component: BackupSettingsPageComponent;
  let fixture: ComponentFixture<BackupSettingsPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BackupSettingsPageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BackupSettingsPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
