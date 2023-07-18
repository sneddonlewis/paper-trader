import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Position } from './position.model';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class PositionService {

  constructor(private readonly httpClient: HttpClient) { }

  getPositions(): Observable<Position[]> {
    return this.httpClient.get<Position[]>(`${environment.apiUrl}/api/positions`);
  }
}
