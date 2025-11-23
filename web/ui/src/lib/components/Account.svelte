<script lang="ts">
	import Button, { buttonVariants } from '$lib/components/ui/button.svelte';
	import Input from '$lib/components/ui/input.svelte';

	import CircleAlert from '@lucide/svelte/icons/circle-alert';
	import * as Dialog from '$lib/components/ui/dialog';
	import { formSchema2 } from '$lib/types/auth';

	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod4 } from 'sveltekit-superforms/adapters';
	import { toast } from 'svelte-sonner';
	import * as Form from '$lib/components/ui/form/index.js';

	import { cn } from '$lib/utils';

	import * as Empty from '$lib/components/ui/empty/index.js';
	import ArrowUpRightIcon from '@lucide/svelte/icons/arrow-up-right';

	import UserPen from '@lucide/svelte/icons/user-pen';
	import Loader2 from '@lucide/svelte/icons/loader-2';

	const { username } = $props();

	let open = $state(false);

	const form = superForm(defaults(zod4(formSchema2)), {
		validators: zod4(formSchema2),
		onSubmit: async () => {
			await new Promise((resolve) => setTimeout(resolve, 800));
		},

		onResult: ({ result }) => {
			if (result.type === 'redirect') {
				toast.promise(
					new Promise((resolve) => {
						setTimeout(resolve, 500);
					}),
					{
						loading: 'Deactivating your account...',
						success: 'Account deleted successfully!',
						error: 'Something went wrong. Please try again.'
					}
				);
			} else {
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
		<Empty.Title class=" text-gray-200">Delete Your Account</Empty.Title>
		<Empty.Description class="text-gray-400">Permanently delete your account</Empty.Description>
	</Empty.Header>
	<Empty.Content>
		<Dialog.Root bind:open>
			<Dialog.Trigger class={cn('cursor-pointer', buttonVariants({ variant: 'destructive' }))}>
				Delete Account</Dialog.Trigger
			>
			<Dialog.Content class="bg-zinc-900">
				<div class="flex flex-col items-center gap-2">
					<div
						class="border-border flex size-9 shrink-0 items-center justify-center rounded-full border"
						aria-hidden="true"
					>
						<CircleAlert class="text-red-400 opacity-80" size={32} />
					</div>
					<Dialog.Header>
						<Dialog.Title class="text-gray-300 sm:text-center">Final confirmation</Dialog.Title>
						<Dialog.Description class="text-gray-400 sm:text-center">
							This action cannot be undone. To confirm, please enter the username
							<span class="text-red-400">{username ?? 'John Doe'}</span>.
						</Dialog.Description>
					</Dialog.Header>
				</div>

				<form method="POST" class="space-y-5" use:enhance>
					<div class="space-y-2">
						<Form.Field {form} name="username">
							<Form.Control>
								{#snippet children({ props })}
									<Form.Label class="font-bold text-gray-300">Username</Form.Label>
									<Input
										class="border-zinc-700 text-white"
										placeholder="Type {username ?? 'John Doe'} to confirm"
										type="text"
										{...props}
										bind:value={$formData.username}
									/>
								{/snippet}
							</Form.Control>
							<Form.FieldErrors />
						</Form.Field>
					</div>
					<Dialog.Footer>
						<Form.Button
							variant="destructive"
							type="submit"
							class="flex-1 cursor-pointer"
							disabled={$formData.username !== username || $submitting}
						>
							{#if $submitting}
								<Loader2 class="size-4 animate-spin" />
							{:else}
								Delete
							{/if}
						</Form.Button>

						{#if !$submitting}
							<Dialog.Close
								type="button"
								class="{buttonVariants({ variant: 'outline' })} flex-1 cursor-pointer"
							>
								Cancel
							</Dialog.Close>
						{/if}
					</Dialog.Footer>
				</form>
			</Dialog.Content>
		</Dialog.Root>
	</Empty.Content>
	<Button variant="link" class="text-gray-400" size="sm">
		<a href="https://github.com/sachinsenal0x64/hifi">
			Learn More <ArrowUpRightIcon class="inline" />
		</a>
	</Button>
</Empty.Root>
