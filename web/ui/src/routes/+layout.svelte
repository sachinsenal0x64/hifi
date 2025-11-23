<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import '@fontsource-variable/inter';
	import { Toaster } from 'svelte-sonner';

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<Toaster closeButton position="top-center" />

{@render children?.()}
