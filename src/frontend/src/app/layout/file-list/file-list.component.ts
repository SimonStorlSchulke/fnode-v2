import { Component, OnInit, inject } from '@angular/core';
import { FileListService } from './file-list.service';
import { FeatherModule } from 'angular-feather';
import { AddDirectoriesFromDialog, AddLooseFilesFromDialog } from '../../../../wailsjs/go/controller/App';
import { FileNameFromPath } from "../../ui/filename-from-path.pipe";

@Component({
    selector: 'app-file-list',
    standalone: true,
    templateUrl: './file-list.component.html',
    styleUrl: './file-list.component.scss',
    imports: [FeatherModule, FileNameFromPath]
})
export class FileListComponent implements OnInit {
  fileListSv = inject(FileListService);

  ngOnInit() {
    this.fileListSv.update();
  }



}
