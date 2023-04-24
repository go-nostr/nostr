import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RemoveRelayFormComponent } from './remove-relay-form.component';

describe('RemoveRelayFormComponent', () => {
  let component: RemoveRelayFormComponent;
  let fixture: ComponentFixture<RemoveRelayFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RemoveRelayFormComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RemoveRelayFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
