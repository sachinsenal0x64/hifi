export const ENHANCED_IMAGE_REGEX = /<enhanced:img/;
export const GITHUB_REPO_URL = 'https://github.com/max-got/originui-svelte/' as const;

export const POSSIBLE_DEPENDENCIES = [
	{
		dev: false,
		name: 'bits-ui',
		packageName: 'bits-ui',
		url: 'https://github.com/huntabyte/bits-ui'
	},
	{
		dev: false,
		name: '@lucide/svelte',
		packageName: '@lucide/svelte',
		url: 'https://github.com/lucide-icons/lucide'
	},
	{
		dev: true,
		name: '~icons/ri',
		packageName: '@iconify-json/ri',
		url: 'https://www.npmjs.com/package/@iconify-json/ri'
	},
	{
		dev: true,
		name: '~icons/ri',
		packageName: 'unplugin-icons',
		url: 'https://github.com/antfu/unplugin-icons'
	},
	{
		dev: false,
		name: '@internationalized/date',
		packageName: '@internationalized/date',
		url: 'https://github.com/adobe/react-spectrum/tree/main/packages/@internationalized/date'
	},
	{
		dev: false,
		name: 'inputmask',
		packageName: 'inputmask',
		url: 'https://github.com/RobinHerbots/Inputmask'
	},
	{
		dev: false,
		name: 'svelte-sonner',
		packageName: 'svelte-sonner',
		url: 'https://github.com/wobsoriano/svelte-sonner'
	},
	{
		dev: false,
		name: '@dnd-kit-svelte/core',
		packageName: '@dnd-kit-svelte/core',
		url: 'https://github.com/HanielU/dnd-kit-svelte'
	},
	{
		dev: false,
		name: '@dnd-kit-svelte/modifiers',
		packageName: '@dnd-kit-svelte/modifiers',
		url: 'https://github.com/HanielU/dnd-kit-svelte'
	},
	{
		dev: false,
		name: '@dnd-kit-svelte/sortable',
		packageName: '@dnd-kit-svelte/sortable',
		url: 'https://github.com/HanielU/dnd-kit-svelte'
	},
	{
		dev: false,
		name: '@dnd-kit-svelte/utilities',
		packageName: '@dnd-kit-svelte/utilities',
		url: 'https://github.com/HanielU/dnd-kit-svelte'
	}
] as const;

export type PossibleDependency = (typeof POSSIBLE_DEPENDENCIES)[number];
