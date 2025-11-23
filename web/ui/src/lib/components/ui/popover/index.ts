/* eslint-disable perfectionist/sort-named-exports */
import Content from './popover-content.svelte';

import { Popover as PopoverPrimitive } from 'bits-ui';
const Root = PopoverPrimitive.Root;
const Trigger = PopoverPrimitive.Trigger;
const Close = PopoverPrimitive.Close;

export {
	Root,
	Content,
	Trigger,
	Close,
	//
	Root as Popover,
	Content as PopoverContent,
	Trigger as PopoverTrigger,
	Close as PopoverClose
};
