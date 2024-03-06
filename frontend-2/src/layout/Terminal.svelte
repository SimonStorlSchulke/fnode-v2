<script lang="ts">
    export let text: string = "";
</script>

<div class="terminal">
    <code>
        {text}
    </code>
</div>

<style>
    .terminal {
        min-height: 200px;
        padding: 12px;
        color: #bbb;
    }
</style>