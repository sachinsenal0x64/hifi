<script lang="ts">
	import Button from '$lib/components/ui/button.svelte';
	import { goto } from '$app/navigation';

	import LogOutIcon from '@lucide/svelte/icons/log-out';
	import { LogInIcon } from 'lucide-svelte';
	import { Avatar, AvatarFallback } from '$lib/components/ui/avatar';
	import {
		DropdownMenu,
		DropdownMenuContent,
		DropdownMenuItem,
		DropdownMenuLabel,
		DropdownMenuSeparator,
		DropdownMenuTrigger
	} from '$lib/components/ui/dropdowns';
	import ProfileIcon from '$lib/components/ProfileIcon.svelte';

	let guest = $state('Guest');

	const { user } = $props();
</script>

<DropdownMenu>
	<DropdownMenuTrigger>
		{#snippet child({ props })}
			<Button
				variant="ghost"
				class="h-auto cursor-pointer p-0 hover:bg-transparent hover:opacity-80"
				{...props}
			>
				<Avatar>
					<!-- <AvatarImage src="" alt="Profile image" /> -->
					<AvatarFallback>
						<ProfileIcon username={user?.username || guest} />
					</AvatarFallback>
				</Avatar>
			</Button>
		{/snippet}
	</DropdownMenuTrigger>
	<DropdownMenuContent class="max-w-64 bg-zinc-700" align="end">
		<DropdownMenuLabel class="flex min-w-0  flex-col">
			{#if user?.username}
				<span class="truncate text-sm font-medium text-zinc-100">{user.username}</span>
			{:else}
				<span class="truncate text-sm font-medium text-zinc-100">{guest}</span>
			{/if}
		</DropdownMenuLabel>

		{#if user?.username}
			<DropdownMenuSeparator class="bg-zinc-600" />
			<DropdownMenuItem
				onclick={() => goto('/signout')}
				class="cursor-pointer text-zinc-100 focus:bg-zinc-600 focus:text-white"
			>
				<LogOutIcon size={16} class="opacity-80" aria-hidden="true" />
				<span>Sign out</span>
			</DropdownMenuItem>
		{/if}

		{#if !user?.username}
			<DropdownMenuSeparator class="bg-zinc-600" />
			<DropdownMenuItem
				onclick={() => goto('/signin')}
				class="cursor-pointer text-zinc-100 focus:bg-zinc-600 focus:text-white"
			>
				<LogInIcon size={16} class="opacity-80" aria-hidden="true" />
				<span>Sign in</span>
			</DropdownMenuItem>
		{/if}
	</DropdownMenuContent>
</DropdownMenu>
