import {
  Component,
  Input,
  ElementRef,
  ViewChild,
  OnInit,
  DestroyRef,
  Output,
  EventEmitter
} from '@angular/core';
import { NodeLink } from '../fnode/fnode';
import { Subject } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Point } from '../point';

@Component({
  selector: 'app-node-link-ghost',
  standalone: true,
  templateUrl: './node-link-ghost.component.html',
  styleUrls: ['./node-link-ghost.component.scss']
})
export class NodeLinkGhostComponent implements OnInit {
  
  @Input({required: true}) requestRedraw$!: Subject<[HTMLElement, Point]>; //fromSocket, mousePosition, snapToSocket
  @Output() removed$ = new EventEmitter<NodeLink>();
  
  startSocket!: HTMLElement | null;
  randomGradientId: string =  crypto.randomUUID()

  constructor(private destroyRef: DestroyRef) {}

  @ViewChild("linkPath") pathElement?: ElementRef<SVGPathElement>;

  ngOnInit() {
    this.requestRedraw$.pipe(takeUntilDestroyed(this.destroyRef))
      .subscribe(drawRequest => {
          this.updatePath(...drawRequest);
      });
  }

  updatePath(startSocket: HTMLElement, position: Point) {
    const rectFrom = startSocket?.querySelector(".socket")?.getBoundingClientRect();
    if(!rectFrom) return;

    let p1 = [
      rectFrom.x + rectFrom.width / 2,
      rectFrom.y + rectFrom.height / 2 - 0.001, //+0.01 as a workaround for invisible straight paths bug in chrome
    ]

    let p2: number[] = [];

    p2 = [ // todo cursor pos
      position.x,
      position.y,
    ]
    
    let p1b = [
      (p1[0]! + p2[0]!) / 2,
      p1[1]!,
    ]

    let p2b = [
      (p1[0]! + p2[0]!) / 2,
      p2[1]!,
    ];

    this.pathElement?.nativeElement.setAttribute("d", `M ${p1[0]} ${p1[1]} C ${p1b[0]} ${p1b[1]} ${p2b[0]} ${p2b[1]} ${p2[0]} ${p2[1]}`);
  }
}
