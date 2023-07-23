import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';
import { Observable } from 'rxjs';
import { Portfolio } from './portfolio.types';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class PortfolioService {

  constructor(private readonly http: HttpClient) { }

  getPortfolioById(id: number): Observable<Portfolio> {
    return this.http.get<Portfolio>(`${environment.apiUrl}/api/portfolio/${id}`);
  }
}
