<script lang="ts">
	import Button from '$lib/components/ui/button.svelte';
	import IconFileOrientation from '@tabler/icons-svelte/icons/file-orientation';
	import * as Empty from '$lib/components/ui/empty/index.js';
	import ArrowUpRightIcon from '@lucide/svelte/icons/arrow-up-right';
</script>

<Empty.Root>
	<Empty.Header>
		<Empty.Media variant="icon">
			<IconFileOrientation />
		</Empty.Media>
		<Empty.Title class=" text-gray-200">Page Not Found</Empty.Title>
		<Empty.Description class="text-gray-400">
			The page you are looking for doesn't exist or has been moved.
		</Empty.Description>
	</Empty.Header>
	<Button variant="link" class="text-gray-400" size="sm">
		<a href="#/">
			Learn More <ArrowUpRightIcon class="inline" />
		</a>
	</Button>
</Empty.Root>
