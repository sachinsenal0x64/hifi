<script lang="ts">
	import { cn } from '$lib/utils';
	import { NavigationMenu as NavigationMenuPrimitive } from 'bits-ui';

	let {
		class: className,
		ref = $bindable(null),
		...restProps
	}: NavigationMenuPrimitive.ContentProps = $props();
</script>

<NavigationMenuPrimitive.Content
	bind:ref
	data-slot="navigation-menu-content"
	class={cn(
		'data-[motion^=from-]:animate-in data-[motion^=to-]:animate-out data-[motion^=from-]:fade-in data-[motion^=to-]:fade-out data-[motion=from-end]:slide-in-from-right-52 data-[motion=from-start]:slide-in-from-left-52 data-[motion=to-end]:slide-out-to-right-52 data-[motion=to-start]:slide-out-to-left-52 top-0 left-0 w-full p-2 pr-2.5 md:absolute md:w-auto',
		'[&[data-slot="navigation-menu-link"]:focus]:outline-hidden',
		'group-data-[viewport=false]/navigation-menu:bg-popover group-data-[viewport=false]/navigation-menu:text-popover-foreground data-[state=open]:group-data-[viewport=false]/navigation-menu:animate-in data-[state=closed]:group-data-[viewport=false]/navigation-menu:animate-out data-[state=closed]:group-data-[viewport=false]/navigation-menu:fade-out-0 data-[state=open]:group-data-[viewport=false]/navigation-menu:fade-in-0 data-[state=closed]:group-data-[viewport=false]/navigation-menu:zoom-out-95 data-[state=open]:group-data-[viewport=false]/navigation-menu:zoom-in-95 group-data-[viewport=false]/navigation-menu:top-full group-data-[viewport=false]/navigation-menu:mt-1.5 group-data-[viewport=false]/navigation-menu:overflow-hidden group-data-[viewport=false]/navigation-menu:rounded-md group-data-[viewport=false]/navigation-menu:border group-data-[viewport=false]/navigation-menu:shadow-sm group-data-[viewport=false]/navigation-menu:duration-200',
		className
	)}
	{...restProps}
/>
