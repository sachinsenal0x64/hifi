/* eslint-disable perfectionist/sort-named-exports */
import Content from './dialog-content.svelte';
import Description from './dialog-description.svelte';
import Footer from './dialog-footer.svelte';
import Header from './dialog-header.svelte';
import Overlay from './dialog-overlay.svelte';
import Title from './dialog-title.svelte';

import { Dialog as DialogPrimitive } from 'bits-ui';

const Root = DialogPrimitive.Root;
const Trigger = DialogPrimitive.Trigger;
const Portal = DialogPrimitive.Portal;
const Close = DialogPrimitive.Close;

export {
	Close,
	Close as DialogClose,
	Content,
	Content as DialogContent,
	Description,
	Description as DialogDescription,
	Footer,
	Footer as DialogFooter,
	Header,
	Header as DialogHeader,
	Overlay,
	Overlay as DialogOverlay,
	Portal,
	Portal as DialogPortal,
	Root,
	Root as DialogRoot,
	Title,
	Title as DialogTitle,
	Trigger,
	Trigger as DialogTrigger
};


	export function close() {
			throw new Error('Function not implemented.');
		}
