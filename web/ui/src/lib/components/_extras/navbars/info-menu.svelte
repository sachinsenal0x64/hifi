<!-- <script lang="ts">
	import Button from '$lib/components/ui/button.svelte';

	import BookIcon from '@lucide/svelte/icons/book';
	import InfoIcon from '@lucide/svelte/icons/info';
	import LifeBuoyIcon from '@lucide/svelte/icons/life-buoy';
	import MessageCircleMoreIcon from '@lucide/svelte/icons/message-circle-more';
	import {
		DropdownMenu,
		DropdownMenuContent,
		DropdownMenuItem,
		DropdownMenuLabel,
		DropdownMenuTrigger
	} from '$lib/components/ui/dropdowns';
</script>

<DropdownMenu>
	<DropdownMenuTrigger>
		{#snippet child({ props })}
			<Button
				size="icon"
				variant="ghost"
				class="size-8 rounded-full shadow-none"
				aria-label="Open edit menu"
				{...props}
			>
				<InfoIcon class="text-muted-foreground" size={16} aria-hidden="true" />
			</Button>
		{/snippet}
	</DropdownMenuTrigger>
	<DropdownMenuContent class="pb-2">
		<DropdownMenuLabel>Need help?</DropdownMenuLabel>
		<DropdownMenuItem class="cursor-pointer py-1 focus:bg-transparent focus:underline">
			{#snippet child({ props })}
				<a href="#" {...props}>
					<BookIcon size={16} class="opacity-60" aria-hidden="true" />
					Documentation
				</a>
			{/snippet}
		</DropdownMenuItem>
		<DropdownMenuItem class="cursor-pointer py-1 focus:bg-transparent focus:underline">
			{#snippet child({ props })}
				<a href="#" {...props}>
					<LifeBuoyIcon size={16} class="opacity-60" aria-hidden="true" />
					Support
				</a>
			{/snippet}
		</DropdownMenuItem>
		<DropdownMenuItem class="cursor-pointer py-1 focus:bg-transparent focus:underline">
			{#snippet child({ props })}
				<a href="#" {...props}>
					<MessageCircleMoreIcon size={16} class="opacity-60" aria-hidden="true" />
					Contact us
				</a>
			{/snippet}
		</DropdownMenuItem>
	</DropdownMenuContent>
</DropdownMenu> -->
