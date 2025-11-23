<script lang="ts">
	import Label from '$lib/components/ui/label.svelte';

	import ClockIcon from '@lucide/svelte/icons/clock';
	import { TimeField, TimeInput } from '$lib/components/ui/datefield-rac';
</script>

<TimeField class="*:not-first:mt-2">
	<Label class="text-foreground text-sm font-medium">Time input with end icon</Label>
	<div class="relative">
		<TimeInput />
		<div
			class="text-muted-foreground/80 pointer-events-none absolute inset-y-0 end-0 z-10 flex items-center justify-center pe-3"
		>
			<ClockIcon size={16} aria-hidden="true" />
		</div>
	</div>
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
