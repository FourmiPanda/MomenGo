import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CaptorsValuesComponent } from './captors-values.component';

describe('CaptorsValuesComponent', () => {
  let component: CaptorsValuesComponent;
  let fixture: ComponentFixture<CaptorsValuesComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CaptorsValuesComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CaptorsValuesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
