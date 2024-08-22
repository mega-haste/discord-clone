import { sveltekit } from "@sveltejs/kit/vite";
import path from "path";
import { defineConfig } from "vite";

export default defineConfig({
    resolve: {
        alias: {
            "@src": path.resolve(__dirname, "./src"),
            "@lib": path.resolve(__dirname, "./src/lib"),
            "@global": path.resolve(__dirname, "./src/globals.ts"),
			"@static": path.resolve(__dirname, "./static")
        },
    },
    plugins: [sveltekit()],
});
