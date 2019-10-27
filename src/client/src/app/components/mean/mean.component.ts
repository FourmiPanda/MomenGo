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
  constructor(private http: HttpClient) { }

  ngOnInit() {
  }
  getMean() {
    console.log(this.meanDate);
    console.log(this.meanUrl + this.meanDate);
    
    //let means = this.http.get(this.meanUrl + this.meanDate);
  }
}

