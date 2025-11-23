<script lang="ts">
	import { minidenticon } from 'minidenticons';

	let {
		username,
		saturation = 95,
		lightness = 45,
		...others
	}: {
		username: string;
		saturation?: number | string;
		lightness?: number | string;
		[key: string]: any;
	} = $props();

	let svgURI: string = $derived(
		'data:image/svg+xml;utf8,' + encodeURIComponent(minidenticon(username, saturation, lightness))
	);
</script>

<img src={svgURI} alt={username} {...others} />
