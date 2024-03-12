import { Injectable } from '@angular/core';
import { NodeCategory } from '../../editor/fnode/fnode';
import { AddNode, GetNodeCategories } from '../../../../wailsjs/go/controller/App';
import { Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class NodeAdderService {

  NodeCategories: NodeCategory[] = [];

  nodeAdded$ = new Subject<void>();

  ngOnInit() {
    GetNodeCategories().then((data: NodeCategory[]) => {
      this.NodeCategories = data
    });
  }
  
  async addNode(ofCategory: string, ofType: string, posX: number = 100, posY: number = 100) {
    await AddNode(`${ofCategory}.${ofType}`, posX, posY);
    this.nodeAdded$.next();
  }

}
