<script lang="ts">
	import { cn } from '$lib/utils';
	import { NavigationMenu as NavigationMenuPrimitive } from 'bits-ui';

	let {
		class: className,
		ref = $bindable(null),
		...restProps
	}: NavigationMenuPrimitive.LinkProps = $props();
</script>

<NavigationMenuPrimitive.Link
	bind:ref
	data-slot="navigation-menu-link"
	class={cn(
		"hover:bg-accent focus:bg-accent focus:text-accent-foreground focus-visible:ring-ring/50 data-active:bg-accent data-active:text-accent-foreground data-active:hover:bg-accent data-active:focus:bg-accent [&_svg:not([class*='text-'])]:text-muted-foreground flex flex-col gap-1 rounded-sm p-2 text-sm outline-hidden transition-all focus-visible:ring-[3px] focus-visible:outline-1 [&_svg:not([class*='size-'])]:size-4",
		className
	)}
	{...restProps}
/>
