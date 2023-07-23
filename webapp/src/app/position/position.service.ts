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

  private positions: Position[] = [];
  private positionSubject = new Subject<Position[]>();

  getPositionObservable(): Observable<Position[]> {
    return this.positionSubject.asObservable();
  }

  getPositions(): Observable<Position[]> {
    return this.httpClient.get<Position[]>(`${environment.apiUrl}/api/positions`)
      .pipe(
        tap((p: Position[]) => {
          this.positions = p;
          this.positionSubject.next(this.positions);
          console.log(p);
        })
      );
  }

  closePosition(id: number): Observable<ClosedPosition> {
    return this.httpClient.post<ClosedPosition>(`${environment.apiUrl}/api/position/${id}/close`, {})
      .pipe(
        tap((cp: ClosedPosition) => {
          this.positions = this.positions.filter(p => p.id !== cp.id)
          this.positionSubject.next(this.positions);
        })
      );
  }
}
