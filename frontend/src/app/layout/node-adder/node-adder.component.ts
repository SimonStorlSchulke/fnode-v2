import { Component, Inject, inject } from '@angular/core';
import { NodeAdderService } from './node-adder.service';
import { AddNode } from '../../../../wailsjs/go/controller/App';
import { TabContainerComponent } from '../../ui/tab-container/tab-container.component';

@Component({
  selector: 'app-node-adder',
  standalone: true,
  imports: [
    TabContainerComponent
  ],
  templateUrl: './node-adder.component.html',
  styleUrl: './node-adder.component.scss'
})
export class NodeAdderComponent {

  nodeAdderSv: NodeAdderService = inject(NodeAdderService);
  currentTab = 0;

  ngOnInit() {
    this.nodeAdderSv.ngOnInit();
  }

  getCategoryNames(): string[] {
    return this.nodeAdderSv.NodeCategories.map((_) => _.Name)
  }
}
