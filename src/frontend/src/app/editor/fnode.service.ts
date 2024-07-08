import { Injectable } from '@angular/core';
import { NodeCategory } from './fnode/fnode';
import { AddConnectedNode, AddNode, GetNodeCategories } from '../../../wailsjs/go/controller/App';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FNodeService {

  NodeCategories: NodeCategory[] = [];
  nodeAdded$ = new Subject<void>();

  selectedNodeIds: string[] = [];
  activeNodeId: string = "";

  ngOnInit() {
    GetNodeCategories().then((data: NodeCategory[]) => {
      this.NodeCategories = data
    });
  }
  
  async addNode(ofCategory: string, ofType: string, posX: number = 100, posY: number = 100) {

    let addedNodeId = "";

    if (this.activeNodeId != "") {
      addedNodeId = await AddConnectedNode(`${ofCategory}.${ofType}`, this.activeNodeId);
    } else {
      addedNodeId = await AddNode(`${ofCategory}.${ofType}`, posX, posY);
    }

    this.activeNodeId = addedNodeId;

    this.nodeAdded$.next();
  }

}
