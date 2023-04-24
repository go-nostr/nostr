import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListRelaysComponent } from './list-relays.component';

describe('ListRelaysComponent', () => {
  let component: ListRelaysComponent;
  let fixture: ComponentFixture<ListRelaysComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListRelaysComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ListRelaysComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
