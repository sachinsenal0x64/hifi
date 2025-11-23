<script lang="ts">
	import Label from '$lib/components/ui/label.svelte';
	import { useLocale } from '$lib/hooks/use-locale.svelte';

	import { DateField } from 'bits-ui';

	const localeCtx = useLocale();
</script>

<div class="*:not-first:mt-2">
	<Label class="text-foreground text-sm font-medium">Date input</Label>
	<DateField.Root locale={localeCtx.locale}>
		<DateField.Input
			class="border-input bg-background ring-offset-background focus-within:border-ring focus-within:ring-ring/30 relative inline-flex h-9 w-full items-center overflow-hidden rounded-lg border px-3 py-2 text-sm whitespace-nowrap shadow-xs shadow-black/[.04] transition-shadow focus-within:ring-2 focus-within:ring-offset-2 focus-within:outline-hidden disabled:opacity-50"
		>
			{#snippet children({ segments })}
				<!-- eslint-disable-next-line svelte/require-each-key -->
				{#each segments as { part, value }}
					<DateField.Segment
						{part}
						class="text-foreground focus:bg-accent focus:text-foreground aria-[valuetext=Empty]:text-muted-foreground/70 focus:aria-[valuetext=Empty]:text-foreground data-invalid:data-focused:bg-destructive data-invalid:text-destructive data-[segment=literal]:text-muted-foreground/70 data-invalid:focus:data-placeholder:text-destructive-foreground data-invalid:focus:text-destructive-foreground data-invalid:aria-[valuetext=Empty]:text-destructive inline rounded p-0.5 caret-transparent outline-0 outline-solid disabled:cursor-not-allowed disabled:opacity-50 data-[type=literal]:px-0"
					>
						{value}
					</DateField.Segment>
				{/each}
			{/snippet}
		</DateField.Input>
	</DateField.Root>
	<p class="text-muted-foreground mt-2 text-xs" role="region" aria-live="polite">
		Built with <a
			class="hover:text-foreground underline"
			href="https://next.bits-ui.com/docs/components/date-field"
			target="_blank"
			rel="noopener nofollow"
		>
			Bits UI DateField
		</a>
	</p>
</div>
