<script lang="ts">
	import Button from '$lib/components/ui/button.svelte';
	import Checkbox from '$lib/components/ui/checkbox.svelte';
	import Label from '$lib/components/ui/label.svelte';

	import ListFilterIcon from '@lucide/svelte/icons/list-filter';
	import { Popover, PopoverContent, PopoverTrigger } from '$lib/components/ui/popover';
	const id = $props.id();
</script>

<div class="flex flex-col gap-4">
	<Popover>
		<PopoverTrigger>
			{#snippet child({ props })}
				<Button variant="outline" size="sm" class="text-sm" {...props}>
					<ListFilterIcon size={16} class="text-muted-foreground/80 -ms-1" aria-hidden="true" />
					Filters
				</Button>
			{/snippet}
		</PopoverTrigger>
		<PopoverContent class="w-36 p-3">
			<div class="space-y-3">
				<div class="text-xs font-medium">Filters</div>
				<form>
					<div class="space-y-3">
						<div class="flex items-center gap-2">
							<Checkbox id="{id}-1" />
							<Label for="{id}-1" class="font-normal">Real Time</Label>
						</div>
						<div class="flex items-center gap-2">
							<Checkbox id="{id}-2" />
							<Label for="{id}-2" class="font-normal">Top Channels</Label>
						</div>
						<div class="flex items-center gap-2">
							<Checkbox id="{id}-3" />
							<Label for="{id}-3" class="font-normal">Last Orders</Label>
						</div>
						<div class="flex items-center gap-2">
							<Checkbox id="{id}-4" />
							<Label for="{id}-4" class="font-normal">Total Spent</Label>
						</div>
					</div>
					<div
						role="separator"
						aria-orientation="horizontal"
						class="bg-border -mx-3 my-3 h-px"
					></div>
					<div class="flex justify-between gap-2">
						<Button size="sm" variant="outline" class="h-7 px-2">Clear</Button>
						<Button size="sm" class="h-7 px-2">Apply</Button>
					</div>
				</form>
			</div>
		</PopoverContent>
	</Popover>
</div>
