<script lang="ts" module>
	import type { WithElementRef } from 'bits-ui';
	import type { HTMLAnchorAttributes, HTMLButtonAttributes } from 'svelte/elements';

	import { tv, type VariantProps } from 'tailwind-variants';

	export const buttonVariants = tv({
		base: "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-[color,box-shadow] disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 [&_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]",
		defaultVariants: {
			size: 'default',
			variant: 'default'
		},
		variants: {
			size: {
				default: 'h-9 px-4 py-2',
				icon: 'size-9',
				lg: 'h-10 rounded-md px-8',
				sm: 'h-8 rounded-md px-3 text-xs'
			},
			variant: {
				default: 'bg-primary text-primary-foreground shadow-sm hover:bg-primary/90',
				destructive:
					'bg-destructive text-white shadow-xs hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40',
				ghost: 'hover:bg-accent hover:text-accent-foreground',
				link: 'text-primary underline-offset-4 hover:underline',
				outline:
					'border border-input bg-background shadow-xs hover:bg-accent hover:text-accent-foreground',
				secondary: 'bg-secondary text-secondary-foreground shadow-xs hover:bg-secondary/80'
			}
		}
	});

	export type ButtonVariant = VariantProps<typeof buttonVariants>['variant'];
	export type ButtonSize = VariantProps<typeof buttonVariants>['size'];

	export type ButtonProps = WithElementRef<HTMLButtonAttributes> &
		WithElementRef<HTMLAnchorAttributes> & {
			class?: string;
			size?: ButtonSize;
			variant?: ButtonVariant;
		};
</script>

<script lang="ts">
	import { cn } from '$lib/utils';

	let {
		children,
		class: className,
		href = null,
		ref = $bindable(null),
		size = 'default',
		variant = 'default',
		...restProps
	}: ButtonProps = $props();
</script>

{#if !href && restProps?.role !== 'link'}
	<button
		bind:this={ref}
		type="button"
		class={cn(buttonVariants({ className, size, variant }))}
		{...restProps}
	>
		{@render children?.()}
	</button>
{:else}
	<a bind:this={ref} {href} class={cn(buttonVariants({ className, size, variant }))} {...restProps}>
		{@render children?.()}
	</a>
{/if}
