import { Component, OnInit } from '@angular/core';
import {Realm} from "./realm";
import {RealmService} from "./realm.service";

@Component({
  selector: 'ahr-realms',
  templateUrl: './realms.component.html',
  styleUrls: ['./realms.component.scss']
})
export class RealmsComponent implements OnInit {
  realmList: Realm[];
  constructor(private realmService: RealmService) { }

  ngOnInit() {
    this.realmService.getRealmList()
      .subscribe(resp => {
        this.realmList = resp;
      }, err => {
        alert(err);
      })
  }

}
