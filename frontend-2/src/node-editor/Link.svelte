<script lang="ts" context="module">
    function linkUpdateEventName(link: NodeLink): string {
        return `UPDATE_${link.FromNode}${link.FromOutput}__${link.ToNode}${link.ToInput}`
    }
</script>

<script lang="ts">

  import { onMount, onDestroy } from 'svelte';
  import type { NodeLink } from '../model';

  export let nodeLink: NodeLink;

  let fromOutputElement: HTMLElement;
  let toInputElement: HTMLElement;
  let pathElement: SVGPathElement;

  let fromColor: string
  let toColor: string
  let randomGradientId = crypto.randomUUID()

  onMount(() => {
    getSocketElements()
    updatePath()

    // Probably not "the Svelte way"
    window.addEventListener(`MOVE_${nodeLink.FromNode}`, updatePath)
    window.addEventListener(`MOVE_${nodeLink.ToNode}`, updatePath)
  })

  onDestroy(() => {
    window.removeEventListener(linkUpdateEventName(nodeLink), updatePath)
  })

  function getSocketElements() {
    fromOutputElement = document.querySelector(`#output_${nodeLink.FromNode}__${nodeLink.FromOutput}`)
    toInputElement = document.querySelector(`#input_${nodeLink.ToNode}__${nodeLink.ToInput}`)


    fromColor = fromOutputElement.style.backgroundColor

    toColor = toInputElement.style.backgroundColor
    console.log("from", fromColor)
    console.log("to", toColor)
  }

    function updatePath() {

        let p1 = [
            fromOutputElement?.getBoundingClientRect().x + 12,
            fromOutputElement?.getBoundingClientRect().y - 24.01, //+0.01 as a workaround for invisible straight paths bug in chrome
        ]

        let p2 = [
            toInputElement?.getBoundingClientRect().x,
            toInputElement?.getBoundingClientRect().y - 24,
        ]

        let p1b = [
            (p1[0]! + p2[0]!) / 2,
            p1[1]!,
        ]

        let p2b = [
            (p1[0]! + p2[0]!) / 2,
            p2[1]!,
        ]


      pathElement.setAttribute("d", `M ${p1[0]} ${p1[1]} C ${p1b[0]} ${p1b[1]} ${p2b[0]} ${p2b[1]} ${p2[0]} ${p2[1]}`);

      // quite hacky, but whe have the element here anyway so it's cheap
      toInputElement.parentElement.querySelector("input").style.visibility = "hidden"
    }
</script>


<svg class="connection input-number output-number socket-input-number-value socket-output-number-value">
    <path style="stroke: url(#{randomGradientId})" bind:this={pathElement} class="main-path" d="M 270 255 C 362 255 408 370 500 370"></path>
    <defs>
        <linearGradient id={randomGradientId} x1="0%" y1="0%" x2="100%" y2="0%">
            <stop offset="0%"   stop-color={fromColor}/>
            <stop offset="100%" stop-color={toColor}/>
        </linearGradient>
    </defs>
</svg>


<style>
.connection .main-path {
    fill: none;
    stroke-width: 3px;
}

.connection {
    overflow: visible !important;
    position: absolute;
    z-index: 1;
    pointer-events: none;
    top: 0
}
</style>