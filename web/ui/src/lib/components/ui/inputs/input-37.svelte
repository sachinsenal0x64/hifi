<script lang="ts">
	import Label from '$lib/components/ui/label.svelte';

	import { TimeField, TimeInput } from '$lib/components/ui/datefield-rac';
</script>

<TimeField class="*:not-first:mt-2">
	<Label class="text-foreground text-sm font-medium">Time input</Label>
	<TimeInput />
	<p class="text-muted-foreground mt-2 text-xs" role="region" aria-live="polite">
		Built with
		<a
			class="hover:text-foreground underline"
			href="https://bits-ui.com/docs/components/time-field"
			target="_blank"
			rel="noopener nofollow"
		>
			Bits UI
		</a>
	</p>
</TimeField>
