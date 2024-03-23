import { Component, Input, HostBinding } from '@angular/core';

@Component({
  selector: 'sui-panel',
  standalone: true,
  imports: [],
  templateUrl: './panel.component.html',
  styleUrl: './panel.component.scss',
})
export class PanelComponent {
  @Input() direction: "vbox" | "hbox" = "vbox";
  @Input() label: string = "";
  @Input() noPadding = false;
}
