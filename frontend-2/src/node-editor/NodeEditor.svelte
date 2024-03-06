<script lang="ts">
  import { GetTestTree, ParseTree } from "../../wailsjs/go/controller/App";
  import Link from "./Link.svelte";
  import Node from "./Node.svelte";
  import type { NodeTree } from '../model';

  let tree: NodeTree

  async function getTree() {
    tree = await GetTestTree()

    console.log(tree)
  }


</script>

<div class="node-editor grow-1">
  <div class="grid">
    {#key tree}
    {#each tree?.Nodes ?? [] as node}
      <Node node={node} posX={node.Meta.PosX} posY={node.Meta.PosY} />
    {/each}

    {#each tree?.Links ?? [] as link}
      <Link nodeLink={link}/>
    {/each}
      {/key}
  </div>
</div>
<div class="hbox">
  <button on:click={getTree}>GetTree</button>
</div>

<style lang="scss">
  @use "../../src/assets/styles/color";

  $grid-lines-1: rgba(255, 255, 255, 0.06);
  $grid-lines-2: rgba(255, 255, 255, 0.171);

  .grid {
    background-color: color.$neutral-1;
    background-size: 80px 80px;
    background-image: linear-gradient(
        to right,
        $grid-lines-1 1px,
        transparent 1px
      ),
      linear-gradient(to bottom, $grid-lines-1 1px, transparent 1px),
      linear-gradient(
        to right,
        transparent 40px,
        $grid-lines-2 40px,
        $grid-lines-2 41px,
        transparent 41px
      ),
      linear-gradient(
        to bottom,
        transparent 40px,
        $grid-lines-2 40px,
        $grid-lines-2 41px,
        transparent 41px
      );
    width: 100%;
    height: 100%;
    position: relative;
  }
</style>
