<script lang="ts">
	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod4 } from 'sveltekit-superforms/adapters';
	import { toast } from 'svelte-sonner';
	import * as Form from '$lib/components/ui/form/index.js';
	import { updateSchema } from '$lib/types/auth';

	import Button, { buttonVariants } from '$lib/components/ui/button.svelte';
	import Input from '$lib/components/ui/input.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import { cn } from '$lib/utils';

	import * as Empty from '$lib/components/ui/empty/index.js';
	import ArrowUpRightIcon from '@lucide/svelte/icons/arrow-up-right';

	import { page } from '$app/state';
	import UserPen from '@lucide/svelte/icons/user-pen';
	import Loader2 from '@lucide/svelte/icons/loader-2';

	let open = $state(false);

	const form = superForm(defaults(zod4(updateSchema)), {
		validators: zod4(updateSchema),
		onSubmit: async () => {
			await new Promise((resolve) => setTimeout(resolve, 800));
		},
		onResult: ({ result }) => {
			if (result.type === 'success') {
				open = false;

				const promise = new Promise<void>((resolve) => {
					setTimeout(() => {
						resolve();
					}, 800);
				});

				toast.promise(promise, {
					loading: 'Updating your account...',
					success: 'Account updated successfully!',
					error: 'Something went wrong. Please try again.'
				});

				promise.then(() => {
					setTimeout(() => {
						location.replace('/connect');
					}, 800);
				});
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
			<UserPen />
		</Empty.Media>
		<Empty.Title class=" text-gray-200">Edit Your Profile</Empty.Title>
		<Empty.Description class="text-gray-400">Update your profile information</Empty.Description>
	</Empty.Header>
	<Empty.Content>
		<div class="flex gap-2">
			<Dialog.Root bind:open>
				<Dialog.Trigger class={cn('cursor-pointer', buttonVariants({ variant: 'outline' }))}
					>Edit Profile</Dialog.Trigger
				>
				<Dialog.Content class="bg-zinc-900">
					<div class="flex flex-col items-center gap-2">
						<Dialog.Header>
							<Empty.Header>
								<Empty.Media variant="icon">
									<UserPen />
								</Empty.Media>
							</Empty.Header>
							<Dialog.Title class=" text-gray-300 sm:text-center">HiFi</Dialog.Title>
							<Dialog.Description class="text-gray-400 sm:text-center">
								Update your HiFi account
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
												placeholder={page.data.user?.username ?? 'John Doe'}
												type="text"
												{...props}
												bind:value={$formData.username}
											/>
										{/snippet}
									</Form.Control>
									<Form.FieldErrors />
								</Form.Field>
							</div>
							<div class="space-y-2">
								<Form.Field {form} name="oldpassword">
									<Form.Control>
										{#snippet children({ props })}
											<Form.Label class="font-bold text-gray-300">Old Password</Form.Label>
											<Input
												class="border-zinc-700 text-white"
												placeholder="SuperSecret123!"
												type="password"
												{...props}
												bind:value={$formData.oldpassword}
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
											<Form.Label class="font-bold text-gray-300">New Password</Form.Label>
											<Input
												class="border-zinc-700 text-white"
												placeholder="SuperSecret123!"
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
							disabled={(!$formData.username && !$formData.oldpassword && !$formData.password) ||
								$submitting}
							>{#if $submitting}
								<Loader2 class="size-4 animate-spin" />
							{:else}
								Update Profile
							{/if}
						</Form.Button>
					</form>
				</Dialog.Content>
			</Dialog.Root>
		</div>
	</Empty.Content>
	<Button variant="link" class="text-gray-400" size="sm">
		<a href="https://github.com/sachinsenal0x64/hifi">
			Learn More <ArrowUpRightIcon class="inline" />
		</a>
	</Button>
</Empty.Root>
