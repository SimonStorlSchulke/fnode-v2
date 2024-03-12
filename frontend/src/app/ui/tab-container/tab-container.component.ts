import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'sui-tab-container',
  standalone: true,
  imports: [],
  templateUrl: './tab-container.component.html',
  styleUrl: './tab-container.component.scss',
})
export class TabContainerComponent {

  @Input({required: true}) tabs: string[] = [];
  @Input() vertical = false;
  @Input() closableTabs = false;

  @Output() switched = new EventEmitter<number>();

  protected currentTab = 0;

  switchTab(index: number) {
    this.currentTab = index
    this.switched.next(index);
  }
}
