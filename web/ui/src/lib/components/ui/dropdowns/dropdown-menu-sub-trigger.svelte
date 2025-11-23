<script lang="ts">
	import { cn } from '$lib/utils.js';

	import ChevronRight from '@lucide/svelte/icons/chevron-right';
	import { DropdownMenu as DropdownMenuPrimitive } from 'bits-ui';

	let {
		children,
		class: className,
		inset,
		ref = $bindable(null),
		...restProps
	}: DropdownMenuPrimitive.SubTriggerProps & {
		inset?: boolean;
	} = $props();
</script>

<DropdownMenuPrimitive.SubTrigger
	bind:ref
	class={cn(
		'focus:bg-accent data-[state=open]:bg-accent flex cursor-default items-center gap-2 rounded-md px-2 py-1.5 text-sm outline-hidden select-none [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0',
		inset && 'pl-8',
		className
	)}
	{...restProps}
>
	{@render children?.()}
	<ChevronRight class="text-muted-foreground/80 ml-auto" />
</DropdownMenuPrimitive.SubTrigger>
