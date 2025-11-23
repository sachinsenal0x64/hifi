<script lang="ts">
	import Button from '$lib/components/ui/button.svelte';

	import SettingsIcon from '@lucide/svelte/icons/settings';
	import {
		DropdownMenu,
		DropdownMenuContent,
		DropdownMenuItem,
		DropdownMenuTrigger
	} from '$lib/components/ui/dropdowns';
</script>

<DropdownMenu>
	<DropdownMenuTrigger>
		{#snippet child({ props })}
			<Button
				size="icon"
				variant="ghost"
				class="rounded-full shadow-none"
				aria-label="Open edit menu"
				{...props}
			>
				<SettingsIcon class="text-muted-foreground" size={16} aria-hidden="true" />
			</Button>
		{/snippet}
	</DropdownMenuTrigger>
	<DropdownMenuContent class="max-w-64">
		<DropdownMenuItem>Appearance</DropdownMenuItem>
		<DropdownMenuItem>Preferences</DropdownMenuItem>
		<DropdownMenuItem>API Settings</DropdownMenuItem>
	</DropdownMenuContent>
</DropdownMenu>
