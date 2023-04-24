import { Component, OnInit } from '@angular/core';
import { AppService } from '../app.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.scss']
})
export class HomePageComponent implements OnInit {
  status = ''
  timestamp = -1

  constructor(
    private readonly appService: AppService,
  ) { }

  ngOnInit() {
    this.appService.GetHealth().subscribe((res) => {
      this.status = res.status
      this.timestamp = res.timestamp
    })
  }
}
