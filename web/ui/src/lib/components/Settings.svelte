<script lang="ts">
	import { defaults, superForm } from 'sveltekit-superforms';
	import { zod4 } from 'sveltekit-superforms/adapters';
	import { toast } from 'svelte-sonner';

	import Button from '$lib/components/ui/button.svelte';

	import * as Empty from '$lib/components/ui/empty/index.js';
	import ArrowUpRightIcon from '@lucide/svelte/icons/arrow-up-right';

	import IconSettingsFilled from '@lucide/svelte/icons/settings';
	import { goto } from '$app/navigation';

	let open = $state(false);
</script>

<Empty.Root>
	<Empty.Header>
		<Empty.Media variant="icon">
			<IconSettingsFilled />
		</Empty.Media>
		<Empty.Title class=" text-gray-200">Settings</Empty.Title>
		<Empty.Description class="text-gray-400"
			>Authenticate with your Tidal account to get tokens that you can use with Selfhosted HiFi.</Empty.Description
		>
	</Empty.Header>
	<Empty.Content>
		<div class="flex gap-2">
			<a href="https://tidal.com/" target="_blank" rel="noopener noreferrer">
				<Button class="flex cursor-pointer items-center gap-2" variant="outline">
					<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 256 171">
						<path
							fill="#0a0b09"
							d="m128.004 85.339l42.664 42.67l-42.664 42.667l-42.669-42.667zM42.667.002L85.335 42.67L42.667 85.34L0 42.67zm170.666 0L256 42.67l-42.667 42.67l-42.666-42.67l-42.663 42.669l-42.669-42.67L128.004 0l42.663 42.665z"
						/>
					</svg>

					<span>Connect Tidal Account</span>
				</Button>
			</a>
		</div>
	</Empty.Content>
	<Button variant="link" class="text-gray-400" size="sm">
		<a href="https://github.com/sachinsenal0x64/hifi">
			Learn More <ArrowUpRightIcon class="inline" />
		</a>
	</Button>
</Empty.Root>
