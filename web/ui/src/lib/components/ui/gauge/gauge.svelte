<script lang="ts">
	import { cn } from '$lib/utils';
	import { Tween } from 'svelte/motion';

	type Props = {
		class?: string | undefined;
		value: number;
		total?: number;
		size?: 'tiny' | 'sm' | 'md' | 'lg';
		show_value?: boolean;
		colors?: Record<number, string> | { primary: string; secondary: string } | undefined;
	};

	let {
		class: class_name = undefined,
		value,
		total = 100,
		size = 'md',
		show_value = false,
		colors = undefined
	}: Props = $props();

	const sizes_map = {
		tiny: { size: 20, stroke_width: 15 },
		sm: { size: 32, stroke_width: 10 },
		md: { size: 64, stroke_width: 10 },
		lg: { size: 128, stroke_width: 10 }
	};

	const arc_size = sizes_map[size].size;
	const stroke_width = sizes_map[size].stroke_width;
	// Store to track the length of the arc
	function get_radius() {
		switch (size) {
			case 'tiny':
				return 42.5;
			case 'sm':
			case 'md':
			case 'lg':
				return 45;
		}
	}
	const radius = get_radius();
	const circumference = 2 * Math.PI * radius;
	let percent_px = circumference / 100;
	let circle_size = 100;
	const progress = new Tween(0, { duration: 1000 });

	$effect(() => {
		progress.target = (value / total) * 100;
	});

	// Colors
	function get_primary_arc_stroke_fill() {
		// if colors is an object with primary key, use only that
		if (colors && 'primary' in colors) {
			return colors.primary;
		} else if (colors) {
			// if colors is an object with a range of percentages, use those
			const percentages = Object.keys(colors)
				.map(Number)
				.sort((a, b) => a - b);
			for (let i = percentages.length - 1; i >= 0; i--) {
				if (value >= percentages[i]) {
					return colors[percentages[i]];
				}
			}
		}
		// if no colors are provided, use the default color scale
		if (value < 33) {
			return 'stroke-red-700';
		} else if (value < 66) {
			return 'stroke-amber-700';
		} else {
			return 'stroke-green-700';
		}
	}
	function get_secondary_arc_stroke_fill() {
		// if colors is an object with secondary key, use only that
		if (colors && 'secondary' in colors) {
			return colors.secondary;
		}
		// if no colors are provided, use the default color scale
		return 'stroke-gray-400';
	}
</script>

<div
	class={cn('relative grid place-items-center', class_name)}
	style:--circle-size={circle_size}
	style:--circumference={circumference}
	style:--percent-to-px="{percent_px}px"
	style:--gap-percent={value === 0 ? 0 : 5}
	style:--offset-factor="0"
	style:--transition-length="1s"
	style:--transition-step="200ms"
	style:--delay="0s"
	style:--percent-to-deg="3.6deg"
	style="transform: translateZ(0);"
>
	<svg
		fill="none"
		height={arc_size}
		width={arc_size}
		stroke-width="2"
		viewBox="0 0 {circle_size} {circle_size}"
		class={cn(class_name)}
	>
		<circle
			cx={50}
			cy={50}
			r={radius}
			stroke-width={stroke_width}
			stroke-dashoffset="0"
			stroke-linecap="round"
			stroke-linejoin="round"
			class={cn(get_secondary_arc_stroke_fill())}
			style="
        			stroke-dasharray: calc(var(--stroke-percent) * var(--percent-to-px)) var(--circumference);
        			transform: rotate(calc(1turn - 90deg - (var(--gap-percent) * var(--percent-to-deg) * var(--offset-factor-secondary)))) scaleY(-1);
        			transition: all var(--transition-length) ease var(--delay);
        			transform-origin:calc(var(--circle-size) / 2) calc(var(--circle-size) / 2);
        		"
			style:--stroke-percent={value === 0 ? 100 - progress.current : 90 - progress.current}
			style:--offset-factor-secondary="calc(1 - var(--offset-factor))"
		/>
		{#if value > 0}
			<circle
				cx={50}
				cy={50}
				r={radius}
				stroke-width={stroke_width}
				stroke-dashoffset="0"
				stroke-linecap="round"
				stroke-linejoin="round"
				class={cn(get_primary_arc_stroke_fill())}
				style="
						stroke-dasharray:calc(var(--stroke-percent) * var(--percent-to-px)) var(--circumference);
        				transition:var(--transition-length) ease var(--delay),stroke var(--transition-length) ease var(--delay);
        				transition-property: stroke-dasharray,transform;
        				transform:rotate(calc(-90deg + var(--gap-percent) * var(--offset-factor) * var(--percent-to-deg)));
        				transform-origin:calc(var(--circle-size) / 2) calc(var(--circle-size) / 2);
					"
				style:--stroke-percent={progress.current}
			/>
		{/if}
	</svg>

	{#if show_value && size !== 'tiny'}
		<p
			data-current-value={progress.current}
			class={cn('absolute tabular-nums ease-linear', {
				'text-[0.6875rem] font-medium': size === 'sm',
				'text-lg leading-6 font-medium': size === 'md',
				'text-[2rem] leading-10 font-semibold': size === 'lg'
			})}
		>
			{Math.round(progress.current)}
		</p>
	{/if}
</div>
