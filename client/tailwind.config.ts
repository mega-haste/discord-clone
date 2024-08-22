import { join } from "path";
import type { Config } from "tailwindcss";

// plugins
import forms from "@tailwindcss/forms";

const config = {
    darkMode: "class",
    content: [
        "./src/**/*.{html,js,svelte,ts,jsx,tsx}",
    ],
    theme: {
        extend: {
		},
    },
    plugins: [
        forms,
    ],
} satisfies Config;

export default config;
