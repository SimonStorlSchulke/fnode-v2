<script lang="ts">
  import { clickOutside } from "../lib-ui/click-outside-directive";
  import { UpdateInputDefaultValue } from "../../wailsjs/go/controller/App"
  import type { FNode } from '../model';
  import { onMount } from 'svelte';
  import { FTypeColors } from './ftype-colors';

  export let node: FNode;

  let selected: boolean = true;

  export let posX: number = 100;
  export let posY: number = 100;

  let moving = false;
  let moveEvent: Event;
	
	function onMouseDown() {
		moving = true;
    }

    onMount(() => {
      console.log(posX)
      moveEvent = new Event(`MOVE_${node.Id}`);
    })

  function updateInputValue(inputIndex: number, value: Event, valueType: number) {
      UpdateInputDefaultValue(node.Id, inputIndex, (value.target as HTMLInputElement).value, valueType)
  }

	
	function onMouseMove(e) {
		if (moving) {
			posX += e.movementX;
			posY += e.movementY;

            window.dispatchEvent(moveEvent)
		}
	}
	
	function onMouseUp() {
		moving = false;
        window.dispatchEvent(moveEvent)
    }

</script>

<div class="fnode" id={node.Id} on:click={() => {selected = true}} use:clickOutside={() => {selected = false}} class:selected style="top: {posY}px; left: {posX}px">
  <div class="fnode-header" on:mousedown={onMouseDown}>
    <span>{node.Type}</span>
  </div>
  <div class="fnode-body vbox gap-1">

    {#each node.Outputs as output, i}
      <div class="foutput hbox center-items gap-3 end">
        <span>{output.Name}</span>
        <div id={"output_" + node.Id + "__" + i} style="background-color: {FTypeColors[output.Type]}" class="socket"></div>
      </div>
    {/each}
    {#each Object.keys(node.Options) as optionKey} <!--TODO ugly af-->
      <div class="hbox gap-3">
        <span>{optionKey}</span>

        <select name="pets" id="pet-select">
          <option value={node.Options[optionKey].SelectedOption}>{node.Options[optionKey].SelectedOption}</option>
          {#each node.Options[optionKey].Choices as choice}
            {#if choice != node.Options[optionKey].SelectedOption}
              <option value={choice}>{choice}</option>
              {/if}
          {/each}
        </select>
      </div>
    {/each}
    {#each node.Inputs as input, i}
      <div class="finput hbox center-items gap-3 space-between">
        <div id={"input_" + node.Id + "__" + i} style="background-color: {FTypeColors[input.Type]}" class="socket"></div>
        <label>{input.Name}</label>
        <input type="number" value={input.DefaultValue} on:input={(e) => updateInputValue(i, e, input.Type)}/>
      </div>
      {/each}
  </div>
</div>

<svelte:window on:mouseup={onMouseUp} on:mousemove={onMouseMove} />

<style lang="scss">
  @use "src/assets/styles/color";

  .fnode {
    user-select: none;
    position: absolute;
    background-color: color.$neutral-3;
    outline: 1px solid color.$neutral-1;
    color: color.$text-0;
    border-radius: 4px;
    box-shadow: 2px 3px 12px #0008;
    -webkit-user-select: none;   
    -moz-user-select: none;
    -ms-user-select: none;
    user-select: none;
    cursor: default;

    &.selected {
      outline: 2px solid color.$accent-1;
    }
  }

  .fnode-header {
    user-select: none;
    color: color.$text-1;
    padding: 6px 8px;
    font-weight: 500;
    background-color: rgb(95, 33, 33);
  }

  .fnode-body {
    padding: 8px;
  }

  .socket {
    position: absolute;
    width: 12px;
    height: 12px;
    outline: 1px solid color.$neutral-0;
    border-radius: 50%;
  }

  .finput {
    height: 30px;
    position: relative;
    padding: 0 4px;

    input {
      width: 60px;
      height: 20px;
    }

    .socket {
      left: -14px;
    }
  }

  .foutput {
    padding: 0 4px;
    .socket {
      right: -6px;
    }
  }
</style>
