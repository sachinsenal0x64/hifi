<script lang="ts">
	import type { Snippet } from 'svelte';

	import DialogOverlay from './dialog-overlay.svelte';
	import { cn } from '$lib/utils.js';

	import XIcon from '@lucide/svelte/icons/x';
	import { Dialog as DialogPrimitive, type WithoutChild } from 'bits-ui';

	let {
		children,
		class: className,
		portalProps,
		ref = $bindable(null),
		...restProps
	}: WithoutChild<DialogPrimitive.ContentProps> & {
		children: Snippet;
		portalProps?: DialogPrimitive.PortalProps;
	} = $props();
</script>

<DialogPrimitive.Portal {...portalProps}>
	<DialogOverlay />
	<DialogPrimitive.Content
		bind:ref
		class={cn(
			'bg-background data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 fixed top-1/2 left-1/2 z-50 grid max-h-[calc(100%-2rem)] w-full max-w-[calc(100%-2rem)] -translate-x-1/2 -translate-y-1/2 gap-4 overflow-y-auto rounded-xl border p-6 shadow-lg duration-200 sm:max-w-100',
			className
		)}
		{...restProps}
	>
		{@render children?.()}
		<DialogPrimitive.Close
			class="group cursor-pointer text-gray-200 focus-visible:border-ring focus-visible:ring-ring/50 absolute top-3 right-3 flex size-7 items-center justify-center rounded transition-[color,box-shadow] outline-none focus-visible:ring-[3px] disabled:pointer-events-none"
		>
			<XIcon size={16} class="opacity-60 transition-opacity group-hover:opacity-100" />
			<span class="sr-only">Close</span>
		</DialogPrimitive.Close>
	</DialogPrimitive.Content>
</DialogPrimitive.Portal>
