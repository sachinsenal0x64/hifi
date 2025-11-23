<script lang="ts">
  import { cn } from "$lib/utils.js";

  import { Tabs as TabsPrimitive } from "bits-ui";

  let {
    class: className,
    ref = $bindable(null),
    value,
    ...restProps
  }: TabsPrimitive.RootProps = $props();
</script>

<TabsPrimitive.Root
  bind:ref
  class={cn("flex min-w-0 flex-col gap-2", className)}
  {value}
  {...restProps}
/>
