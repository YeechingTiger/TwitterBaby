import { Injectable } from '@angular/core';
import { Tweet } from '../models/tweet.model';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import 'rxjs/add/operator/toPromise';
@Injectable()
export class DataService {

  constructor(private http: Http) { }
  getTweetList(id: string): Promise<Tweet[]> {
    return this.http.get(`api/v1/tweetlist/${id}`)
                      .toPromise()
                      .then((res: Response) => res.json())
                      .catch(this.handleError);
  }

  // ERROR handler
  private handleError(error: any): Promise<any> {
    console.error('An error occurred', error);
    return Promise.reject(error.body || error);
  }
}
