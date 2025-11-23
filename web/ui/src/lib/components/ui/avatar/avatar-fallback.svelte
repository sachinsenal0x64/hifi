<script lang="ts">
	import { cn } from '$lib/utils.js';

	import { Avatar as AvatarPrimitive } from 'bits-ui';

	let {
		class: className,
		ref = $bindable(null),
		...restProps
	}: AvatarPrimitive.FallbackProps = $props();
</script>

<AvatarPrimitive.Fallback
	bind:ref
	class={cn(
		'bg-secondary flex size-full items-center justify-center rounded-[inherit] text-xs',
		className
	)}
	{...restProps}
/>
