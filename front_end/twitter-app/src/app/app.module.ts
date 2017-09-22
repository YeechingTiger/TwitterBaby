import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppComponent } from './app.component';
import { NavBarComponent } from './components/nav-bar/nav-bar.component';

import { DataService } from './services/data.service';

import { rooting } from './app.router';

import { MdToolbarModule, MdInputModule, MdMenuModule, 
         MdIconModule, MdButtonModule, MdCardModule } from '@angular/material';
import { UserInfoComponent } from './components/user-info/user-info.component';
import { TweetlistComponent } from './components/tweetlist/tweetlist.component';
import { PostareaComponent } from './components/postarea/postarea.component';
import { LoginComponent } from './components/login/login.component';
import { SignupComponent } from './components/signup/signup.component';
@NgModule({
  declarations: [
    AppComponent,
    NavBarComponent,
    UserInfoComponent,
    TweetlistComponent,
    PostareaComponent,
    LoginComponent,
    SignupComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    MdToolbarModule,
    MdInputModule,
    MdMenuModule,
    MdIconModule,
    MdButtonModule,
    MdCardModule,
    rooting
  ],
  providers: [{provide: 'data',
              useClass: DataService}],
  bootstrap: [AppComponent]
})
export class AppModule { }
