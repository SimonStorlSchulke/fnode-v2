import { MenuEntry } from './../../ui/menu-bar/menu-bar.component';
import { Component, inject } from '@angular/core';
import { MenuContent, MenuBarComponent } from '../../ui/menu-bar/menu-bar.component';
import { WindowMaximise, WindowMinimise, WindowIsMaximised, WindowUnmaximise } from '../../../../wailsjs/runtime';
import { FeatherModule } from 'angular-feather';
import { IconsModule } from '../../icons/icons.module';
import { IOService } from '../../editor/io-service';

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
          key: "save-as",
          label: "Save as",
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
  ];

  ioSv = inject(IOService);

  MenuEntryCallback(entryKey: [string, string]) {
    console.log("WAAAdasA", entryKey)
    const entry = entryKey.join(".")
    switch (entry) {
      case "file.save-as":
        console.log("WAAAA")
        this.ioSv.SaveTree();
    }
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
