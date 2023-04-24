import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListNotificationsComponent } from './list-notifications.component';

describe('ListNotificationsComponent', () => {
  let component: ListNotificationsComponent;
  let fixture: ComponentFixture<ListNotificationsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListNotificationsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ListNotificationsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
