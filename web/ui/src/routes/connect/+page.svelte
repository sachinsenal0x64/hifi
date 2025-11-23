<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import Profile from '$lib/components/Profile.svelte';
	import { Tabs, TabsContent, TabsList, TabsTrigger } from '$lib/components/ui/tabs';
	import { Unplug, ShieldMinus } from 'lucide-svelte';
	import Textarea from '$lib/components/Connect.svelte';
	import Deactivate from '$lib/components/Account.svelte';
	import { goto } from '$app/navigation';

	let currentTab = 'tab-5';
	const { data } = $props();
</script>

<div class="flex min-h-screen items-start justify-center overflow-hidden bg-black">
	<div class="fixed mx-auto w-full max-w-7xl py-4">
		<div class="px-4 py-4">
			<Header />
		</div>
		<Tabs value={currentTab} class="flex max-h-[calc(100dvh-3rem)] w-full flex-col md:flex-row">
			<TabsList class="sticky top-6 flex w-80 flex-col gap-2 self-start bg-transparent p-4">
				<TabsTrigger
					value="tab-5"
					onclick={() => goto('/connect')}
					class="hidden w-full cursor-pointer items-center  justify-start gap-3 rounded-md px-6 py-4 text-3xl font-bold text-white transition data-[state=active]:bg-transparent data-[state=active]:text-white data-[state=inactive]:text-zinc-400 data-[state=inactive]:hover:bg-zinc-800 data-[state=inactive]:hover:text-white md:flex"
				>
					<Unplug class="h-8 w-8" /> Connect
				</TabsTrigger>

				<TabsTrigger
					value="tab-6"
					onclick={() => goto('/deactivate')}
					class="hidden w-full cursor-pointer items-center  justify-start gap-3 rounded-md px-6 py-4 text-3xl font-bold text-white transition data-[state=active]:bg-transparent data-[state=active]:text-white data-[state=inactive]:text-zinc-400 data-[state=inactive]:hover:bg-zinc-800 data-[state=inactive]:hover:text-white md:flex"
				>
					<ShieldMinus class="h-8 w-8" /> Deactivate
				</TabsTrigger>
			</TabsList>

			<div class="grid w-full grid-cols-1 gap-4 overflow-y-auto px-10 pb-20 md:grid-cols-2">
				<div
					class="border-border relative col-span-1 min-w-[300px] rounded-lg border bg-zinc-900 p-8 md:col-span-2 md:min-h-[480px] md:overflow-y-hidden"
				>
					<TabsContent value="tab-6" class="h-auto min-h-[300px]">
						<Tabs value="tab-6" class=" items-center">
							<TabsList
								class="border-border h-auto gap-2 rounded-full border-b bg-zinc-800 px-10 py-2 text-zinc-400"
							>
								<TabsTrigger
									value="tab-6"
									class="relative cursor-pointer after:absolute  after:inset-x-0 after:bottom-0 after:-mb-1 after:h-0.5  hover:text-white hover:after:bg-white data-[state=active]:bg-transparent data-[state=active]:text-white data-[state=active]:shadow-none data-[state=active]:after:bg-white data-[state=active]:hover:text-white"
								>
									Account
								</TabsTrigger>
							</TabsList>
							<TabsContent value="tab-6">
								<Deactivate username={data.user?.username} />
							</TabsContent>
						</Tabs>
					</TabsContent>
					<TabsContent value="tab-5" class="h-auto min-h-[300px]">
						<Tabs value="tab-4" class=" items-center">
							<TabsList
								class="border-border h-auto gap-2 rounded-full border-b bg-zinc-800 px-10 py-2 text-zinc-400"
							>
								<TabsTrigger
									value="tab-4"
									class="relative cursor-pointer after:absolute  after:inset-x-0 after:bottom-0 after:-mb-1 after:h-0.5  hover:text-white hover:after:bg-white data-[state=active]:bg-transparent data-[state=active]:text-white data-[state=active]:shadow-none data-[state=active]:after:bg-white data-[state=active]:hover:text-white"
								>
									Connect
								</TabsTrigger>

								<TabsTrigger
									value="tab-5"
									class="relative cursor-pointer after:absolute  after:inset-x-0 after:bottom-0 after:-mb-1 after:h-0.5  hover:text-white hover:after:bg-white data-[state=active]:bg-transparent data-[state=active]:text-white data-[state=active]:shadow-none data-[state=active]:after:bg-white data-[state=active]:hover:text-white"
								>
									Profile
								</TabsTrigger>
							</TabsList>

							<TabsContent value="tab-4">
								<div class=" *:not-first:mt-4 mt-8">
									<Textarea user={data.user} />
								</div>
							</TabsContent>
							<TabsContent value="tab-5">
								<Profile />
							</TabsContent>
						</Tabs>
					</TabsContent>
				</div>
			</div>
		</Tabs>
	</div>
</div>
