import { Component, OnDestroy, OnInit } from '@angular/core';
import { Position } from './position.model';
import { PositionService } from './position.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-position',
  templateUrl: './position.component.html',
})
export class PositionComponent implements OnInit, OnDestroy {

  positions: Position[] = [];
  private positionSubscription: Subscription | undefined;

  constructor(private readonly positionService: PositionService) {}

  ngOnInit(): void {
    this.positionSubscription = this.positionService.getPositionObservable()
      .subscribe(ps => {
        this.positions = ps;
      });
    this.positionService.getPositions().subscribe();
  }

  closePositionHandler(id: number) {
    this.positionService.closePosition(id).subscribe();
  }

  ngOnDestroy(): void {
    this.positionSubscription?.unsubscribe();
  }

}
