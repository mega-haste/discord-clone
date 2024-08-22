<script lang="ts">
    import { useQuery } from "@sveltestack/svelte-query";
	import * as globals from "@global";
	import { getCookie } from "@src/lib/cookies";
	import { type Models, type ApiResult } from "@lib/typedefs";

	let loginQuery = useQuery("login-query", () =>
			fetch(new URL(`/api/login?login_token=${getCookie("login_token")}`, globals.serverHost))
				.then((res) => res.json() as Promise<ApiResult | Models.UserModel>));

</script>

{#if $loginQuery.isLoading}
	<p>Loging in</p>
{:else if $loginQuery.isError}
	<pre class="block font-mono whitespace-pre">Error while loging in: {JSON.stringify($loginQuery.error, undefined, 4)}</pre>
{:else if $loginQuery.isSuccess}
	<pre class="block font-mono whitespace-pre">Logged in: {JSON.stringify($loginQuery.data, undefined, 4)}</pre>
{/if}

Home
