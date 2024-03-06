import { Component, Input, HostListener, EventEmitter, Output } from '@angular/core';

export type MenuEntry = {
  label: string;
  callback: () => void;
};

export type MenuContent = {
  label: string,
  entries: MenuEntry[],
}[]

@Component({
  selector: 'app-menu-bar',
  standalone: true,
  imports: [],
  templateUrl: './menu-bar.component.html',
  styleUrl: './menu-bar.component.scss'
})
export class MenuBarComponent {
  @Input({required: true}) content: MenuContent | undefined;
  private inside: boolean = false;
  private someText: string = "";
  @Output() entrySelected = new EventEmitter<string>();

  @HostListener("click")
  clicked() {
    this.inside = true;
  }

  @HostListener("document:click")
  clickedOut() {
    this.someText = this.inside
      ? "Event Triggered"
      : "Event Triggered Outside Component";
    this.inside = false;
  }
}
