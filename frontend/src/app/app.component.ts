import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MenuBarComponent } from './ui/menu-bar/menu-bar.component';
import { HeaderComponent } from './layout/header/header.component';
import { PanelComponent } from './ui/panel/panel.component';
import { DragAndDropModule, DragEndEvent } from 'angular-draggable-droppable';
import { NgStyle } from '@angular/common';
import { FNodeComponent } from './editor/fnode/fnode.component';
import { FNodeEditorComponent } from './editor/fnode-editor/fnode-editor.component';
import { NodeAdderComponent } from './layout/node-adder/node-adder.component';
import { TerminalComponent } from './layout/terminal/terminal.component';
import { FileListComponent } from './layout/file-list/file-list.component';
import { NodeTreeComponent } from './node-tree/node-tree.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, NodeTreeComponent, MenuBarComponent, HeaderComponent, PanelComponent, DragAndDropModule, NgStyle, FNodeComponent, FNodeEditorComponent, NodeAdderComponent, TerminalComponent, FileListComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'frontend';

  x = 0;
  y = 0;

  dragEnd($event: DragEndEvent) {
    this.x += $event.x;
    this.y += $event.y;
  }
}
