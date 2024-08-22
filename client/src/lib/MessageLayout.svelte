<script lang="ts">
    import type { Message } from "./typedefs";

    export let message: Message;
    export let count: number;
</script>

<div class="{count % 2 === 0 ? 'bg-slate-500' : ''} px-5 py-1 rounded">
    {#if !message.fromMe}
        <h3>From {message.fromWho}:</h3>
        <p>{message.content}</p>
    {:else}
        <p>{message.content}</p>
    {/if}
</div>
