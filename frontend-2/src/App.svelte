<script lang="ts">
  import { GetTestTree, Greet, ParseTree } from '../wailsjs/go/controller/App.js'
  import TitleBar from './layout/HeaderMenu.svelte';
  import NodeEditor from './node-editor/NodeEditor.svelte';
  import Terminal from './layout/Terminal.svelte';

  let output: string[]
  async function parseTree() {
    output = await ParseTree()
  }

</script>

<main>
  <TitleBar/>
  <NodeEditor/>
  <Terminal text={output?.join("\n")}/>
  <button on:click={parseTree}>Parse</button>
</main>

<style>

main {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.tree {
  background-color: #222;
  flex-grow: 1;
}


</style>