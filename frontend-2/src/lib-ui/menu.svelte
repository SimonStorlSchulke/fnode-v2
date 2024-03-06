<script context="module" lang="ts">
  export type MenuEntry = {
    label: string;
    callback: () => void;
  };

  export type MenuContent = {
        label: string,
        entries: MenuEntry[],
    }[]

</script>

<script lang="ts">
  import { onMount } from "svelte";
  import { clickOutside } from "./click-outside-directive";

  export let content: MenuContent = [];
  let openedIndex: number = -1;

  function openOnHover(menuIndex: number) {
        if (openedIndex !== -1) {
            openedIndex = menuIndex;
        }
    }

    function closeMenu() {
        openedIndex = -1;
    }

    function toggleMenu(index: number) {
        console.log("wtf is this???", openedIndex)
        if (index == openedIndex) {
            closeMenu();
        } else {
            openedIndex = index;
        }
    }

    function entrySelected(callback: ()=>void) {
        closeMenu();
        callback();
    }

</script>

<div class="menu-row hbox" use:clickOutside={closeMenu}>
    {#each content as menu, i}
    <div class="menu">
        <button on:mouseenter={() => openOnHover(i)} on:click={() => toggleMenu(i)}>{menu.label}</button>
        {#if openedIndex==i}
        <div class="menu-body vbox">
            {#each menu.entries as entry}
            <button on:click={() => entrySelected(entry.callback)}>{entry.label}</button>
            {/each}
        </div>
        {/if}
    </div>
    {/each}
</div>

<style>
  button {
    text-align: left;
  }
  .menu-body {
    z-index: 100;
    position: absolute;
  }
</style>
