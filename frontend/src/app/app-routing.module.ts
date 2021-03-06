import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {HomeComponent} from "./home/home.component";
import {AboutComponent} from "./about/about.component";
import {ActiveauctionsComponent} from "./ranking/activeauctions/activeauctions.component";
import {RealmsComponent} from "./realms/realms.component";

const routes: Routes = [
  {
    path: '',
    component: HomeComponent
  },
  {
    path: 'about',
    component: AboutComponent
  },
  {
    path: 'realms',
    component: RealmsComponent
  },
  {
    path: 'ranking/active-auctions',
    component: ActiveauctionsComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
