import { Component, OnInit } from '@angular/core';
import { PortfolioService } from './portfolio.service';
import { Portfolio } from './portfolio.types';

@Component({
  selector: 'app-portfolio',
  templateUrl: './portfolio.component.html',
  styleUrls: ['./portfolio.component.css']
})
export class PortfolioComponent implements OnInit {

  public portfolio: Portfolio | undefined;
  constructor(private readonly portfolioService: PortfolioService) {}

  ngOnInit(): void {
    this.portfolioService.getPortfolioById(1)
      .subscribe(p => this.portfolio = p);
  }
}
