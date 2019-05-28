import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs/Observable";
import {AuctionCountScore} from "./auction-count-score";
import {environment} from "../environments/environment";

@Injectable()
export class RankingService {
  apiUrl = environment.apiHost;

  constructor(private http: HttpClient) { }
  public fetchActiveAuctionRanking() {
    const apiPath = "auction-count-scores";
    return this.http.get<AuctionCountScore[]>(this.apiUrl + apiPath);
  }
}
