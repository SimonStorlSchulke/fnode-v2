import { Component } from '@angular/core';
import { MenuContent, MenuBarComponent } from '../../ui/menu-bar/menu-bar.component';
import { WindowMaximise, WindowMinimise, WindowIsMaximised, WindowUnmaximise } from '../../../../wailsjs/runtime';
import { FeatherModule } from 'angular-feather';
import { IconsModule } from '../../icons/icons.module';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [
    MenuBarComponent,
    IconsModule
  ],
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss'
})
export class HeaderComponent {
  menuContent: MenuContent = [
    {
      key: "file",
      label: "File",
      entries: [
        {
          key: "save",
          label: "Save",
        }
      ]
    },
    {
      key: "edit",
      label: "Edit",
      entries: [
        {
          key: "undo",
          label: "undo",
          disabled: true,
        },
        {
          key: "redo",
          label: "Redo"
        },
      ]
    },
  ]

  cbTest() {
    console.log("this.menuContent")
    console.log(this.menuContent)
  }

  async maximize() {
    const isMaximized = await WindowIsMaximised();
    if(isMaximized) {
      WindowUnmaximise();
    } else {
      WindowMaximise();
    }
  }

  minimize() {
    WindowMinimise();
  }
}
