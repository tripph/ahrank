import { Component, OnInit } from '@angular/core';
import {RankingService} from "../../ranking.service";
import {AuctionCountScore} from "../../auction-count-score";

@Component({
  selector: 'ahr-activeauctions',
  templateUrl: './activeauctions.component.html',
  styleUrls: ['./activeauctions.component.scss']
})
export class ActiveauctionsComponent implements OnInit {
  scores: AuctionCountScore[];
  constructor(private rankingService: RankingService) { }



  ngOnInit() {
    this.rankingService.fetchActiveAuctionRanking()
      .subscribe(results => {
        this.scores = results;
      }, err => {
        console.log('error fetching ranking: ' + err);
      })
  }

}

