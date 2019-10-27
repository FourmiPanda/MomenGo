import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { MeanComponent } from './components/mean/mean.component';
import { CaptorsValuesComponent } from './components/captors-values/captors-values.component';
import {CalendarModule} from 'primeng/calendar';
import { FormsModule } from '@angular/forms';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';
@NgModule({
  declarations: [
    AppComponent,
    MeanComponent,
    CaptorsValuesComponent
  ],
  imports: [
    BrowserModule,
    CalendarModule,
    FormsModule,
    BrowserAnimationsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
