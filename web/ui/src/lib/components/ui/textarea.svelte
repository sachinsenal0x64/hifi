<script lang="ts">
	import type { WithElementRef } from 'bits-ui';
	import type { HTMLTextareaAttributes } from 'svelte/elements';

	import { cn } from '$lib/utils.js';

	let {
		class: className,
		ref = $bindable(null),
		value = $bindable(),
		...restProps
	}: WithElementRef<HTMLTextareaAttributes> = $props();
</script>

<textarea
	bind:this={ref}
	bind:value
	class={cn(
		'border-input placeholder:text-muted-foreground/70 focus-visible:border-ring focus-visible:ring-ring/50 aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive flex min-h-19.5 w-full rounded-md border bg-transparent px-3 py-2 text-sm shadow-xs transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-50',
		className
	)}
	{...restProps}
></textarea>
