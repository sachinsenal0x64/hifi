<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import LoginForm from '$lib/components/Signin.svelte';
	import { Tabs, TabsContent, TabsList, TabsTrigger } from '$lib/components/ui/tabs';

	let currentTab = 'tab-0';
</script>

<div class="flex min-h-screen items-start justify-center overflow-hidden bg-black">
	<div class="fixed mx-auto w-full max-w-7xl py-4">
		<div class="px-4 py-4">
			<Header />
		</div>

		<Tabs
			value={currentTab}
			class="flex max-h-[calc(100dvh-3rem)] min-h-[600px] w-full flex-col gap-8 md:flex-row lg:max-h-[calc(95dvh-3rem)]"
		>
			<div class="grid w-full grid-cols-1 gap-4 overflow-y-auto px-10 pb-20 md:grid-cols-2">
				<div
					class="border-border relative col-span-1 min-w-[210px] rounded-lg border bg-zinc-900 p-5 md:col-span-2 md:min-h-[480px] md:overflow-y-hidden md:p-10"
				>
					<TabsContent value="tab-0" class="md:min-h-[300px] xl:h-[600px]">
						<Tabs value="tab-1" class=" items-center">
							<TabsList
								class="border-border h-auto gap-2 rounded-full border-b bg-zinc-800 px-10 py-2 text-zinc-400"
							>
								<TabsTrigger
									value="tab-1"
									class="relative cursor-pointer after:absolute  after:inset-x-0 after:bottom-0 after:-mb-1 after:h-0.5  hover:text-white hover:after:bg-white data-[state=active]:bg-transparent data-[state=active]:text-white data-[state=active]:shadow-none data-[state=active]:after:bg-white data-[state=active]:hover:text-white"
								>
									Sign In
								</TabsTrigger>
							</TabsList>
							<TabsContent value="tab-1">
								<div class="mt-5 flex w-full max-w-sm flex-col gap-6">
									<LoginForm />
								</div>
							</TabsContent>
						</Tabs>
					</TabsContent>
				</div>
			</div>
		</Tabs>
	</div>
</div>
