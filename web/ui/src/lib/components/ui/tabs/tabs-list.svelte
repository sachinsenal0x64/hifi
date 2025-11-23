<script lang="ts">
  import { cn } from "$lib/utils.js";

  import { Tabs as TabsPrimitive } from "bits-ui";

  let {
    class: className,
    ref = $bindable(null),
    ...restProps
  }: TabsPrimitive.ListProps = $props();
</script>

<TabsPrimitive.List
  bind:ref
  class={cn(
    "bg-muted text-muted-foreground/70 inline-flex w-fit items-center justify-center rounded-md p-0.5",
    className,
  )}
  {...restProps}
/>
