import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { HttpClientModule } from '@angular/common/http';
import { AppComponent } from './app.component';
import { HomePageComponent } from './home-page/home-page.component';
import { ProfilePageComponent } from './profile-page/profile-page.component';
import { AppearanceSettingsPageComponent } from './appearance-settings-page/appearance-settings-page.component';
import { AccountSettingsPageComponent } from './account-settings-page/account-settings-page.component';
import { BackupSettingsPageComponent } from './backup-settings-page/backup-settings-page.component';
import { CreateShortTextNoteEventFormComponent } from './create-short-text-note-event-form/create-short-text-note-event-form.component';
import { NotificationsPageComponent } from './notifications-page/notifications-page.component';
import { NetworkSettingsPageComponent } from './network-settings-page/network-settings-page.component';
import { SignInPageComponent } from './sign-in-page/sign-in-page.component';
import { SignUpPageComponent } from './sign-up-page/sign-up-page.component';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { NotFoundPageComponent } from './not-found-page/not-found-page.component';
import { AddRelayFormComponent } from './add-relay-form/add-relay-form.component';
import { ListRelaysComponent } from './list-relays/list-relays.component';
import { RemoveRelayFormComponent } from './remove-relay-form/remove-relay-form.component';
import { ListEventsComponent } from './list-events/list-events.component';
import { ListNotificationsComponent } from './list-notifications/list-notifications.component';

@NgModule({
  declarations: [
    AppComponent,
    HomePageComponent,
    ProfilePageComponent,
    AppearanceSettingsPageComponent,
    AccountSettingsPageComponent,
    BackupSettingsPageComponent,
    CreateShortTextNoteEventFormComponent,
    NotificationsPageComponent,
    NetworkSettingsPageComponent,
    SignInPageComponent,
    SignUpPageComponent,
    LandingPageComponent,
    NotFoundPageComponent,
    AddRelayFormComponent,
    ListRelaysComponent,
    RemoveRelayFormComponent,
    ListEventsComponent,
    ListNotificationsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
  ],
  providers: [

  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
