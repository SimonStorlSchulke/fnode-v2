import { Injectable } from '@angular/core';
import { FTree, NodeOption } from './fnode/fnode';
import { ClearTree, GetTree } from '../../../wailsjs/go/controller/App';

@Injectable({
  providedIn: 'root'
})
export class ProjectModel {

  tree: FTree | null = null;

  async UpdateTree() {
    this.tree =  await GetTree();
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

  async clearTree() {
    await ClearTree();
    await this.UpdateTree();
  }

}
