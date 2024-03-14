import { LoadFile, Save, SaveAs } from '../../../wailsjs/go/controller/App';
import { FTree } from './fnode/fnode';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class IOService {

  async SaveTree() {
    SaveAs()
  }

  async LoadTree() {
    await LoadFile()
  }

}
