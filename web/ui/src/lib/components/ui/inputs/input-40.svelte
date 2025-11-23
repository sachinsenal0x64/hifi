<script lang="ts">
	import Label from '$lib/components/ui/label.svelte';

	import { DateField, DateInput } from '$lib/components/ui/datefield-rac';
</script>

<DateField class="*:not-first:mt-2" granularity="minute" hourCycle={24}>
	<Label class="text-foreground text-sm font-medium">Date and time input</Label>
	<DateInput />
	<p class="text-muted-foreground mt-2 text-xs" role="region" aria-live="polite">
		Built with
		<a
			class="hover:text-foreground underline"
			href="https://bits-ui.com/docs/components/date-field"
			target="_blank"
			rel="noopener nofollow"
		>
			Bits UI
		</a>
	</p>
</DateField>
