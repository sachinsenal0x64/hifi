import type {
	OUIComponent,
	OUIDirectory,
	OUIDirectoryToComponent
} from '$lib/componentRegistry.types';

import { type ComponentState, OUI_DIRECTORIES } from '$lib/componentRegistry.components';

import { GITHUB_REPO_URL } from './constants';

interface ComponentMetadata {
	githubUrl: string;
	name: string;
	path: string;
	state: ComponentState;
}

interface DirectoryMetadata {
	apiUrl: string;
	componentCount: number;
	components: ComponentMetadata[];
	directoryPath: string;
	githubUrl: string;
	llmsTextUrl: string;
	stateBreakdown: Record<ComponentState, number>;
}

interface RegistryMetadata {
	directoriesBreakdown: Record<OUIDirectory, DirectoryMetadata>;
	totalComponents: number;
	totalDirectories: number;
}

// =========================================
// Component Registry Implementation
// =========================================

class ComponentRegistry {
	#components: Map<OUIDirectory, Set<OUIComponent>>;

	constructor() {
		this.#components = new Map();
		this.refresh();
	}

	// =========================================
	// Initialization
	// =========================================

	async refresh() {
		this.#components = await this.#initializeComponents();
	}

	#getFileImports = () => {
		return import.meta.glob<OUIComponent>(
			[
				'$lib/components/**/*.svelte',
				'!$lib/components/_extras/**/*.svelte',
				'!$lib/components/ui/**/*.svelte'
			],
			{ eager: false, import: 'default' }
		);
	};

	async #initializeComponents(): Promise<Map<OUIDirectory, Set<OUIComponent>>> {
		const files = this.#getFileImports();
		const componentMap = new Map<OUIDirectory, Set<OUIComponent>>();

		// In dev mode, we just need the paths
		const paths = Object.keys(files);

		for (const path of paths) {
			const match = path.match(/\/components\/([^/]+)\/([^/]+)\.svelte$/);
			if (!match) continue;

			const [, directory, filename] = match as [string, OUIDirectory, OUIComponent];

			if (!componentMap.has(directory)) {
				componentMap.set(directory, new Set());
			}
			componentMap.get(directory)?.add(filename);
		}

		return componentMap;
	}

	/**
	 * Determines the state of a component based on its filename
	 */
	#getComponentState(filename: string): ComponentState {
		if (filename.includes('.todo')) return 'todo';
		return 'ready';
	}

	/**
	 * Gets the GitHub URL for a directory or component
	 */
	#generateGithubUrl(directory: OUIDirectory, filename?: string): string {
		const basePath = `${GITHUB_REPO_URL}/tree/main/src/lib/components/${directory}`;
		return filename ? `${basePath}/${filename}.svelte` : basePath;
	}

	#getFile<T extends OUIDirectory>(directory: T): OUIDirectoryToComponent[T][] {
		const components = this.#components.get(directory);
		if (!components?.size) {
			throw new Error(`Components ${directory} not found in components/${directory}`);
		}
		return Array.from(components) as OUIDirectoryToComponent[T][];
	}

	getAllDirectories(): OUIDirectory[] {
		return Array.from(this.#components.keys());
	}

	getDirectories<T extends keyof typeof OUI_DIRECTORIES>(
		directories: T[]
	): (typeof OUI_DIRECTORIES)[T]['directory'][] {
		return directories.map((directory) => OUI_DIRECTORIES[directory].directory);
	}

	getFiles<T extends OUIDirectory>(directories: T[]): OUIDirectoryToComponent[T][] {
		return directories.flatMap((directory) => this.#getFile(directory));
	}

	getRegistryMetaData(): RegistryMetadata {
		const metadata: RegistryMetadata = {
			directoriesBreakdown: {} as Record<OUIDirectory, DirectoryMetadata>,
			totalComponents: 0,
			totalDirectories: Object.keys(OUI_DIRECTORIES).length
		};

		Object.values(OUI_DIRECTORIES).forEach((directoryConfig) => {
			const directory = directoryConfig.directory;
			const directoryPath = `src/lib/components/${directory}`;
			const components = this.#components.get(directory);
			const componentsMetadata: ComponentMetadata[] = [];
			const stateBreakdown = {
				ready: 0,
				soon: 0,
				todo: 0
			};

			components?.forEach((component) => {
				const state = this.#getComponentState(component);
				stateBreakdown[state]++;

				componentsMetadata.push({
					githubUrl: this.#generateGithubUrl(directory, component),
					name: component,
					path: `${directoryPath}/${component}`,
					state
				});
			});

			metadata.directoriesBreakdown[directory as OUIDirectory] = {
				apiUrl: `/api/v1/components/${directory}.json`,
				componentCount: components?.size || 0,
				components: componentsMetadata,
				directoryPath,
				githubUrl: this.#generateGithubUrl(directory),
				llmsTextUrl: `/llms/${directory}.txt`,
				stateBreakdown
			};

			metadata.totalComponents += components?.size || 0;
		});

		return metadata;
	}
}

let componentRegistryInstance: ComponentRegistry | null = null;

const initComponentRegistry = async () => {
	const registry = new ComponentRegistry();
	await registry.refresh();
	return registry;
};

export const getComponentRegistry = async () => {
	if (!componentRegistryInstance) {
		componentRegistryInstance = await initComponentRegistry();
	}
	return componentRegistryInstance;
};

export const getComponentDirectories = async () =>
	(await getComponentRegistry()).getAllDirectories();

export const getComponentFileNames = async <T extends OUIDirectory>(
	directory: T
): Promise<OUIDirectoryToComponent[T][]> => (await getComponentRegistry()).getFiles([directory]);
