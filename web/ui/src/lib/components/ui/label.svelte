<script lang="ts">
	import { cn } from '$lib/utils.js';

	import { Label as LabelPrimitive } from 'bits-ui';

	let {
		class: className,
		ref = $bindable(null),
		...restProps
	}: LabelPrimitive.RootProps = $props();
</script>

<LabelPrimitive.Root
	bind:ref
	class={cn('text-foreground text-sm font-medium', className)}
	{...restProps}
/>
