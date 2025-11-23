import devtoolsJson from 'vite-plugin-devtools-json';
import Icons from 'unplugin-icons/vite';
import { defineConfig } from 'vite';
import { enhancedImages } from '@sveltejs/enhanced-img';
import { sveltekit } from '@sveltejs/kit/vite';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
	plugins: [
		tailwindcss(),
		enhancedImages(),
		sveltekit(),
		Icons({ compiler: 'svelte' }),
		devtoolsJson()
	]
});
