import { Component, OnInit } from '@angular/core';
import { Position } from './position.model';
import { PositionService } from './position.service';

@Component({
  selector: 'app-position',
  templateUrl: './position.component.html',
})
export class PositionComponent implements OnInit {

  positions: Position[] = [];

  constructor(private readonly positionService: PositionService) {}

  ngOnInit(): void {
    this.positionService.getPositions()
      .subscribe(result => this.positions = result);
  }

  closePositionHandler(id: number) {
    this.positionService.closePosition(id)
      .subscribe(res => console.log(res))
  }

}
