import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';

import { AppComponent } from './app.component';
import {MaterialModule} from "./material/material.module";
import { HomeComponent } from './home/home.component';
import { TopnavComponent } from './topnav/topnav.component';
import { AboutComponent } from './about/about.component';
import { RankingComponent } from './ranking/ranking.component';
import { ActiveauctionsComponent } from './ranking/activeauctions/activeauctions.component';
import { RankingService } from './ranking.service';
import {HttpClient, HttpClientModule} from "@angular/common/http";


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    TopnavComponent,
    AboutComponent,
    RankingComponent,
    ActiveauctionsComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    MaterialModule
  ],
  providers: [RankingService, HttpClient],
  bootstrap: [AppComponent]
})
export class AppModule { }
