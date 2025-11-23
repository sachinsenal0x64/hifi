<script lang="ts">
	import { superForm } from 'sveltekit-superforms';
	import { zod4 } from 'sveltekit-superforms/adapters';
	import { toast } from 'svelte-sonner';
	import * as Form from '$lib/components/ui/form/index.js';
	import { formSchema } from '$lib/types/auth';

	import Button, { buttonVariants } from '$lib/components/ui/button.svelte';
	import Input from '$lib/components/ui/input.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import { cn } from '$lib/utils';

	import * as Empty from '$lib/components/ui/empty/index.js';
	import IconUserCircle from '@tabler/icons-svelte/icons/user-circle';
	import ArrowUpRightIcon from '@lucide/svelte/icons/arrow-up-right';
	import Loader2 from '@lucide/svelte/icons/loader-2';
	import { goto } from '$app/navigation';

	let open = $state(false);

	const { data } = $props();

	const form = superForm(data.form, {
		resetForm: true,
		validators: zod4(formSchema),
		onSubmit: async () => {
			await new Promise((resolve) => setTimeout(resolve, 800));
		},
		onResult: ({ result }) => {
			if (result.type === 'success') {
				open = false;
				toast.promise(
					new Promise((resolve) => {
						setTimeout(resolve, 500);
					}),
					{
						loading: 'Account created successfully!',
						success: () => {
							setTimeout(() => {
								goto('/signin');
							}, 300);
							return 'Redirecting to Sign In...';
						},

						error: 'Something went wrong. Please try again.'
					}
				);
			} else {
				open = false;
				toast.error('Something went wrong. Please try again.');
			}
		}
	});

	const { form: formData, submitting, enhance } = form;
</script>

<Empty.Root>
	<Empty.Header>
		<Empty.Media variant="icon">
			<IconUserCircle />
		</Empty.Media>

		<Empty.Title class=" text-gray-200">
			{#if data.user}
				Welcome back, {data.user.username}!
			{:else}
				Let’s Get Started
			{/if}</Empty.Title
		>
		<Empty.Description class="text-gray-400">
			{#if data.user}
				<div class="inline-flex items-center gap-1">
					<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
						><path
							fill="#d7c997"
							d="M15.252 10.689c-.987-1.18-1.48-1.77-2.048-1.68c-.567.091-.832.803-1.362 2.227l-.138.368c-.15.405-.226.607-.373.756c-.146.149-.348.228-.75.386l-.367.143c-1.417.555-2.126.833-2.207 1.4s.52 1.049 1.721 2.011l.31.25c.342.273.513.41.611.597c.1.187.115.404.146.837l.029.394c.11 1.523.166 2.285.683 2.545s1.154-.155 2.427-.983l.329-.215c.362-.235.543-.353.75-.387c.208-.033.42.022.841.132l.385.1c1.485.386 2.228.58 2.629.173s.193-1.144-.221-2.62l-.108-.38c-.117-.42-.176-.63-.147-.837c.03-.208.145-.39.374-.756l.21-.332c.807-1.285 1.21-1.927.94-2.438c-.269-.511-1.033-.553-2.562-.635l-.396-.022c-.434-.023-.652-.035-.841-.13c-.19-.095-.33-.263-.61-.599z"
						/><path
							fill="#e1d3a0"
							d="m14.878 5.167l.183.522c.201.573.302.86.497 1.07c.196.212.464.324 1.001.547l.489.204c1.89.786 2.835 1.18 2.942 1.983c.045.334-.067.648-.33.998a4 4 0 0 0-.254-.06c-.526-.107-1.195-.143-1.877-.18l-.49-.026l-.23-.013l-.151-.18l-.316-.377c-.44-.528-.87-1.043-1.272-1.401c-.427-.382-1.141-.88-2.103-.726c-.972.156-1.49.863-1.77 1.367c-.259.468-.492 1.094-.728 1.727l-.17.458l-.072.193q-.08.03-.193.076l-.456.178c-.63.247-1.253.49-1.717.757c-.498.287-1.2.817-1.339 1.793c-.137.967.378 1.673.767 2.092c.29.311.681.64 1.098.978c-1.569.425-2.403.584-2.877.072c-.534-.576-.258-1.62.295-3.71l.142-.54c.157-.595.236-.891.197-1.186c-.04-.294-.193-.553-.499-1.07l-.278-.47C4.29 8.422 3.752 7.512 4.11 6.787c.36-.724 1.379-.783 3.418-.9l.527-.03c.58-.034.869-.05 1.122-.185c.252-.135.439-.372.813-.848l.34-.432c1.316-1.673 1.974-2.509 2.73-2.38s1.11 1.137 1.817 3.154"
						/></svg
					>
					<p class="mt-4 text-gray-200">New releases:</p>
				</div>
			{:else}
				Listen to Tidal music on your favorite Subsonic client by creating a HiFi account.
			{/if}
		</Empty.Description>
	</Empty.Header>
	<Empty.Content>
		{#if !data.user}
			<div class="flex gap-2">
				<Dialog.Root bind:open>
					<Dialog.Trigger
						class={cn('cursor-pointer hover:opacity-95', buttonVariants({ variant: 'outline' }))}
						>Get Started</Dialog.Trigger
					>
					<Dialog.Content class="bg-zinc-900">
						<div class="flex flex-col items-center gap-2">
							<Dialog.Header>
								<Empty.Header>
									<Empty.Media variant="icon">
										<IconUserCircle />
									</Empty.Media>
								</Empty.Header>
								<Dialog.Title class="text-gray-300 sm:text-center">HiFi</Dialog.Title>
								<Dialog.Description class="text-gray-400 sm:text-center">
									Create your HiFi account
								</Dialog.Description>
							</Dialog.Header>
						</div>

						<form method="POST" class="space-y-5" use:enhance>
							<div class="space-y-4">
								<div class="space-y-2">
									<Form.Field {form} name="username">
										<Form.Control>
											{#snippet children({ props })}
												<Form.Label class="font-bold text-gray-300">Username</Form.Label>
												<Input
													class="border-zinc-700 text-white"
													placeholder="Joe Doe"
													type="text"
													autocomplete="username"
													{...props}
													bind:value={$formData.username}
												/>
											{/snippet}
										</Form.Control>
										<Form.FieldErrors />
									</Form.Field>
								</div>
								<div class="space-y-2">
									<Form.Field {form} name="password">
										<Form.Control>
											{#snippet children({ props })}
												<Form.Label class="font-bold text-gray-300">Password</Form.Label>
												<Input
													class="border-zinc-700 text-white"
													placeholder="Secure Password"
													type="password"
													{...props}
													bind:value={$formData.password}
												/>
											{/snippet}
										</Form.Control>
										<Form.FieldErrors />
									</Form.Field>
								</div>
							</div>
							<Form.Button
								class="mt-2 w-full cursor-pointer "
								type="submit"
								variant="outline"
								disabled={$submitting}
								>{#if $submitting}
									<Loader2 class="size-4 animate-spin" />
								{:else}
									Create Account
								{/if}
							</Form.Button>
						</form>
					</Dialog.Content>
				</Dialog.Root>
			</div>
		{:else}
			<span
				class="inline-flex items-center gap-x-1.5 text-nowrap rounded-full bg-zinc-700 px-3 py-1.5 text-xs font-semibold text-white dark:bg-white/10 dark:text-white"
			>
				<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 48 48"
					><g fill="#dfdede"><circle cx="19" cy="33" r="9" /><path d="M24 6v27h4V14l11 3v-7z" /></g
					></svg
				>
				{data.titles[0]}</span
			>
			<span
				class="inline-flex items-center gap-x-1.5 text-nowrap rounded-full bg-zinc-700 px-3 py-1.5 text-xs font-semibold text-white dark:bg-white/10 dark:text-white"
			>
				<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 48 48"
					><g fill="#dfdede"><circle cx="19" cy="33" r="9" /><path d="M24 6v27h4V14l11 3v-7z" /></g
					></svg
				>
				{data.titles[1]}</span
			>
		{/if}
	</Empty.Content>
	<Button variant="link" class="text-gray-400" size="sm">
		<a href="https://github.com/sachinsenal0x64/hifi">
			Learn More <ArrowUpRightIcon class="inline" />
		</a>
	</Button>
</Empty.Root>
