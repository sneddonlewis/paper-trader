import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { TechComponent } from './tech/tech.component';
import { HttpClientModule } from '@angular/common/http';
import { PositionComponent } from './position/position.component';

@NgModule({
  declarations: [
    AppComponent,
    TechComponent,
    PositionComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
