<script lang="ts" module>
	import type { ToggleRootProps } from 'bits-ui';

	import { tv, type VariantProps } from 'tailwind-variants';

	export const toggleVariants = tv({
		base: "inline-flex items-center justify-center gap-2 rounded-md text-sm font-medium transition-[color,box-shadow] hover:bg-muted hover:text-muted-foreground disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 [&_svg]:shrink-0 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] outline-none aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive",
		defaultVariants: {
			size: 'default',
			variant: 'default'
		},
		variants: {
			size: {
				default: 'h-9 px-3',
				lg: 'h-10 px-3',
				sm: 'h-8 px-2'
			},
			variant: {
				default: 'bg-transparent',
				outline:
					'border border-input bg-transparent shadow-xs hover:bg-accent hover:text-accent-foreground'
			}
		}
	});

	export type ToggleVariant = VariantProps<typeof toggleVariants>['variant'];
	export type ToggleSize = VariantProps<typeof toggleVariants>['size'];
	export type ToggleVariants = VariantProps<typeof toggleVariants>;
	export type ToggleProps = ToggleRootProps & {
		class?: string;
		size?: ToggleSize;
		variant?: ToggleVariant;
	};
</script>

<script lang="ts">
	import { cn } from '$lib/utils.js';

	import { Toggle as TogglePrimitive } from 'bits-ui';

	let {
		class: className,
		pressed = $bindable(false),
		ref = $bindable(null),
		size = 'default',
		variant = 'default',
		...restProps
	}: ToggleProps = $props();
</script>

<TogglePrimitive.Root
	bind:ref
	bind:pressed
	class={cn(toggleVariants({ className, size, variant }))}
	{...restProps}
/>
