<script lang="ts">
	import { totalChars, incorrectChars } from "../../stores/words";

	let accuracy;

	$: accuracy = 100 - (($incorrectChars/$totalChars)*100).toFixed(0) || 100
</script>

<div class="p-2">
	<p class="text-[10rem] text-gray-900 font-medium">
		{accuracy}%
	</p>
</div>