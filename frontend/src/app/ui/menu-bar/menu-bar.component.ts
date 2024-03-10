import { Component, Input, HostListener, EventEmitter, Output } from '@angular/core';

export type MenuEntry = {
  key: string,
  label: string,
  disabled?: boolean,
  tooltip?: string,
};

export type MenuContent = {
  key: string,
  label: string,
  entries: MenuEntry[],
}[]

@Component({
  selector: 'sui-menu-bar',
  standalone: true,
  imports: [],
  templateUrl: './menu-bar.component.html',
  styleUrl: './menu-bar.component.scss'
})
export class MenuBarComponent {
  @Input({required: true}) content: MenuContent | undefined;
  @Input() disabledKeys: [string, string][] = [];
  @Output() entrySelected = new EventEmitter<[string, string]>();

  protected openIndex: number = -1;

  private inside: boolean = false;
  private someText: string = "";

  @HostListener("click")
  clicked() {
    this.inside = true;
  }

  @HostListener("document:click")
  clickedOut() {
    if(!this.inside) {
      this.closeMenu();
    }
    this.inside = false;
  }

  toggleMenu(index: number) {
    if (index == this.openIndex) {
      this.closeMenu();
    } else {
      this.openIndex = index;
    }
  }

  openOnHover(menuIndex: number) {
    if (this.openIndex !== -1) {
      this.openIndex = menuIndex;
    }
  }

  closeMenu() {
    this.openIndex = -1;
  }
}
