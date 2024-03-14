import { Injectable } from '@angular/core';
import { FTree } from '../fnode/fnode';

@Injectable({
  providedIn: 'root',
})
export class TreeService {
    trees: FTree[] = [];
}
