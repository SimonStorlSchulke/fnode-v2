import { Component, Input, HostBinding, OnInit, ViewChild, ElementRef, Output, EventEmitter, inject} from '@angular/core';
import { FNode, FInput, FType } from './fnode';
import {
  UpdateInputDefaultValue,
  AddLink,
  UpdateNodePosition,
  UpdateUption,
  RemoveNode
} from '../../../../wailsjs/go/controller/App';
import { FTypeColors } from './ftype-colors';
import { DraggableDirective, DragAndDropModule, DragMoveEvent, DragEndEvent } from 'angular-draggable-droppable';
import { NgStyle } from '@angular/common';
import { FeatherModule } from 'angular-feather';
import { FNodeService } from '../fnode.service';
import {CdkDrag} from '@angular/cdk/drag-drop';

type SocketType = "input" | "output";

@Component({
  selector: 'app-fnode',
  standalone: true,
  imports: [DragAndDropModule, NgStyle, FeatherModule],
  templateUrl: './fnode.component.html',
  styleUrl: './fnode.component.scss',
})
export class FNodeComponent implements OnInit{
  @Input({required: true}) fnode!: FNode;

  @HostBinding("style.left") posX = "100px";
  @HostBinding("style.top") posY = "100px";
  @Output() redrawLinks = new EventEmitter<string>();
  @Output() changedNode = new EventEmitter<void>();
  @Output() removedNode = new EventEmitter<void>();

  @ViewChild("header") headerElement!: ElementRef;
  @ViewChild("content") contentElement!: ElementRef;

  fNodeSv = inject(FNodeService);

  ngOnInit() {
    this.updatePosition();
  }

  selectNode(event: MouseEvent) {
    this.fNodeSv.activeNodeId = this.fnode.Id;
    if(event.ctrlKey) {
      this.fNodeSv.selectedNodeIds.push(this.fnode.Id);
    } else {
      this.fNodeSv.selectedNodeIds = [this.fnode.Id];
    }

  }

  updatePosition() {
    this.posX = `${this.fnode.Meta.PosX}px`;
    this.posY = `${this.fnode.Meta.PosY}px`;

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
    const success = await UpdateUption(this.fnode.Id, optionKey, value);
    if (success) {
      this.fnode.Options.get(optionKey)!.SelectedOption = value;
      this.changedNode.next();
    }
  }

  protected readonly FTypeColors = FTypeColors;

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
    this.fnode.Meta.PosX += event.x;
    this.fnode.Meta.PosY += event.y;
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
