<script lang="ts">
	import Input from '$lib/components/ui/input.svelte';

	const uid = $props.id();
</script>

<div class="group relative">
	<label
		for={uid}
		class="has-[+input:not(:placeholder-shown)):-translate-y-1/2 origin-start text-muted-foreground/70 group-focus-within:text-foreground has-[+input:not(:placeholder-shown)]:text-foreground absolute top-1/2 block -translate-y-1/2 cursor-text px-1 text-sm transition-all group-focus-within:pointer-events-none group-focus-within:top-0 group-focus-within:-translate-y-1/2 group-focus-within:cursor-default group-focus-within:text-xs group-focus-within:font-medium has-[+input:not(:placeholder-shown)]:pointer-events-none has-[+input:not(:placeholder-shown)]:top-0 has-[+input:not(:placeholder-shown)]:cursor-default has-[+input:not(:placeholder-shown)]:text-xs has-[+input:not(:placeholder-shown)]:font-medium"
	>
		<span class="bg-background inline-flex px-2">Input with label animation</span>
	</label>
	<Input id={uid} type="email" placeholder="" />
</div>
