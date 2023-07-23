import { Component, OnDestroy, OnInit } from '@angular/core';
import { ClosedPosition, Position } from './position.model';
import { PositionService } from './position.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-position',
  templateUrl: './position.component.html',
})
export class PositionComponent implements OnInit, OnDestroy {

  positions: Position[] = [];
  closedPositions: Map<number, ClosedPosition[]> = new Map();
  private positionSubscription: Subscription | undefined;
  private closedPositionSubscription: Subscription | undefined;

  constructor(private readonly positionService: PositionService) {}

  ngOnInit(): void {
    this.positionSubscription = this.positionService.getOpenPositionsObservable()
      .subscribe(ps => this.positions = ps);
    this.closedPositionSubscription = this.positionService.getClosedPositionsObservable()
      .subscribe(portfolio => this.closedPositions = portfolio);
    this.positionService.getOpenPositions().subscribe();
    this.positionService.getClosedPositions(1).subscribe();
  }

  ngOnDestroy(): void {
    this.positionSubscription?.unsubscribe();
    this.closedPositionSubscription?.unsubscribe();
  }

  closePositionHandler(id: number) {
    this.positionService.closePosition(id).subscribe();
  }
}
