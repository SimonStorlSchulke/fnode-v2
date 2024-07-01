import { Component, ElementRef, Injector, ViewChild } from '@angular/core';

@Component({
  selector: 'app-node-tree',
  standalone: true,
  imports: [],
  templateUrl: './node-tree.component.html',
  styleUrl: './node-tree.component.scss'
})
export class NodeTreeComponent {

    constructor(private injector: Injector) {}

    @ViewChild("rete") container!: ElementRef;

    ngAfterViewInit(): void {
      const el = this.container.nativeElement;

      if (el) {
      }
    }
}
