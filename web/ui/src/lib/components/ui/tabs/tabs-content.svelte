<script lang="ts">
	import { cn } from '$lib/utils.js';

	import { Tabs as TabsPrimitive } from 'bits-ui';

	let {
		class: className,
		ref = $bindable(null),
		value,
		...restProps
	}: TabsPrimitive.ContentProps = $props();
</script>

<TabsPrimitive.Content
	bind:ref
	class={cn('flex-1 outline-hidden', className)}
	{value}
	{...restProps}
/>
