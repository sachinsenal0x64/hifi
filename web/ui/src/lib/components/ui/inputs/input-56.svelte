<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	import Inputmask from 'inputmask';

	let inputElement = $state<HTMLInputElement | null>(null);

	$effect(() => {
		if (!inputElement) return;
		const im = new Inputmask('99:99:99', {
			placeholder: '-',
			showMaskOnHover: false
		}).mask(inputElement);

		return () => im.remove();
	});

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Timestamp</Label>
	<Input id={uid} placeholder="00:00:00" type="text" bind:ref={inputElement} />
	<p class="text-muted-foreground mt-2 text-xs" role="region" aria-live="polite">
		Built with <a
			class="hover:text-foreground underline"
			href="https://github.com/RobinHerbots/inputmask"
			target="_blank"
			rel="noopener nofollow">inputmask</a
		>
	</p>
</div>
