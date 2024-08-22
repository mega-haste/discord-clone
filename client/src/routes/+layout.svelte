<script lang="ts">
    import "../app.css";
	import { QueryClient, QueryClientProvider, useQuery } from "@sveltestack/svelte-query";
    import { onMount } from "svelte";
    import io from "../webSocketConnection";

	const queryClient = new QueryClient();

    let socketLoaded = false;

    onMount(() => {
        io.onConnect(() => {
            console.log("connected as:", io.ID);
            socketLoaded = true;
        });
        io.onClose(() => {
            console.warn("socket closed");
        });
    });
</script>

<QueryClientProvider client={queryClient}>
	{#if socketLoaded}
		<slot/>
	{:else}
		<div>Loading websocket...</div>
	{/if}
</QueryClientProvider>
