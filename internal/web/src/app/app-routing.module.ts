import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomePageComponent } from './home-page/home-page.component';
import { ProfilePageComponent } from './profile-page/profile-page.component';
import { NotFoundPageComponent } from './not-found-page/not-found-page.component';
import { AccountSettingsPageComponent } from './account-settings-page/account-settings-page.component';
import { NetworkSettingsPageComponent } from './network-settings-page/network-settings-page.component';
import { NotificationsPageComponent } from './notifications-page/notifications-page.component';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: '/'
  },
  {
    path: '',
    component: HomePageComponent
  },
  {
    path: 'notifications',
    component: NotificationsPageComponent
  },
  {
    path: 'settings/account',
    component: AccountSettingsPageComponent
  },
  {
    path: 'settings/account',
    component: AccountSettingsPageComponent
  },
  {
    path: 'settings/backup',
    component: NetworkSettingsPageComponent
  },
  {
    path: 'settings/network',
    component: NetworkSettingsPageComponent
  },
  {
    path: ':id',
    component: ProfilePageComponent
  },
  {
    path: '**',
    component: NotFoundPageComponent
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
