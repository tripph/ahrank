import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs/Observable";
import {AuctionCountScore} from "./auction-count-score";

@Injectable()
export class RankingService {
  apiUrl = "http://localhost:1232/";

  constructor(private http: HttpClient) { }
  public fetchActiveAuctionRanking() {
    const apiPath = "auction-count-scores";
    return this.http.get<AuctionCountScore[]>(this.apiUrl + apiPath);
  }
}
