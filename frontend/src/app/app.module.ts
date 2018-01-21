import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';
import {MaterialModule} from "./material/material.module";
import { HomeComponent } from './home/home.component';
import { SidenavComponent } from './sidenav/sidenav.component';
import { RankingMenuComponent } from './sidenav/ranking-menu/ranking-menu.component';


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    SidenavComponent,
    RankingMenuComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    MaterialModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
