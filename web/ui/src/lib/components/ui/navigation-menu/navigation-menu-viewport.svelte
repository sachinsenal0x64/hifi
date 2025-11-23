<script lang="ts">
	import { cn } from '$lib/utils';
	import { NavigationMenu as NavigationMenuPrimitive } from 'bits-ui';

	let {
		class: className,
		ref = $bindable(null),
		...restProps
	}: NavigationMenuPrimitive.ViewportProps = $props();
</script>

<div class={cn('absolute top-full left-0 isolate z-50 flex justify-center')}>
	<NavigationMenuPrimitive.Viewport
		bind:ref
		data-slot="navigation-menu-viewport"
		class={cn(
			'origin-top-center bg-popover text-popover-foreground data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-90 relative mt-1.5 h-(--bits-navigation-menu-viewport-height) w-full overflow-hidden rounded-md border shadow-sm md:w-(--bits-navigation-menu-viewport-width)',

			className
		)}
		{...restProps}
	/>
</div>
