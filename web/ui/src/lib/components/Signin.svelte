<script lang="ts">
	import * as Card from '$lib/components/ui/card/index.js';
	import { FieldDescription } from '$lib/components/ui/field/index.js';
	import { zod4 } from 'sveltekit-superforms/adapters';
	import { formSchema } from '$lib/types/auth';

	import { superForm } from 'sveltekit-superforms';
	import { toast } from 'svelte-sonner';
	import * as Form from '$lib/components/ui/form/index.js';

	import Input from '$lib/components/ui/input.svelte';

	import * as Empty from '$lib/components/ui/empty/index.js';
	import Loader2 from '@lucide/svelte/icons/loader-2';
	import UserCircle from '@lucide/svelte/icons/user-circle';
	import { page } from '$app/state';

	const form = superForm(page.data.form, {
		resetForm: true,
		validators: zod4(formSchema),
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
						loading: 'Signing you in...',
						success: 'Signed in successfully!',
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

<div class="flex flex-col gap-6">
	<Card.Root
		class="mx-auto max-w-md border border-zinc-700 bg-transparent xl:h-[450px] xl:w-[350px]"
	>
		<Card.Header class="text-center">
			<Empty.Header>
				<Empty.Media variant="icon">
					<UserCircle />
				</Empty.Media>
			</Empty.Header>
			<Card.Title class="text-xl text-gray-200">Login to your account</Card.Title>
		</Card.Header>
		<Card.Content>
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
						Log In
					{/if}
				</Form.Button>
			</form>
		</Card.Content>
		<FieldDescription class="text-center ">
			Don't have an account?
			<a href="/" class="underline underline-offset-4">Sign up</a>
		</FieldDescription>
		<FieldDescription class="px-6 text-center">
			By clicking Log in you agree to our <a
				href="https://github.com/sachinsenal0x64/hifi?tab=readme-ov-file#project-terms">Terms</a
			>.
		</FieldDescription>
	</Card.Root>
</div>
