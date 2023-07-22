import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, tap } from 'rxjs';
import { ClosedPosition, Position } from './position.model';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class PositionService {

  constructor(private readonly httpClient: HttpClient) { }

  private positions: Position[] = [];

  getPositions(): Observable<Position[]> {
    return this.httpClient.get<Position[]>(`${environment.apiUrl}/api/positions`)
      .pipe(
        tap((positions: Position[]) => {
          this.positions = positions;
        })
      );
  }

  closePosition(id: number): Observable<ClosedPosition> {
    return this.httpClient.post<ClosedPosition>(`${environment.apiUrl}/api/position/${id}/close`, {})
      .pipe(
        tap((closedPosition: ClosedPosition) => {
          // Find the position in the array by ID and update its properties with the received data.
          const index = this.positions.findIndex(position => position.id === id);
          if (index !== -1) {
            this.positions[index].closePrice = closedPosition.closePrice;
            this.positions[index].closedAt = closedPosition.closedAt;
            // You can update other properties as needed based on the response from the server.
          }
        })
      );
  }
}
