<script lang="ts">
	import { cn } from '$lib/utils.js';

	import { DropdownMenu as DropdownMenuPrimitive } from 'bits-ui';

	let {
		class: className,
		onCloseAutoFocus,
		onPointerDown,
		onPointerDownOutside,
		portalProps,
		ref = $bindable(null),
		sideOffset = 4,
		...restProps
	}: DropdownMenuPrimitive.ContentProps & {
		onCloseAutoFocus?: (e: FocusEvent) => void;
		onPointerDown?: (e: PointerEvent) => void;
		onPointerDownOutside?: (e: PointerEvent) => void;
		portalProps?: DropdownMenuPrimitive.PortalProps;
	} = $props();

	let isCloseFromMouse = $state(false);

	function handlePointerDown(e: PointerEvent) {
		isCloseFromMouse = true;
		onPointerDown?.(e);
	}

	function handlePointerDownOutside(e: PointerEvent) {
		isCloseFromMouse = true;
		onPointerDownOutside?.(e);
	}

	function handleCloseAutoFocus(e: Event) {
		if (onCloseAutoFocus) {
			return onCloseAutoFocus(e);
		}

		if (!isCloseFromMouse) {
			return;
		}

		e.preventDefault();
		isCloseFromMouse = false;
	}
</script>

<DropdownMenuPrimitive.Portal {...portalProps}>
	<DropdownMenuPrimitive.Content
		bind:ref
		{sideOffset}
		class={cn(
			'border-border bg-popover text-popover-foreground data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 z-50 min-w-40 overflow-hidden rounded-lg border p-1 shadow-lg shadow-black/5 data-dropdown-menu-content:focus:outline-hidden',
			className
		)}
		onpointerdown={handlePointerDown}
		onInteractOutside={handlePointerDownOutside}
		onCloseAutoFocus={handleCloseAutoFocus}
		{...restProps}
	/>
</DropdownMenuPrimitive.Portal>
