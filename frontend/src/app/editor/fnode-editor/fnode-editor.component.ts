import { Component, OnInit, inject } from '@angular/core';
import { GetTree, ParseTree, ClearTree, ParseTreePreview } from '../../../../wailsjs/go/controller/App';
import { FTree, NodeOption } from '../fnode/fnode';
import { FNodeComponent } from '../fnode/fnode.component';
import { NodeLinkComponent } from '../node-link/node-link.component';
import { Subject, tap } from 'rxjs';
import { NodeAdderService } from '../../layout/node-adder/node-adder.service';

@Component({
  selector: 'app-fnode-editor',
  standalone: true,
  imports: [
    FNodeComponent,
    NodeLinkComponent
  ],
  templateUrl: './fnode-editor.component.html',
  styleUrl: './fnode-editor.component.scss'
})
export class FNodeEditorComponent implements OnInit {

  tree?: FTree;
  nodeChanged$ = new Subject<string>()

  nodAdderSv = inject(NodeAdderService);

  constructor() {
    this.nodAdderSv.nodeAdded$.subscribe(() => {
      this.getTree();
    });
  }

  async  ngOnInit() {
    await this.getTree();
  }

  async getTree() {
    this.tree =  await GetTree();
    console.log(this.tree)

    for (let node of this.tree!.Nodes) {
      const newOptions = new Map<string, NodeOption>();
      for (let optionKey of Object.keys(node.Options)) {
        newOptions.set(optionKey, (node.Options as any)[optionKey])
      }

      node.Options = newOptions;
      console.log("node.Options")
      console.log(node.Options)
    }
  }

  async parseTree() {
    await ParseTree();
  }

  async parseTreePreview() {
    await ParseTreePreview();
  }

  async clearTree() {
    await ClearTree();
    await this.getTree();
  }

  emitNodePositionChangedEvent(nodeId: string) {
    this.nodeChanged$.next(nodeId);
  }
}
