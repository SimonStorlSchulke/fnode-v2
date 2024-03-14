import {
  Component,
  Input,
  AfterViewInit,
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
import { AddLink, RemoveLink } from '../../../../wailsjs/go/controller/App';


@Component({
  selector: 'app-node-link',
  standalone: true,
  templateUrl: './node-link.component.html',
  styleUrls: ['./node-link.component.scss']
})
export class NodeLinkComponent implements AfterViewInit, OnInit {
  @Input({required: true}) editor!: HTMLElement;
  @Input({required: true}) nodeLink!: NodeLink;
  @Input({required: true}) nodeChanged$!: Subject<string>;

  @Output() removed$ = new EventEmitter<NodeLink>();

  fromOutputElement?: HTMLElement | null;
  toInputElement?: HTMLElement | null;

  randomGradientId: string =  crypto.randomUUID()

  fromColor: string = "#55f";
  toColor: string = "#55f";

  constructor(private destroyRef: DestroyRef) {}

  @ViewChild("linkPath") pathElement?: ElementRef<SVGPathElement>;

  ngOnInit() {
    this.findSockets();
    this.nodeChanged$.pipe(takeUntilDestroyed(this.destroyRef))
      .subscribe(nodeId => {
        if(nodeId == this.nodeLink.FromNode || nodeId == this.nodeLink.ToNode) {
          this.updatePath();
        }
      });
  }

  ngAfterViewInit() {
    this.updatePath();
  }

  findSockets() {
    this.fromOutputElement = document.querySelector(
      `#output_${this.nodeLink.FromNode}__${this.nodeLink.FromOutput}`
    ) ?? null;

    this.toInputElement = document.querySelector(
      `#input_${this.nodeLink.ToNode}__${this.nodeLink.ToInput}`
    ) ?? null;

    this.fromColor = this.fromOutputElement?.style.backgroundColor ?? "#f00";

    this.toColor = this.toInputElement?.style.backgroundColor ?? "#f00";

    if (this.fromColor != this.toColor) {
      this.addImplicitConversionNote();
    }
  }

  addImplicitConversionNote() {

  }

  updatePath() {

    let p1 = [
      this.fromOutputElement!.getBoundingClientRect().x + 12,
      this.fromOutputElement!.getBoundingClientRect().y - 24.01, //+0.01 as a workaround for invisible straight paths bug in chrome
    ]

    let p2 = [
      this.toInputElement!.getBoundingClientRect().x,
      this.toInputElement!.getBoundingClientRect().y - 24,
    ]

    let p1b = [
      (p1[0]! + p2[0]!) / 2,
      p1[1]!,
    ]

    let p2b = [
      (p1[0]! + p2[0]!) / 2,
      p2[1]!,
    ]


    this.pathElement?.nativeElement.setAttribute("d", `M ${p1[0]} ${p1[1]} C ${p1b[0]} ${p1b[1]} ${p2b[0]} ${p2b[1]} ${p2[0]} ${p2[1]}`);

    // quite hacky, but whe have the element here anyway so it's cheap
    this.toInputElement!.parentElement!.querySelector("input")!.style.display = "none";
  }

  async remove() {
    await RemoveLink(this.nodeLink);
    this.toInputElement!.parentElement!.querySelector("input")!.style.display = "block"
    this.removed$.next(this.nodeLink);
  }
}
