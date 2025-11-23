<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';
	import { usePasswordStrength } from '$lib/hooks/use-password-strength.svelte';
	import { cn } from '$lib/utils.js';

	import Check from '@lucide/svelte/icons/check';
	import Eye from '@lucide/svelte/icons/eye';
	import EyeOff from '@lucide/svelte/icons/eye-off';
	import X from '@lucide/svelte/icons/x';

	const uid = $props.id();
	const passwordStrength = usePasswordStrength({ id: uid });
</script>

<div>
	<!-- Password input field with toggle visibility button -->
	<div class="*:not-first:mt-2">
		<Label for={uid}>Input with password strength indicator</Label>
		<div class="relative">
			<Input
				id={passwordStrength.id}
				class="pe-9"
				placeholder="Password"
				type={passwordStrength.isVisible ? 'text' : 'password'}
				bind:value={passwordStrength.password}
				aria-describedby={passwordStrength.id}
			/>
			<button
				class="text-muted-foreground/80 hover:text-foreground focus-visible:border-ring focus-visible:ring-ring/50 absolute inset-y-0 end-0 flex h-full w-9 items-center justify-center rounded-e-md transition-[color,box-shadow] outline-none focus:z-10 focus-visible:ring-[3px] disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50"
				type="button"
				onclick={passwordStrength.toggleVisibility}
				aria-label={passwordStrength.isVisible ? 'Hide password' : 'Show password'}
				aria-pressed={passwordStrength.isVisible}
				aria-controls={passwordStrength.id}
			>
				{#if passwordStrength.isVisible}
					<EyeOff size={16} aria-hidden="true" />
				{:else}
					<Eye size={16} aria-hidden="true" />
				{/if}
			</button>
		</div>
	</div>

	<!-- Password strength indicator -->
	<div
		class="bg-border mt-3 mb-4 h-1 w-full overflow-hidden rounded-full"
		role="progressbar"
		aria-valuenow={passwordStrength.strengthScore}
		aria-valuemin={0}
		aria-valuemax={4}
		aria-label="Password strength"
	>
		<div
			class={cn(`h-full transition-all duration-500 ease-out`, passwordStrength.strengthColor)}
			style:width="{(passwordStrength.strengthScore / 4) * 100}%"
		></div>
	</div>

	<!-- Password strength description -->
	<p id="password-strength" class="text-foreground mb-2 text-sm font-medium">
		{passwordStrength.strengthText}. Must contain:
	</p>

	<!-- Password requirements list -->
	<ul class="space-y-1.5" aria-label="Password requirements">
		{#each passwordStrength.strength as req (req.text)}
			<li class="flex items-center space-x-2">
				{#if req.met}
					<Check size={16} class="text-emerald-500" aria-hidden="true" />
				{:else}
					<X size={16} class="text-muted-foreground/80" aria-hidden="true" />
				{/if}
				<span class={`text-xs ${req.met ? 'text-emerald-600' : 'text-muted-foreground'}`}>
					{req.text}
					<span class="sr-only">
						{req.met ? ' - Requirement met' : ' - Requirement not met'}
					</span>
				</span>
			</li>
		{/each}
	</ul>
</div>
