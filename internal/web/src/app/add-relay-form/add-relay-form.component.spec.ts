import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddRelayFormComponent } from './add-relay-form.component';

describe('AddRelayFormComponent', () => {
  let component: AddRelayFormComponent;
  let fixture: ComponentFixture<AddRelayFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddRelayFormComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddRelayFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
