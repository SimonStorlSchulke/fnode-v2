import {
  ChangeDetectionStrategy,
  ChangeDetectorRef,
  Component,
  ElementRef,
  HostListener,
  OnInit,
  ViewChild,
  inject,
} from "@angular/core";
import {
  GetTree,
  ParseTree,
  ClearTree,
  ParseTreePreview,
} from "../../../../wailsjs/go/controller/App";
import { FTree, NodeOption } from "../fnode/fnode";
import { FNodeComponent } from "../fnode/fnode.component";
import { NodeLinkComponent } from "../node-link/node-link.component";
import { Subject } from "rxjs";
import { FNodeService } from "../fnode.service";
import { Point } from "../point";
import { NodeLinkGhostComponent } from "../node-link-ghost/node-link-ghost.component";

@Component({
  selector: "app-fnode-editor",
  standalone: true,
  imports: [FNodeComponent, NodeLinkComponent, NodeLinkGhostComponent],
  templateUrl: "./fnode-editor.component.html",
  styleUrl: "./fnode-editor.component.scss",
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class FNodeEditorComponent implements OnInit {
  fNodeSv = inject(FNodeService);
  changeDetectorRef = inject(ChangeDetectorRef);
  elRef = inject(ElementRef);

  tree?: FTree;
  nodeChanged$ = new Subject<string>();
  viewTransform = {
    zoom: 1,
    scrollX: 1,
    scrollY: 1,
  };

  @ViewChild("editor") grid!: ElementRef<HTMLElement>;

  @HostListener("wheel", ["$event"])
  onMousewheel(event: WheelEvent) {
    this.zoom(event.deltaY > 0 ? -0.1 : 0.1, event.clientX, event.clientY);
  }

  dragStartingPoint: Point = {
    x: 0,
    y: 0,
  };
  dragging = false;

  constructor() {
    this.fNodeSv.nodeAdded$.subscribe(() => {
      this.getTree();
    });
  }

  async ngOnInit() {
    await this.getTree();
    this.updateViewTransform();
  }

  onDragStart(event: MouseEvent) {
    this.dragging = true;
    this.dragStartingPoint = {
      x: event.clientX - this.viewTransform.scrollX,
      y: event.clientY - this.viewTransform.scrollY,
    };
  }

  onDragEnd(event: MouseEvent) {
    if (!this.dragging) return;
    this.dragging = false;
  }

  onDrag(event: MouseEvent) {
    if (this.dragging) {
      this.viewTransform.scrollX = event.clientX - this.dragStartingPoint.x;
      this.viewTransform.scrollY = event.clientY - this.dragStartingPoint.y;
      this.updateViewTransform();
    }
  }

  async getTree() {
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

  async onRemovedNode(nodeId: string) {
    await this.getTree();
    this.nodeChanged$.next(nodeId);
  }

  zoom(step: number, cursorX = 0, cursorY = 0) {
    this.viewTransform.zoom = Math.min(
      Math.max(this.viewTransform.zoom + step, 0.25),
      1.5
    );
    this.updateViewTransform();
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

  @HostListener("click", ["$event"])
  deselectNodes(event: MouseEvent) {
    const clickedGrid = (event.target as HTMLElement)?.classList.contains(
      "grid"
    ); //hacky...
    if (clickedGrid) {
      this.fNodeSv.activeNodeId = "";
      this.fNodeSv.selectedNodeIds = [];
      this.changeDetectorRef.markForCheck();
    }
  }

  emitNodePositionChangedEvent(nodeId: string) {
    window.setTimeout(() => {
      this.nodeChanged$.next(nodeId);
    }, 0);
  }
  private updateViewTransform() {
    this.elRef.nativeElement.style.setProperty(
      "--zoom",
      this.viewTransform.zoom
    );
    this.elRef.nativeElement.style.setProperty(
      "--scrollX",
      `${this.viewTransform.scrollX}px`
    );
    this.elRef.nativeElement.style.setProperty(
      "--scrollY",
      `${this.viewTransform.scrollY}px`
    );
    this.nodeChanged$.next("all");
  }
}
