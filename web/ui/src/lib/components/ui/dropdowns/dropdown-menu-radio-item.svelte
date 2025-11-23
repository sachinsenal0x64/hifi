<script lang="ts">
	import { cn } from '$lib/utils.js';

	import Circle from '@lucide/svelte/icons/circle';
	import { DropdownMenu as DropdownMenuPrimitive, type WithoutChild } from 'bits-ui';

	let {
		children: childrenProp,
		class: className,
		ref = $bindable(null),
		...restProps
	}: WithoutChild<DropdownMenuPrimitive.RadioItemProps> = $props();
</script>

<DropdownMenuPrimitive.RadioItem
	bind:ref
	class={cn(
		'focus:bg-accent focus:text-accent-foreground relative flex cursor-default items-center rounded-md py-1.5 pr-2 pl-8 text-sm outline-hidden transition-colors select-none disabled:pointer-events-none disabled:opacity-50',
		className
	)}
	{...restProps}
>
	{#snippet children({ checked })}
		<span class="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
			{#if checked}
				<Circle class="size-2 fill-current" />
			{/if}
		</span>
		{@render childrenProp?.({ checked })}
	{/snippet}
</DropdownMenuPrimitive.RadioItem>
