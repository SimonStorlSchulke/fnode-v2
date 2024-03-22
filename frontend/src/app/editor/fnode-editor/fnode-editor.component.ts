import {
  ChangeDetectionStrategy,
  ChangeDetectorRef,
  Component,
  ElementRef,
  HostListener,
  OnInit,
  ViewChild,
  inject,
} from '@angular/core';
import {
  GetTree,
  ParseTree,
  ClearTree,
  ParseTreePreview,
} from '../../../../wailsjs/go/controller/App';
import { FTree, NodeOption } from '../fnode/fnode';
import { FNodeComponent } from '../fnode/fnode.component';
import { NodeLinkComponent } from '../node-link/node-link.component';
import { Subject } from 'rxjs';
import { FNodeService } from '../fnode.service';

@Component({
  selector: 'app-fnode-editor',
  standalone: true,
  imports: [FNodeComponent, NodeLinkComponent],
  templateUrl: './fnode-editor.component.html',
  styleUrl: './fnode-editor.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class FNodeEditorComponent implements OnInit {
  protected tree?: FTree;
  protected nodeChanged$ = new Subject<string>();
  @ViewChild('editor') grid!: ElementRef<HTMLElement>;

  protected fNodeSv = inject(FNodeService);
  changeDetectorRef = inject(ChangeDetectorRef);

  constructor() {
    this.fNodeSv.nodeAdded$.subscribe(() => {
      this.getTree();
    });
  }

  async ngOnInit() {
    await this.getTree();
  }

  protected async getTree() {
    this.tree = await GetTree();

    for (let node of this.tree!.Nodes) {
      const newOptions = new Map<string, NodeOption>();
      for (let optionKey of Object.keys(node.Options)) {
        newOptions.set(optionKey, (node.Options as any)[optionKey]);
      }

      node.Options = newOptions;
    }
    this.changeDetectorRef.markForCheck();
  }

  protected async onRemovedNode(nodeId: string) {
    await this.getTree();
    this.nodeChanged$.next(nodeId);
  }

  currentZoom = 1;
  zoom(step: number) {
    this.currentZoom += step;
    this.grid.nativeElement.style.transform = `scale(${this.currentZoom})`;
    this.nodeChanged$.next('all');
  }

  protected async parseTree() {
    await ParseTree();
  }

  protected async parseTreePreview() {
    await ParseTreePreview();
  }

  async clearTree() {
    await ClearTree();
    await this.getTree();
  }

  @HostListener('click', ['$event'])
  protected deselectNodes(event: MouseEvent) {
    const clickedGrid = (
      event.target as HTMLElement
    ).children[0]?.classList.contains('grid'); //hacky...
    if (clickedGrid) {
      this.fNodeSv.activeNodeId = '';
      this.fNodeSv.selectedNodeIds = [];
      this.changeDetectorRef.markForCheck();
    }
  }

  protected emitNodePositionChangedEvent(nodeId: string) {
    this.nodeChanged$.next(nodeId);
  }
}
