<script lang="ts">
	import type { EventHandler } from 'svelte/elements';

	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	import LoaderCircle from '@lucide/svelte/icons/loader-circle';
	import Mic from '@lucide/svelte/icons/mic';
	import Search from '@lucide/svelte/icons/search';
	import { onDestroy } from 'svelte';

	let inputValue = $state('');
	let isLoading = $state(false);
	let timer: null | ReturnType<typeof setTimeout> = $state(null);

	const handleInput: EventHandler<Event, HTMLInputElement> = () => {
		if (timer) clearTimeout(timer);

		if (!inputValue) {
			isLoading = false;
			return;
		}

		isLoading = true;
		timer = setTimeout(() => {
			isLoading = false;
			timer = null;
		}, 500);
	};

	const uid = $props.id();

	onDestroy(() => {
		if (timer) clearTimeout(timer);
	});
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Search input with loader</Label>
	<div class="relative">
		<Input
			id={uid}
			class="peer ps-9 pe-9"
			placeholder="Search..."
			type="search"
			bind:value={inputValue}
			oninput={handleInput}
		/>
		<div
			class="text-muted-foreground/80 pointer-events-none absolute inset-y-0 start-0 flex items-center justify-center ps-3 peer-disabled:opacity-50"
		>
			{#if isLoading}
				<LoaderCircle class="animate-spin" size={16} aria-hidden="true" role="presentation" />
			{:else}
				<Search size={16} aria-hidden="true" />
			{/if}
		</div>
		<button
			class="text-muted-foreground/80 ring-offset-background hover:text-foreground focus-visible:border-ring focus-visible:text-foreground focus-visible:ring-ring/30 absolute inset-y-px end-px flex h-full w-9 items-center justify-center rounded-e-lg transition-shadow focus-visible:border focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-hidden disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50"
			aria-label="Press to speak"
			type="submit"
		>
			<Mic size={16} aria-hidden="true" />
		</button>
	</div>
</div>
