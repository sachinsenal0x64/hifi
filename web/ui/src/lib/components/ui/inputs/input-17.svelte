<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';

	import ChevronDown from '@lucide/svelte/icons/chevron-down';

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Input with start select</Label>
	<div class="flex rounded-lg shadow-xs shadow-black/[.04]">
		<div class="relative">
			<select
				class="peer border-input bg-background text-muted-foreground ring-offset-background hover:bg-accent hover:text-foreground focus-visible:border-ring focus-visible:text-foreground focus-visible:ring-ring/30 inline-flex h-full appearance-none items-center rounded-s-lg border ps-3 pe-8 text-sm transition-shadow focus:z-10 focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-hidden disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50"
				aria-label="Protocol"
			>
				<option value="https://">https://</option>
				<option value="http://">http://</option>
				<option value="ftp://">ftp://</option>
				<option value="sftp://">sftp://</option>
				<option value="ws://">ws://</option>
				<option value="wss://">wss://</option>
			</select>
			<span
				class="text-muted-foreground/80 pointer-events-none absolute inset-y-0 end-px flex h-full w-9 items-center justify-center peer-disabled:opacity-50"
			>
				<ChevronDown size={16} aria-hidden="true" />
			</span>
		</div>
		<Input
			id={uid}
			class="-ms-px rounded-s-none shadow-none focus-visible:z-10"
			placeholder="192.168.1.1"
			type="text"
		/>
	</div>
</div>
