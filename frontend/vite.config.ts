import {defineConfig} from 'vite'
import {svelte} from '@sveltejs/vite-plugin-svelte'
import path from "path";

export default defineConfig({
    plugins: [svelte()],
    server: {
        host: '0.0.0.0',
    },
    base: "/",
    resolve: {
        alias: {
            "$": path.resolve("./src"),
        },
    },
    css: {
        preprocessorOptions: {
            scss: {
                api: 'modern'
            }
        }
    }
})
