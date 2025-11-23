<script lang="ts">
  import { cn } from "$lib/utils.js";

  import { Tabs as TabsPrimitive } from "bits-ui";

  let {
    class: className,
    ref = $bindable(null),
    value,
    ...restProps
  }: TabsPrimitive.TriggerProps = $props();
</script>

<TabsPrimitive.Trigger
  bind:ref
  class={cn(
    "hover:text-muted-foreground focus-visible:border-ring focus-visible:ring-ring/50 aria-selected:text-foreground data-[state=active]:bg-background inline-flex items-center justify-center rounded-sm px-3 py-1.5 text-sm font-medium whitespace-nowrap outline-hidden  focus-visible:ring-[3px] disabled:pointer-events-none disabled:opacity-50 aria-selected:shadow-xs [&_svg]:shrink-0",
    className,
  )}
  {value}
  {...restProps}
/>
