import { Injectable, OnInit } from '@angular/core';
import { EventsOn } from '../../../../wailsjs/runtime';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TerminalService {

  outputReceived$ = new Subject<string>();
  logReceived$ = new Subject<string>();


  constructor() {
    EventsOn("output", (line: string) => {
      console.log(line);
      this.outputReceived$.next(line);
    });

    EventsOn("log", (line: any) => {
      console.log(line);
      this.logReceived$.next(line);
    });
  }
}
