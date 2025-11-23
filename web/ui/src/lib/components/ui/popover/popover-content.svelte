<script lang="ts">
	import type { Snippet } from 'svelte';

	import { cn } from '$lib/utils.js';

	import { Popover as PopoverPrimitive, type WithoutChildren } from 'bits-ui';

	let {
		align = 'center',
		children,
		class: className,
		portalProps,
		ref = $bindable(null),
		showArrow = false,
		sideOffset = 4,
		...restProps
	}: WithoutChildren<PopoverPrimitive.ContentProps> & {
		children: Snippet;
		portalProps?: PopoverPrimitive.PortalProps;
		showArrow?: boolean;
	} = $props();
</script>

<PopoverPrimitive.Portal {...portalProps}>
	<PopoverPrimitive.Content
		bind:ref
		{sideOffset}
		{align}
		class={cn(
			'border-border bg-popover text-popover-foreground data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 z-50 max-h-(--bits-popover-content-available-height) min-w-32 overflow-x-hidden overflow-y-auto rounded-lg border p-4 shadow-lg shadow-black/5 outline-hidden',
			className
		)}
		{...restProps}
	>
		{#if showArrow}
			<PopoverPrimitive.Arrow
				class="text-popover -my-px drop-shadow-[0_1px_0_hsl(var(--border))]"
			/>
		{/if}
		{@render children()}
	</PopoverPrimitive.Content>
</PopoverPrimitive.Portal>
