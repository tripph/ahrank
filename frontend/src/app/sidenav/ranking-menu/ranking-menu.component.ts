import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'ahr-ranking-menu',
  templateUrl: './ranking-menu.component.html',
  styleUrls: ['./ranking-menu.component.scss']
})
export class RankingMenuComponent implements OnInit {
  rankingMenuOpen = false;

  constructor() { }

  ngOnInit() {
  }

}
