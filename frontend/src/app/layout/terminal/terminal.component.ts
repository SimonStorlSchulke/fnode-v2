import { Component, ChangeDetectionStrategy, inject, ChangeDetectorRef, OnInit } from '@angular/core';
import { PanelComponent } from '../../ui/panel/panel.component';
import { ReversePipe } from '../../ui/reverse.pipe';
import { TerminalService } from './terminal.service';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { TabContainerComponent } from '../../ui/tab-container/tab-container.component';
import { DomSanitizer } from '@angular/platform-browser';

@Component({
  selector: 'app-terminal',
  standalone: true,
  imports: [
    PanelComponent,
    ReversePipe,
    TabContainerComponent
  ],
  templateUrl: './terminal.component.html',
  styleUrl: './terminal.component.scss',
})
export class TerminalComponent {

  changeDetectorRef = inject(ChangeDetectorRef);
  terminalSv = inject(TerminalService);
  domSanitizer = inject(DomSanitizer);
  currentTab = 0;


  text: string = ""
  log: string = ""

  constructor() {
    this.terminalSv.outputReceived$.pipe(takeUntilDestroyed()).subscribe(line => {
      this.text += line;
      this.text += "\n";
      this.changeDetectorRef.detectChanges()
    });

    this.terminalSv.logReceived$.pipe(takeUntilDestroyed()).subscribe(line => {
      if(line.startsWith("Info")) line = `<span style="color: #00b2ff">${line}</span>`
      if(line.startsWith("Warn")) line = `<span style="color: #d9be38">${line}</span>`
      if(line.startsWith("Error")) line = `<span style="color: #ff4444">${line}</span>`
      this.log += line;
      this.log += "\n";
      this.changeDetectorRef.detectChanges()
    });
  }

  clearOutput() {
    this.text = "";
    this.changeDetectorRef.markForCheck();
  }

}
