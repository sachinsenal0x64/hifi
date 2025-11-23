<script lang="ts">
	import { cn } from '$lib/utils';
	import { NavigationMenu as NavigationMenuPrimitive } from 'bits-ui';

	let {
		class: className,
		ref = $bindable(null),
		...restProps
	}: NavigationMenuPrimitive.ListProps = $props();
</script>

<NavigationMenuPrimitive.List
	bind:ref
	data-slot="navigation-menu-list"
	class={cn('group flex flex-1 list-none items-center justify-center gap-1', className)}
	{...restProps}
/>
