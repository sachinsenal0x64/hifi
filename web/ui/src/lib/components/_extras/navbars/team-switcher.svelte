<script lang="ts">
	import Button from '$lib/components/ui/button.svelte';

	import ChevronUpIcon from '@lucide/svelte/icons/chevrons-up';
	import {
		DropdownMenu,
		DropdownMenuContent,
		DropdownMenuItem,
		DropdownMenuTrigger
	} from '$lib/components/ui/dropdowns';

	let {
		defaultTeam,
		teams
	}: {
		defaultTeam: string;
		teams: string[];
	} = $props();

	let selectedTeam = $state(defaultTeam);
</script>

<DropdownMenu>
	<DropdownMenuTrigger>
		{#snippet child({ props })}
			<Button variant="ghost" class="p-0 hover:bg-transparent" {...props}>
				<span
					class="bg-primary text-primary-foreground flex size-8 items-center justify-center rounded-full"
				>
					{selectedTeam.charAt(0).toUpperCase()}
				</span>
				<div class="flex flex-col gap-0.5 leading-none">
					<span class="">{selectedTeam}</span>
				</div>
				<ChevronUpIcon size={14} class="text-muted-foreground/80" />
			</Button>
		{/snippet}
	</DropdownMenuTrigger>
	<DropdownMenuContent align="start">
		{#each teams as team, i (i)}
			<DropdownMenuItem
				onSelect={() => {
					selectedTeam = team;
				}}
			>
				{team}
			</DropdownMenuItem>
		{/each}
	</DropdownMenuContent>
</DropdownMenu>
