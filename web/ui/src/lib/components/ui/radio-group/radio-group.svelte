<script lang="ts" module>
	import { RadioGroup as RadioGroupPrimitive } from 'bits-ui';

	export const RadioGroupItem = RadioGroupPrimitive.Item;
</script>

<script lang="ts">
	import { cn } from '$lib/utils.js';

	let {
		children,
		class: className,
		ref = $bindable(null),
		value = $bindable(undefined),
		...restProps
	}: RadioGroupPrimitive.RootProps = $props();
</script>

<RadioGroupPrimitive.Root bind:ref bind:value class={cn('grid gap-3', className)} {...restProps}>
	{@render children?.()}
</RadioGroupPrimitive.Root>
