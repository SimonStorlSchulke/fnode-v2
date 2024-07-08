import { Injectable } from '@angular/core';
import { FileList } from '../../editor/fnode/fnode';
import { AddDirectoriesFromDialog, AddLooseFilesFromDialog, ClearFileList, GetFileList, RemoveDirectory, RemoveLooseFile } from '../../../../wailsjs/go/controller/App';

@Injectable({
  providedIn: 'root'
})
export class FileListService {

  public list: FileList = {
    LooseFiles: [],
    Directories: [],
  }

  async update() {
    this.list = await GetFileList();
    console.log(this.list)
  }


  public async addFilesDialog() {
    await AddLooseFilesFromDialog();
    this.update();
  }


  public async addDirectoryDialog() {
    await AddDirectoriesFromDialog();
    this.update();
  }

  public async removeLooseFile(index: number) {
    await RemoveLooseFile(index);
    this.update();
  }

  public async removeDirectory(index: number) {
    await RemoveDirectory(index);
    this.update();
  }


  public async clear(index: number) {
    await ClearFileList();
    this.update();
  }
}
