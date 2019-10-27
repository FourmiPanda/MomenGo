import { Component, OnInit, Input } from '@angular/core';
import { CalendarModule } from 'primeng/calendar';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient } from '@angular/common/http';
@Component({
  selector: 'app-mean',
  templateUrl: './mean.component.html',
  styleUrls: ['./mean.component.css']
})
export class MeanComponent implements OnInit {
  //variables
  @Input() meanDate: string
  meanUrl =" http://localhost:2019/mean?date="
  tempMean = "50"
  windMean="120"
  pressurMean="12"
  constructor(private http: HttpClient) { }

  ngOnInit() {
  }
  getMean() {
    console.log(this.meanUrl + this.meanDate);
    if(this.meanDate==undefined){
      alert("Please enter a date")
    }else{
      //let means = this.http.get(this.meanUrl + this.meanDate);
    }
    
  }
}

