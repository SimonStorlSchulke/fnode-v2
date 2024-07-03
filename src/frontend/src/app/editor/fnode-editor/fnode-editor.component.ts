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

  viewTransform = {
    zoom: 1,
    scrollX: 1,
    scrollY: 1,
  }


  @HostListener('wheel', ['$event'])
  onMousewheel(event: WheelEvent) {
    this.zoom(event.deltaY > 0 ? -0.1 : 0.1, event.clientX, event.clientY);
  }

  protected fNodeSv = inject(FNodeService);
  changeDetectorRef = inject(ChangeDetectorRef);
  elRef = inject(ElementRef);

  constructor() {
    this.fNodeSv.nodeAdded$.subscribe(() => {
      this.getTree();
    });
  }

  async ngOnInit() {
    await this.getTree();
    this.updateViewTransform();
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

  zoom(step: number, cursorX = 0, cursorY = 0) {
    this.viewTransform.zoom = Math.min(Math.max(this.viewTransform.zoom + step, 0.25), 1.5);
    this.updateViewTransform();
    //this.grid.nativeElement.style.transform = `scale(${this.viewTransorm.zoom})`;
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
  private updateViewTransform() {
    this.elRef.nativeElement.style.setProperty('--zoom', this.viewTransform.zoom);
    this.elRef.nativeElement.style.setProperty('--scrollX', `${this.viewTransform.scrollX}px`);
    this.elRef.nativeElement.style.setProperty('--scrollY', `${this.viewTransform.scrollY}px`);
  }
}
