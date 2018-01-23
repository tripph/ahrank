import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {environment} from "../../environments/environment";
import {Realm} from "./realm";

@Injectable()
export class RealmService {

  constructor(private http: HttpClient) { }
  apiUrl = environment.apiHost;

  getRealmList() {
    return this.http.get<Realm[]>(this.apiUrl + 'api/realms/list');
  }
}
