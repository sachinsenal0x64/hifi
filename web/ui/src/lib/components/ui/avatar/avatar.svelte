<script lang="ts">
	import { cn } from '$lib/utils.js';

	import { Avatar as AvatarPrimitive } from 'bits-ui';

	let {
		class: className,
		ref = $bindable(null),
		...restProps
	}: AvatarPrimitive.RootProps = $props();
</script>

<AvatarPrimitive.Root
	bind:ref
	class={cn('relative flex size-8 shrink-0 overflow-hidden rounded-full', className)}
	{...restProps}
/>
