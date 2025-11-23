<script lang="ts">
	import type { Snippet } from 'svelte';

	import { cn } from '$lib/utils.js';

	import Check from '@lucide/svelte/icons/check';
	import Minus from '@lucide/svelte/icons/minus';
	import { DropdownMenu as DropdownMenuPrimitive, type WithoutChildrenOrChild } from 'bits-ui';

	let {
		checked = $bindable(false),
		children: childrenProp,
		class: className,
		ref = $bindable(null),
		...restProps
	}: WithoutChildrenOrChild<DropdownMenuPrimitive.CheckboxItemProps> & {
		children?: Snippet;
	} = $props();
</script>

<DropdownMenuPrimitive.CheckboxItem
	bind:ref
	bind:checked
	class={cn(
		'focus:bg-accent focus:text-accent-foreground relative flex cursor-default items-center rounded-md py-1.5 pr-2 pl-8 text-sm outline-hidden transition-colors select-none disabled:pointer-events-none disabled:opacity-50',
		className
	)}
	{...restProps}
>
	{#snippet children({ checked, indeterminate })}
		<span class="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
			{#if indeterminate}
				<Minus class="size-4" />
			{:else}
				<Check class={cn('size-4', !checked && 'text-transparent')} />
			{/if}
		</span>
		{@render childrenProp?.()}
	{/snippet}
</DropdownMenuPrimitive.CheckboxItem>
