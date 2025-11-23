<script lang="ts">
	import { cn } from '$lib/utils.js';

	import { DropdownMenu as DropdownMenuPrimitive } from 'bits-ui';

	let {
		class: className,
		inset,
		ref = $bindable(null),
		...restProps
	}: DropdownMenuPrimitive.ItemProps & {
		inset?: boolean;
	} = $props();
</script>

<DropdownMenuPrimitive.Item
	bind:ref
	class={cn(
		'focus:bg-accent focus:text-accent-foreground relative flex cursor-default items-center gap-2 rounded-md px-2 py-1.5 text-sm outline-hidden transition-colors select-none disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:shrink-0',
		inset && 'pl-8',
		className
	)}
	{...restProps}
/>
