import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateShortTextNoteEventFormComponent } from './create-short-text-note-event-form.component';

describe('CreateShortTextNoteEventFormComponent', () => {
  let component: CreateShortTextNoteEventFormComponent;
  let fixture: ComponentFixture<CreateShortTextNoteEventFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateShortTextNoteEventFormComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateShortTextNoteEventFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
