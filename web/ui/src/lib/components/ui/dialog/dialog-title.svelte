<script lang="ts">
	import { cn } from '$lib/utils.js';

	import { Dialog as DialogPrimitive } from 'bits-ui';

	let {
		class: className,
		ref = $bindable(null),
		...restProps
	}: DialogPrimitive.TitleProps = $props();
</script>

<DialogPrimitive.Title
	bind:ref
	class={cn('text-lg font-semibold tracking-tight', className)}
	{...restProps}
/>
