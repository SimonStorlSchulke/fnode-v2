import { Component, Inject, inject } from '@angular/core';
import { FNodeService } from '../../editor/fnode.service';
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

  fNodeSv: FNodeService = inject(FNodeService);
  currentTab = 0;

  ngOnInit() {
    this.fNodeSv.ngOnInit();
  }

  getCategoryNames(): string[] {
    return this.fNodeSv.NodeCategories.map((_) => _.Name)
  }
}
