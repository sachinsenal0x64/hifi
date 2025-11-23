<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	import { cn } from '$lib/utils.js';

	import { type WithElementRef } from 'bits-ui';

	let {
		children,
		class: className,
		ref = $bindable(null),
		...restProps
	}: WithElementRef<HTMLAttributes<HTMLSpanElement>> = $props();
</script>

<span
	bind:this={ref}
	class={cn(
		'border-border bg-background text-muted-foreground/70 ms-auto -me-1 inline-flex h-5 max-h-full items-center rounded border px-1 font-[inherit] text-[0.625rem] font-medium',
		className
	)}
	{...restProps}
>
	{@render children?.()}
</span>
