<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import { Tabs, TabsContent, TabsList, TabsTrigger } from '$lib/components/ui/tabs';
	import Page from '$lib/components/Signup.svelte';
	import Settings from '$lib/components/Settings.svelte';
	import Marqueeck, { type MarqueeckOptions } from '@arisbh/marqueeck';
	import { Skeleton } from '$lib/components/ui/skeleton/index.js';

	const { data } = $props();

	const options: MarqueeckOptions = {
		gap: 10,
		speed: 30,
		hoverSpeed: 0
	};

	let currentTab = 'tab-0';
</script>

<div class="flex min-h-screen items-start justify-center overflow-hidden bg-black">
	<div class="fixed mx-auto w-full max-w-7xl py-4">
		<div class="px-4 py-4">
			<Header />
		</div>

		<Tabs
			value={currentTab}
			class="flex max-h-[calc(85dvh-3rem)] w-full flex-col gap-8 md:max-h-[calc(100dvh-3rem)] md:flex-row"
		>
			<div class="grid w-full grid-cols-1 gap-4 overflow-y-auto px-10 pb-20 md:grid-cols-2">
				<div
					class="border-border relative col-span-1 min-w-[300px] rounded-lg border bg-zinc-900 p-10 md:col-span-2 md:min-h-[480px] md:overflow-y-hidden"
				>
					<TabsContent value="tab-0" class="h-[600px] md:min-h-[400px] xl:h-[600px]">
						<Tabs value="tab-1" class=" items-center">
							<TabsList
								class="border-border h-auto  gap-2 rounded-full border-b bg-zinc-800 px-10 py-2 text-zinc-400"
							>
								<TabsTrigger
									value="tab-1"
									class="relative cursor-pointer after:absolute  after:inset-x-0 after:bottom-0 after:-mb-1 after:h-0.5  hover:text-white hover:after:bg-white data-[state=active]:bg-transparent data-[state=active]:text-white data-[state=active]:shadow-none data-[state=active]:after:bg-white data-[state=active]:hover:text-white"
								>
									Home
								</TabsTrigger>

								<TabsTrigger
									value="tab-2"
									class="relative cursor-pointer after:absolute  after:inset-x-0 after:bottom-0 after:-mb-1 after:h-0.5  hover:text-white hover:after:bg-white data-[state=active]:bg-transparent data-[state=active]:text-white data-[state=active]:shadow-none data-[state=active]:after:bg-white data-[state=active]:hover:text-white"
								>
									Settings
								</TabsTrigger>
							</TabsList>
							<TabsContent value="tab-1">
								<Page {data} />
								<div
									class="fade-mask mx-auto grid cursor-pointer items-center justify-center md:mt-[-2rem] md:w-[80%]"
								>
									<Marqueeck {options}>
										{#each data.albums as album}
											<a href="/signup">
												<img
													src={album.cover}
													alt={album.title}
													class="md:h-50 md:w-50 h-35 w-35 mx-auto rounded-2xl object-cover text-white hover:opacity-85"
												/>
											</a>
										{/each}

										<svelte:fragment slot="stickyEnd">
											<div
												class="rounded-3xl bg-[rgba(19,19,19,0.4)] px-4 py-2 text-white backdrop-blur-xl md:text-2xl"
											>
												Recent Albums
											</div>
										</svelte:fragment>
									</Marqueeck>

									{#if data.albums.length === 0}
										<Skeleton
											class="md:h-50 md:w-50 h-35 w-35 mx-auto rounded-2xl  bg-zinc-800 object-cover"
										/>
									{/if}
								</div>
							</TabsContent>
							<TabsContent value="tab-2">
								<Settings />
							</TabsContent>
						</Tabs>
					</TabsContent>
				</div>
			</div>
		</Tabs>
	</div>
</div>
