import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject, tap } from 'rxjs';
import { ClosedPosition, Position } from './position.model';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class PositionService {

  constructor(private readonly httpClient: HttpClient) { }

  private openPositions: Position[] = [];
  private closedPositions: ClosedPosition[] = [];

  private openPositionsSubject= new Subject<Position[]>();
  private closedPositionsSubject= new Subject<ClosedPosition[]>();

  getOpenPositionsObservable(): Observable<Position[]> {
    return this.openPositionsSubject.asObservable();
  }

  getClosedPositionsObservable(): Observable<ClosedPosition[]> {
    return this.closedPositionsSubject.asObservable();
  }

  getOpenPositions(): Observable<Position[]> {
    return this.httpClient.get<Position[]>(`${environment.apiUrl}/api/positions`)
      .pipe(
        tap((p: Position[]) => {
          this. openPositions = p;
          this.openPositionsSubject.next(this.openPositions);
        })
      );
  }

  getClosedPositions(): Observable<ClosedPosition[]> {
    return this.httpClient.get<ClosedPosition[]>(`${environment.apiUrl}/api/positions/closed`)
      .pipe(
        tap((cp: ClosedPosition[]) => {
          this.closedPositions = cp;
          this.closedPositionsSubject.next(this.closedPositions);
        })
      );
  }
  closePosition(id: number): Observable<ClosedPosition> {
    return this.httpClient.post<ClosedPosition>(`${environment.apiUrl}/api/position/${id}/close`, {})
      .pipe(
        tap((cp: ClosedPosition) => {
          this.openPositions = this.openPositions.filter(p => p.id !== cp.id)
          this.openPositionsSubject.next(this.openPositions);
          this.closedPositions === null ? this.closedPositions = [cp] : this.closedPositions.push(cp)
          this.closedPositionsSubject.next(this.closedPositions);
        })
      );
  }
}
