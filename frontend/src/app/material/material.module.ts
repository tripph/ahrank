import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  MatButtonModule, MatCardModule, MatExpansionModule, MatIconModule, MatListModule, MatMenuModule,
  MatSidenavModule, MatToolbarModule
} from "@angular/material";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
@NgModule({
  imports: [
    CommonModule,
    BrowserAnimationsModule,
    MatListModule,
    MatMenuModule,
    MatButtonModule,
    MatIconModule,
    MatSidenavModule,
    MatExpansionModule,
    MatCardModule,
    MatToolbarModule
  ], exports: [
    BrowserAnimationsModule,
    MatListModule,
    MatMenuModule,
    MatIconModule,
    MatButtonModule,
    MatToolbarModule,
    MatSidenavModule,
    MatExpansionModule,
    MatCardModule

  ],
  declarations: []
})
export class MaterialModule { }
