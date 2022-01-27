<script lang="ts">
	import Typing from "./Typing.svelte"
	import Words from "./Words.svelte"
	import Wpm from "./Wpm.svelte"
	import Accuracy from "./Accuracy.svelte"
</script>

<div class="overflow-hidden flex flex-col justify-center p-12">
	<Words />
	<Typing />
	<div class="flex justify-around">
		<Wpm />
		<Accuracy />
	</div>
</div>