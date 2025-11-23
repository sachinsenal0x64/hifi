<script lang="ts">
	import NavigationMenuViewport from './navigation-menu-viewport.svelte';

	import { cn } from '$lib/utils';
	import { NavigationMenu as NavigationMenuPrimitive } from 'bits-ui';

	let {
		children,
		class: className,
		ref = $bindable(null),
		viewport = true,
		...restProps
	}: NavigationMenuPrimitive.RootProps & {
		viewport?: boolean;
	} = $props();
</script>

<NavigationMenuPrimitive.Root
	bind:ref
	data-slot="navigation-menu"
	data-viewport={viewport}
	class={cn(
		'group/navigation-menu relative flex max-w-max flex-1 items-center justify-center',
		className
	)}
	{...restProps}
>
	{@render children?.()}

	{#if viewport}
		<NavigationMenuViewport />
	{/if}
</NavigationMenuPrimitive.Root>
