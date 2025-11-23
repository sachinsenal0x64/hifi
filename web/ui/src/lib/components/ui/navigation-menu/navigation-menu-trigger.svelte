<script lang="ts" module>
	import { cn } from '$lib/utils';
	import { tv } from 'tailwind-variants';

	export const navigationMenuTriggerStyle = tv({
		base: 'group inline-flex h-9 w-max items-center justify-center rounded-md bg-background px-4 py-2 text-sm font-medium hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=open]:hover:bg-accent data-[state=open]:text-accent-foreground data-[state=open]:focus:bg-accent data-[state=open]:bg-accent/50 focus-visible:ring-ring/50 outline-hidden transition-[color,box-shadow] focus-visible:ring-[3px] focus-visible:outline-1'
	});
</script>

<script lang="ts">
	import ChevronDownIcon from '@lucide/svelte/icons/chevron-down';
	import { NavigationMenu as NavigationMenuPrimitive } from 'bits-ui';

	let {
		children,
		class: className,
		ref = $bindable(null),
		...restProps
	}: NavigationMenuPrimitive.TriggerProps = $props();
</script>

<NavigationMenuPrimitive.Trigger
	bind:ref
	data-slot="navigation-menu-trigger"
	class={cn(navigationMenuTriggerStyle(), 'group', className)}
	{...restProps}
>
	{@render children?.()}

	<ChevronDownIcon
		class="relative top-px ml-1 size-3 transition duration-300 group-data-[state=open]:rotate-180"
		aria-hidden="true"
	/>
</NavigationMenuPrimitive.Trigger>
