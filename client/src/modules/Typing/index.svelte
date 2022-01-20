<script lang="ts">
	import Typing from "./Typing.svelte"
	import Stats from "./stats/index.svelte"
	import Words from "./Words.svelte"
</script>

<div class="overflow-hidden flex flex-col justify-center p-12">
	<Stats />
	<Typing />
	<Words />
</div>