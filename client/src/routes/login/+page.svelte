<script lang="ts">
import { onMount } from "svelte";
import type { PageData } from "./$types";
import * as globals from "@global";
import { setCookie, getAllCookies } from "@src/lib/cookies";

let userName = "";
let password = "";

async function submiting() {
	if (userName === "") return;
	if (password === "") return;
	try {
		const loginReq = await fetch(new URL("/api/login", globals.serverHost).href, {
			method: "POST",
			body: JSON.stringify({ userName, password })
		});
		const res: {
			"status": string,
			"msg": string,
			"token": string,
		} = await loginReq.json();
		if (res.token) {
			setCookie("login_token", res.token, 20);
		}
		console.log(getAllCookies());
	} catch (e) {
		console.error(e);
	}
}

export let data: PageData;
</script>

<div class="flex flex-row items-center justify-center bg-white h-full">
	<div class="inline-block w-fit p-5 shadow-xl rounded-md">
		<h2 class="h2 px-4 py-2 text-black text-3xl font-semibold">
			Welcome Back
		</h2>
		<form class="flex gap-1 flex-col mt-5" on:submit|preventDefault={submiting}>
			<input bind:value={userName} type="text" placeholder="User name" class="input !w-full"/>
			<input bind:value={password} type="password" placeholder="Password" class="input !w-full"/>
			<div class="flex mt-2 gap-1">
				<button type="submit" class="inline-flex items-center justify-center rounded-xl bg-blue-600 shadow-[0px_0.25rem_blue] transition px-4 py-3 font-medium leading-none text-white shadow hover:opacity-75 active:translate-y-1 active:shadow-[0px_0px_blue]">Submit</button>
				<a href="#" type="submit" class="inline-flex items-center justify-center rounded-xl bg-gray-200 shadow-[0px_0.25rem_#ccc] transition px-4 py-3 font-medium leading-none text-black shadow hover:opacity-75 active:translate-y-1 active:shadow-[0px_0px_grey]">Create acount</a>
			</div>
		</form>
	</div>
</div>

<style scoped>
@tailwind base;
@tailwind utilities;

@layer utilities {
	.input {
		@apply bg-gray-50 border border-gray-300 
			text-gray-900 text-sm 
			rounded-lg 
			focus:ring-blue-500 focus:border-blue-500
			block w-full p-2.5 
			dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500;
	}
}
@layer base {
	section {
		@apply w-full;
	}
}
</style>

