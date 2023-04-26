import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

const defaultHostname = '0.0.0.0'
const defaultPort = '4317'

export interface GetHealthResponse {
  status: string
  timestamp: number
}

export interface GetInternetIdentifier {
  names: string[]
  relays: string[]
}

@Injectable({
  providedIn: 'root'
})
export class AppService {

  constructor(
    private readonly httpClient: HttpClient,
  ) { }
}
