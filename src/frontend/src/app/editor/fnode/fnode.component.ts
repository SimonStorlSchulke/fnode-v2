import { Component, Input, HostBinding, OnInit, ViewChild, ElementRef, Output, EventEmitter, inject, HostListener} from '@angular/core';
import { FNode, FInput, FType } from './fnode';
import {
  UpdateInputDefaultValue,
  AddLink,
  UpdateNodePosition,
  UpdateUption as UpdateOption,
  RemoveNode
} from '../../../../wailsjs/go/controller/App';
import { FTypeColors } from './ftype-colors';
import { DragAndDropModule, DragMoveEvent, DragEndEvent, DragStartEvent } from 'angular-draggable-droppable';
import { NgStyle } from '@angular/common';
import { FeatherModule } from 'angular-feather';
import { FNodeService } from '../fnode.service';
import { NodeLinkGhostComponent } from '../node-link-ghost/node-link-ghost.component';
import { Subject } from 'rxjs';
import { Point } from '../point';

type SocketType = "fromInput" | "fromOutput";

@Component({
  selector: 'app-fnode',
  standalone: true,
  imports: [DragAndDropModule, NgStyle, FeatherModule, NodeLinkGhostComponent],
  templateUrl: './fnode.component.html',
  styleUrl: './fnode.component.scss',
})
export class FNodeComponent implements OnInit{
  @Input({required: true}) fnode!: FNode;
  @Input({required: true}) viewTransform!: {zoom: number, scrollX: number, scrollY: number};
  @HostBinding("style.left") cssPosX = "calc(100px * var(--zoom))";
  @HostBinding("style.top") cssPosY = "calc(100px * var(--zoom))";
  @Output() redrawLinks = new EventEmitter<string>();
  @Output() changedNode = new EventEmitter<void>();
  @Output() removedNode = new EventEmitter<void>();

  @ViewChild("header") headerElement!: ElementRef;
  @ViewChild("content") contentElement!: ElementRef;

  fNodeSv = inject(FNodeService);

  requestRedrawGhostLink$ = new Subject<[HTMLElement, Point]>();
  isConnectingNodes = false;

  currentDraggingSocket?: HTMLElement;

  ngOnInit() {
    this.updatePosition();
  }

  onStartDraggingSocket(socket: HTMLElement) {
    window.setTimeout(() => {
      this.currentDraggingSocket = socket;
      this.isConnectingNodes = true;
    }, 0)
  }


  onEndDraggingSocket() {
    this.isConnectingNodes = false;
  }

  @HostListener('document:mousemove', ['$event'])
  onDraggingSocket(event: MouseEvent) {
    if(!this.isConnectingNodes) return;
    this.requestRedrawGhostLink$.next([this.currentDraggingSocket!, {x: event.clientX, y: event.clientY}]);
   }

  @HostListener('click', ['$event'])
  selectNode(event: MouseEvent) {
    this.fNodeSv.activeNodeId = this.fnode.Id;
    if(event.ctrlKey) {
      this.fNodeSv.selectedNodeIds.push(this.fnode.Id);
    } else {
      this.fNodeSv.selectedNodeIds = [this.fnode.Id];
    }

  }

  updatePosition() {
    this.cssPosX = `calc(${this.fnode.Meta.PosX}px * var(--zoom) + var(--scrollX)  )`;
    this.cssPosY = `calc(${this.fnode.Meta.PosY}px * var(--zoom) + var(--scrollY) )`;

    this.redrawLinks.next(this.fnode.Id);
    UpdateNodePosition(this.fnode.Id, this.fnode.Meta.PosX, this.fnode.Meta.PosY);
  }

  getHtmlInputType(input: FInput) {
    switch (input.Type) {
      case FType.Float:
        return "number";
      case FType.Int:
        return "number";
      case FType.String:
        return "text";
      case FType.Bool:
        return "checkbox";
      default:
        return "text";
    }
  }

  updateInputValue(inputIndex: number, event: Event, valueType: number) {
    let value: any = (event.target as HTMLInputElement).value;
    if(valueType == FType.Bool) {
      value = (event.target as HTMLInputElement).checked;
    }
    UpdateInputDefaultValue(this.fnode.Id, inputIndex, value, valueType)
  }

  async updatedOption(optionKey: string, selectedChoice: EventTarget | null) {
    const value = (selectedChoice as HTMLSelectElement).value
    const success = await UpdateOption(this.fnode.Id, optionKey, value);
    if (success) {
      this.fnode.Options.get(optionKey)!.SelectedOption = value;
      this.changedNode.next();
    }
  }

  readonly FTypeColors = FTypeColors;

  dragging = false;
  dragOffsetX = 0;
  dragOffsetY = 0;
  onDrag(event: DragMoveEvent) {
    this.dragging = true;
    this.dragOffsetX = event.x;
    this.dragOffsetY = event.y;
    this.redrawLinks.next(this.fnode.Id);

  }

  onDragEnd(event: DragEndEvent) {
    this.dragging = false;
    this.fnode.Meta.PosX = Math.floor(this.fnode.Meta.PosX + event.x / this.viewTransform.zoom);
    this.fnode.Meta.PosY = Math.floor(this.fnode.Meta.PosY + event.y / this.viewTransform.zoom);
    this.updatePosition();
  }

  async addLink(fromNode: string, fromInput: number, toNode: string, toInput: number, fromSocketType: SocketType, toSocketType: SocketType) {
    if(fromSocketType == toSocketType) return;
    await AddLink({
      FromNode: fromNode,
      FromOutput: fromInput,
      ToNode: toNode,
      ToInput: toInput,
    });
    this.changedNode.next();
    this.redrawLinks.next(this.fnode.Id);
  }

  async removeNode() {
    if(this.fNodeSv.activeNodeId == this.fnode.Id) this.fNodeSv.activeNodeId = "";
    await RemoveNode(this.fnode.Id);
    this.removedNode.next();
  }
}
