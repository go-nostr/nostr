import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

import { AppService, GetHealthResponse } from './app.service';

describe('AppService', () => {
  let service: AppService;
  let httpTestingController: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [AppService]
    });
    service = TestBed.inject(AppService);
    httpTestingController = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpTestingController.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
  
  describe('AppService.prototype.GetHealth', () => { 
    it('should get health', () => {
      const mockResponse: GetHealthResponse = {
        status: 'ok',
        timestamp: Date.now()
      };

      service.GetHealth().subscribe(response => {
        expect(response).toBeTruthy();
        expect(response).toEqual(mockResponse);
      });

      const req = httpTestingController.expectOne('http://0.0.0.0:4317/health');
      expect(req.request.method).toEqual('GET');
      req.flush(mockResponse);
    });
  });
});
