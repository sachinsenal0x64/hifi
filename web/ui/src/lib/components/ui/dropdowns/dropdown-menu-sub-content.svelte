<script lang="ts">
	import { cn } from '$lib/utils.js';

	import { DropdownMenu as DropdownMenuPrimitive } from 'bits-ui';

	let {
		class: className,
		ref = $bindable(null),
		...restProps
	}: DropdownMenuPrimitive.SubContentProps = $props();
</script>

<DropdownMenuPrimitive.SubContent
	bind:ref
	class={cn(
		'border-border bg-popover text-popover-foreground data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 z-50 min-w-40 overflow-hidden rounded-md border p-1 shadow-lg shadow-black/5',
		className
	)}
	{...restProps}
/>
