<script lang="ts">
    import Menu, { type MenuContent, type MenuEntry } from '../lib-ui/menu.svelte';
  
    let resultText: string = "Please enter your name below 👇"
    let name: string
  
  
    const menuContent: MenuContent = [
      {
        label: "File",
        entries: [
          {
            label: "Save",
            callback: () => {}
          },
          {
            label: "Save as",
            callback: () => {}
          },
        ]
      },
      {
        label: "Help",
        entries: [
          {
            label: "About",
            callback: () => {}
          },
          {
            label: "Nah",
            callback: () => {}
          },
        ]
      },
    ]
  
  </script>
  
  <div class="hbox">
    <Menu content={menuContent}/>
    <div class="grow-1" style="--wails-draggable:drag"></div>
  </div>
  
  <style>
  
  
  </style>