<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';
	import Label from '$lib/components/ui/label.svelte';
	import {
		Tooltip,
		TooltipContent,
		TooltipProvider,
		TooltipTrigger
	} from '$lib/components/ui/tooltip/index.js';
	import { cn } from '$lib/utils.js';

	import Check from '@lucide/svelte/icons/check';
	import Copy from '@lucide/svelte/icons/copy';

	let copied = $state(false);
	let inputElement = $state<HTMLInputElement | null>(null);

	async function handleCopy() {
		if (!inputElement) return;

		await navigator.clipboard.writeText(inputElement.value);
		copied = true;
		setTimeout(() => {
			copied = false;
		}, 1500);
	}

	const uid = $props.id();
</script>

<div class="*:not-first:mt-2">
	<Label for={uid}>Copy to clipboard</Label>
	<div class="relative">
		<Input
			bind:ref={inputElement}
			id={uid}
			class="pe-9"
			type="text"
			value="npx sv create my-app"
			readonly
		/>
		<TooltipProvider>
			<Tooltip>
				<TooltipTrigger
					onclick={handleCopy}
					class="text-muted-foreground/80 ring-offset-background hover:text-foreground focus-visible:border-ring focus-visible:text-foreground focus-visible:ring-ring/30 absolute inset-y-0 end-0 flex h-full w-9 items-center justify-center rounded-e-lg border border-transparent transition-shadow focus-visible:ring-2 focus-visible:ring-offset-2 focus-visible:outline-hidden disabled:pointer-events-none disabled:cursor-not-allowed"
					aria-label={copied ? 'Copied' : 'Copy to clipboard'}
					disabled={copied}
				>
					{#snippet child({ props })}
						<button {...props}>
							<div
								class={cn('transition-all', copied ? 'scale-100 opacity-100' : 'scale-0 opacity-0')}
							>
								<Check class="stroke-emerald-500" size={16} aria-hidden="true" />
							</div>
							<div
								class={cn(
									'absolute transition-all',
									copied ? 'scale-0 opacity-0' : 'scale-100 opacity-100'
								)}
							>
								<Copy size={16} aria-hidden="true" />
							</div>
						</button>
					{/snippet}
				</TooltipTrigger>
				<TooltipContent
					class="border-input bg-popover text-muted-foreground border px-2 py-1 text-xs"
				>
					Copy to clipboard
				</TooltipContent>
			</Tooltip>
		</TooltipProvider>
	</div>
</div>
