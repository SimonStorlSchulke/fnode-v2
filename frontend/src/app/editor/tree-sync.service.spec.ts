import { TestBed } from '@angular/core/testing';

import { TreeSyncService } from './tree-sync.service';

describe('TreeSyncService', () => {
  let service: TreeSyncService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TreeSyncService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
