import { RouterModule, Routes } from '@angular/router';
import { UserInfoComponent } from './components/user-info/user-info.component';
import { LoginComponent } from './components/login/login.component';
import { TweetlistComponent } from './components/tweetlist/tweetlist.component';
import { UserPageComponent } from './components/user-page/user-page.component';

const appRoutes: Routes = [
    { path: 'home', component:  LoginComponent},
    { path: 'user/:id', component: UserPageComponent},
    { path: '**', redirectTo: 'home' }
];
export const rooting = RouterModule.forRoot(appRoutes);
