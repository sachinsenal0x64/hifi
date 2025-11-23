<script lang="ts" module>
	import type { WithElementRef } from 'bits-ui';
	import type { HTMLAttributes } from 'svelte/elements';

	import { tv, type VariantProps } from 'tailwind-variants';

	export const badgeVariants = tv({
		base: 'inline-flex items-center justify-center rounded-full border px-1.5 text-xs font-medium w-fit whitespace-nowrap shrink-0 gap-1 [&>svg]:pointer-events-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] transition-[color,box-shadow] [&>svg]:shrink-0 leading-normal',
		defaultVariants: {
			variant: 'default'
		},
		variants: {
			variant: {
				default: 'border-transparent bg-primary text-primary-foreground [a&]:hover:bg-primary/90',
				destructive:
					'border-transparent bg-destructive text-white [a&]:hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40',
				outline: 'text-foreground [a&]:hover:bg-accent [a&]:hover:text-accent-foreground',
				secondary:
					'border-transparent bg-secondary text-secondary-foreground [a&]:hover:bg-secondary/90'
			}
		}
	});

	export type BadgeVariant = VariantProps<typeof badgeVariants>['variant'];

	export type BadgeProps = WithElementRef<HTMLAttributes<HTMLDivElement>, HTMLDivElement> & {
		class?: string;
		variant?: BadgeVariant;
	};
</script>

<script lang="ts">
	import { cn } from '$lib/utils.js';

	let {
		children,
		class: className,
		ref = $bindable(null),
		variant = 'default',
		...restProps
	}: BadgeProps = $props();
</script>

<div class={cn(badgeVariants({ className, variant }))} bind:this={ref} {...restProps}>
	{@render children?.()}
</div>
