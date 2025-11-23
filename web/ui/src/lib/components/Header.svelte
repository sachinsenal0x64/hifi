<script lang="ts">
	import Button from '$lib/components/ui/button.svelte';
	import { NotificationMenu, UserMenu } from '$lib/components/_extras/navbars';
	import { Popover, PopoverTrigger } from '$lib/components/ui/popover';
	import { page } from '$app/state';
	import { Unplug, ShieldMinus } from 'lucide-svelte';
	import { goto } from '$app/navigation';

	const title = 'Hifi';
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<header class="border-b px-4 md:px-6">
	<div class="flex h-16 items-center justify-between gap-4">
		<div class="flex items-center gap-2">
			<Popover>
				<PopoverTrigger>
					{#snippet child({ props })}
						<Button class="group size-8 md:hidden" variant="ghost" size="icon" {...props}>
							<svg
								class="pointer-events-none"
								width={16}
								height={16}
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
								xmlns="http://www.w3.org/2000/svg"
							>
								<path
									d="M4 12L20 12"
									class="group-aria-expanded:rotate-315 origin-center -translate-y-[7px] transition-all duration-300 [transition-timing-function:cubic-bezier(.5,.85,.25,1.1)] group-aria-expanded:translate-x-0 group-aria-expanded:translate-y-0"
								/>
								<path
									d="M4 12H20"
									class="origin-center transition-all duration-300 [transition-timing-function:cubic-bezier(.5,.85,0.25,1.8)] group-aria-expanded:rotate-45"
								/>
								<path
									d="M4 12H20"
									class="group-aria-expanded:rotate-135 origin-center translate-y-[7px] transition-all duration-300 [transition-timing-function:cubic-bezier(.5,.85,.25,1.1)] group-aria-expanded:translate-y-0"
								/>
							</svg>
						</Button>
					{/snippet}
				</PopoverTrigger>
			</Popover>
			<div class="flex items-center">
				<button
					on:click={() => goto('/')}
					aria-label="Home"
					title="Home"
					class="text-primary hover:text-primary/90 cursor-pointer rounded-full p-3 hover:bg-zinc-600 hover:opacity-80"
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						aria-label="home"
						width="24"
						height="24"
						viewBox="0 0 24 24"
						><g
							fill="none"
							stroke="white"
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							><path d="M15 21v-8a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v8" /><path
								d="M3 10a2 2 0 0 1 .709-1.528l7-6a2 2 0 0 1 2.582 0l7 6A2 2 0 0 1 21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"
							/></g
						></svg
					>
				</button>

				{#if (page.data.user?.username && page.url.pathname !== '/connect') || (page.data.user?.username && page.url.pathname == '/') || (page.data.user?.username && page.url.pathname.startsWith('/deactivate'))}
					<a
						href="/connect"
						aria-label="Connect account"
						title="Connect"
						class=" text-primary hover:text-primary/90 rounded-full p-3 hover:bg-zinc-600 hover:opacity-80"
					>
						<Unplug size={24} class="text-white" aria-hidden="true" />
					</a>
				{/if}

				{#if page.data.user?.username && page.url.pathname !== '/deactivate' && page.data.user?.username && page.url.pathname !== '/'}
					<a
						href="/deactivate"
						aria-label="Deactivate account"
						title="Deactivate"
						class=" text-primary hover:text-primary/90 rounded-full p-3 hover:bg-zinc-600 hover:opacity-80 md:hidden"
					>
						<ShieldMinus size={24} class="text-white" aria-hidden="true" />
					</a>
				{/if}
			</div>
		</div>
		<!-- Right side -->
		<div class="flex items-center gap-4">
			<div class="flex items-center gap-2">
				<!-- Notification -->
				<NotificationMenu />
			</div>
			<!-- User menu -->
			<UserMenu user={page.data.user} />
		</div>
	</div>
</header>
