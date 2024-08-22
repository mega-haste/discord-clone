<script lang="ts">
    import Icon from "@iconify/svelte";
    import type { PageData } from "./$types";
    import { Message, type Models } from "@lib/typedefs";
    import MessageLayout from "@lib/MessageLayout.svelte";
    import { onMount } from "svelte";
    import io from "../../webSocketConnection";
    import axios from "axios";
	import { useQuery } from "@sveltestack/svelte-query"
    import * as globals from "@global";

    let theMessage: string = "";
    let cachedMessages: Message[] = [];

    function sendMessage() {
        if (theMessage.length <= 0) return;
        cachedMessages.push(new Message(theMessage, 0, true, undefined));
        io.broadcast().emit("chat-send", theMessage);
        theMessage = "";
        cachedMessages = cachedMessages;
    }

    function onSendHandler(ev: MouseEvent) {
        ev.preventDefault();
        sendMessage();
    }

    function onInputEnter(ev: KeyboardEvent) {
        if (ev.key === "Enter") {
            if (theMessage.length <= 0) return;
            sendMessage();
        }
    }

    onMount(async () => {
        io.on("receive-message", (message) => {
            cachedMessages.push(
                new Message(message as any, 0, true, undefined)
            );
            cachedMessages = cachedMessages;
        });
    });

	const localMessagesQuery = useQuery("global-messgaes", () =>
		axios({
			url: new URL("api/global/messages", globals.serverHost).href,
			method: "GET",
		})
			.then(res => {
				cachedMessages = res.data;
				return res.data;
			})
	);

    export let data: PageData;
</script>

<div class="flex flex-nowrap flex-col justify-between h-full py-3">
    <div class="h-full">
		{#if $localMessagesQuery.isLoading}
			<div>
				Wait our global messages are loading rn
			</div>
		{:else if $localMessagesQuery.isError}
			<span>An error has occurred: {$localMessagesQuery.error.message}</span>
		{:else}
			{#each cachedMessages as message, count}
				<MessageLayout {count} {message} />
			{/each}
		{/if}
    </div>
    <div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
        <div class="input-group-shim">
            <Icon icon="material-symbols:business-messages" />
        </div>
        <input
            type="text"
            placeholder="Message..."
            on:keyup={onInputEnter}
            bind:value={theMessage}
        />
        <button
            type="button"
            class="btn variant-filled-secondary rounded-l-none"
            on:click={onSendHandler}
        >
            send
        </button>
    </div>
</div>
