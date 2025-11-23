<script lang="ts">
	import type { CountryCode, E164Number } from 'svelte-tel-input/types';
	import type { ChangeEventHandler } from 'svelte/elements';

	import Label from '$lib/components/ui/label.svelte';

	import ChevronDown from '@lucide/svelte/icons/chevron-down';
	import Phone from '@lucide/svelte/icons/phone';
	import { normalizedCountries, TelInput } from 'svelte-tel-input';
	import 'svelte-tel-input/styles/flags.css';

	let selectedCountry = $state<CountryCode | null>(null);
	let value = $state<E164Number | null>(null);

	const handleCountryChange: ChangeEventHandler<HTMLSelectElement> = (e) => {
		const { value } = e.currentTarget;
		selectedCountry = (value as CountryCode) || null;
	};

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2" dir="ltr">
	<Label for={uid}>Phone number input</Label>
	<div class="flex rounded-lg shadow-xs shadow-black/[.04]">
		<div
			class="border-input bg-background text-muted-foreground ring-offset-background focus-within:border-ring focus-within:text-foreground focus-within:ring-ring/30 hover:bg-accent hover:text-foreground relative inline-flex items-center self-stretch rounded-l-lg border py-2 ps-3 pe-2 transition-shadow focus-within:z-10 focus-within:ring-2 focus-within:ring-offset-2 focus-within:outline-hidden has-disabled:pointer-events-none has-disabled:opacity-50"
		>
			<div class="inline-flex items-center gap-1" aria-hidden="true">
				<span class="flex h-[16px] w-5 items-center overflow-hidden rounded-sm">
					{#if selectedCountry}
						<span
							class="flag flag-{selectedCountry.toLowerCase()} h-[13px]! w-5!"
							aria-hidden="true"
						></span>
					{:else}
						<Phone size={16} aria-hidden="true" />
					{/if}
				</span>
				<span class="text-muted-foreground/80">
					<ChevronDown size={16} aria-hidden="true" />
				</span>
			</div>
			<select
				onchange={handleCountryChange}
				class="absolute inset-0 text-sm opacity-0"
				aria-label="Select country"
			>
				<option value="">Select a country</option>
				{#each normalizedCountries as country (country.id)}
					<option value={country.id}>
						{country.label}
					</option>
				{/each}
			</select>
		</div>
		<TelInput
			id={uid}
			required
			placeholder="Enter phone number"
			class="border-input bg-background text-foreground ring-offset-background placeholder:text-muted-foreground/70 focus-visible:border-ring focus-visible:ring-ring/30 -ml-px flex h-9 w-full rounded-lg rounded-l-none border px-3 py-2 text-sm shadow-none shadow-black/[.04] transition-shadow focus-visible:z-10 focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-hidden disabled:cursor-not-allowed disabled:opacity-50"
			bind:country={selectedCountry}
			bind:value
			options={{
				format: 'international'
			}}
		/>
	</div>
	<p class="text-muted-foreground mt-2 text-xs" role="region" aria-live="polite">
		Built with <a
			class="hover:text-foreground underline"
			href="https://github.com/gyurielf/svelte-tel-input/tree/main"
			target="_blank"
			rel="noopener nofollow"
		>
			svelte-tel-input
		</a>
	</p>
</div>
