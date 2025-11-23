<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	import Mail from '@lucide/svelte/icons/mail';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Input with end icon</Label>
	<div class="relative">
		<Input id={uid} class="peer pe-9" placeholder="Email" type="email" />
		<div
			class="text-muted-foreground/80 pointer-events-none absolute inset-y-0 end-0 flex items-center justify-center pe-3 peer-disabled:opacity-50"
		>
			<Mail size={16} aria-hidden="true" />
		</div>
	</div>
</div>
