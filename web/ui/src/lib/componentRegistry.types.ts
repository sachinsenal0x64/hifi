/* eslint-disable */
/**
 * !!!!!!!!!!
 * This file is auto-generated. Do not edit manually
 * Last generated at: 6/18/2025, 8:46:01 PM
 * To update, run: pnpm generate:registry --format
 * @version 0.0.1
 * !!!!!!!!!!
 */

import type { Prettify } from '$lib/types/helpers';
import type { OUIDirectories } from './componentRegistry.components';

// Directory Type
export type OUIDirectory = OUIDirectoryHelper<keyof OUIDirectories>;

// Helpers
export type OUIDirectoryHelper<T extends keyof OUIDirectories> = OUIDirectories[T]['directory'];
export type OUIComponentHelper<T extends keyof OUIDirectories> =
	OUIDirectories[T]['components'][number];

// Status Helpers

// Status Type Helpers
export type OUITodoComponent<T extends keyof OUIDirectories> = Extract<
	OUIDirectories[T]['components'][number],
	`${string}todo.svelte`
>;

// Get all todo components
export type OUITodoComponents = {
	[K in keyof OUIDirectories]: OUITodoComponent<K>;
}[keyof OUIDirectories];

// Get todo components by directory
export type OUIDirectoryTodoComponents = {
	[K in keyof OUIDirectories]: {
		directory: K;
		components: OUITodoComponent<K>[];
	};
};

// Ready Component Helpers
export type OUIReadyComponent<T extends keyof OUIDirectories> = Exclude<
	OUIDirectories[T]['components'][number],
	`${string}todo.svelte`
>;

// Get all ready components
export type OUIReadyComponents = {
	[K in keyof OUIDirectories]: OUIReadyComponent<K>;
}[keyof OUIDirectories];

// Get ready components by directory
export type OUIDirectoryReadyComponents = {
	[K in keyof OUIDirectories]: {
		directory: K;
		components: OUIReadyComponent<K>[];
	};
};

// Component Count Types
export type OUIDirectoryComponentCounts = {
	[K in keyof OUIDirectories]: {
		directory: K;
		total: number;
		status: OUIDirectories[K]['status'];
	};
};

// Directory Status Filters
export type OUIDirectoriesWithTodo = {
	[K in keyof OUIDirectories as OUIDirectories[K]['status']['todo'] extends 0
		? never
		: K]: OUIDirectories[K];
};

// Component Metadata
export type OUIComponentMetadata = {
	[K in keyof OUIDirectories]: {
		directory: K;
		name: OUIDirectories[K]['name'];
		totalComponents: number;
		status: OUIDirectories[K]['status'];
		hasInProgress: boolean;
	};
};

// Component Types
export type OUIAccordionsComponents = OUIComponentHelper<'ACCORDIONS'>;
export type OUIAlertsComponents = OUIComponentHelper<'ALERTS'>;
export type OUIAvatarsComponents = OUIComponentHelper<'AVATARS'>;
export type OUIBadgesComponents = OUIComponentHelper<'BADGES'>;
export type OUIBannersComponents = OUIComponentHelper<'BANNERS'>;
export type OUIBreadcrumbsComponents = OUIComponentHelper<'BREADCRUMBS'>;
export type OUIButtonsComponents = OUIComponentHelper<'BUTTONS'>;
export type OUICheckboxesComponents = OUIComponentHelper<'CHECKBOXES'>;
export type OUIDialogsComponents = OUIComponentHelper<'DIALOGS'>;
export type OUIDropdownsComponents = OUIComponentHelper<'DROPDOWNS'>;
export type OUIInputsComponents = OUIComponentHelper<'INPUTS'>;
export type OUINavbarsComponents = OUIComponentHelper<'NAVBARS'>;
export type OUINotificationsComponents = OUIComponentHelper<'NOTIFICATIONS'>;
export type OUIPaginationsComponents = OUIComponentHelper<'PAGINATIONS'>;
export type OUIPopoversComponents = OUIComponentHelper<'POPOVERS'>;
export type OUIRadiosComponents = OUIComponentHelper<'RADIOS'>;
export type OUISelectsComponents = OUIComponentHelper<'SELECTS'>;
export type OUISlidersComponents = OUIComponentHelper<'SLIDERS'>;
export type OUISwitchesComponents = OUIComponentHelper<'SWITCHES'>;
export type OUITablesComponents = OUIComponentHelper<'TABLES'>;
export type OUITabsComponents = OUIComponentHelper<'TABS'>;
export type OUITextareasComponents = OUIComponentHelper<'TEXTAREAS'>;
export type OUITimelinesComponents = OUIComponentHelper<'TIMELINES'>;
export type OUITooltipsComponents = OUIComponentHelper<'TOOLTIPS'>;

// All Component Types
export type OUIComponent = Prettify<
	| OUIAccordionsComponents
	| OUIAlertsComponents
	| OUIAvatarsComponents
	| OUIBadgesComponents
	| OUIBannersComponents
	| OUIBreadcrumbsComponents
	| OUIButtonsComponents
	| OUICheckboxesComponents
	| OUIDialogsComponents
	| OUIDropdownsComponents
	| OUIInputsComponents
	| OUINavbarsComponents
	| OUINotificationsComponents
	| OUIPaginationsComponents
	| OUIPopoversComponents
	| OUIRadiosComponents
	| OUISelectsComponents
	| OUISlidersComponents
	| OUISwitchesComponents
	| OUITablesComponents
	| OUITabsComponents
	| OUITextareasComponents
	| OUITimelinesComponents
	| OUITooltipsComponents
>;

// Directory To Component
export type OUIDirectoryToComponent = Prettify<{
	accordions: OUIAccordionsComponents;
	alerts: OUIAlertsComponents;
	avatars: OUIAvatarsComponents;
	badges: OUIBadgesComponents;
	banners: OUIBannersComponents;
	breadcrumbs: OUIBreadcrumbsComponents;
	buttons: OUIButtonsComponents;
	checkboxes: OUICheckboxesComponents;
	dialogs: OUIDialogsComponents;
	dropdowns: OUIDropdownsComponents;
	inputs: OUIInputsComponents;
	navbars: OUINavbarsComponents;
	notifications: OUINotificationsComponents;
	paginations: OUIPaginationsComponents;
	popovers: OUIPopoversComponents;
	radios: OUIRadiosComponents;
	selects: OUISelectsComponents;
	sliders: OUISlidersComponents;
	switches: OUISwitchesComponents;
	tables: OUITablesComponents;
	tabs: OUITabsComponents;
	textareas: OUITextareasComponents;
	timelines: OUITimelinesComponents;
	tooltips: OUITooltipsComponents;
}>;
