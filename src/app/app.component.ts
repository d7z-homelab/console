import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import {MatSlideToggle} from "@angular/material/slide-toggle";

@Component({
    selector: 'app-root',
  imports: [MatSlideToggle],
    templateUrl: './app.component.html',
    styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'web';
}
