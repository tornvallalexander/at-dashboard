<script lang="ts">
	import { state } from "../../stores/states";
	import { States} from "$utils/constants";
	import { char, incorrectChars, totalChars, word, wordsPerMinute } from '../../stores/words';

	const startTyping = () => {
		state.set(States.typing)
		word.set(0)
		char.set(0)
		wordsPerMinute.set(0)
		incorrectChars.set(0)
		totalChars.set(0)
	}

	const accuracy = 100 - (($incorrectChars/$totalChars)*100).toFixed(0) || 100
</script>

<div class="text-center mt-36">
	<p class="text-9xl text-gray-300 font-bold mb-6">Awesome!</p>

	<div class="flex justify-around mb-20">
		<div class="p-12 shadow-xl">
			<p class="text-9xl text-gray-300 font-bold">{$wordsPerMinute}</p>
			<p class="text-md text-brand-500 font-bold uppercase">Words Per Minute</p>
		</div>
		<div class="p-12 shadow-xl">
			<p class="text-9xl text-gray-300 font-bold">{$word + 1}</p>
			<p class="text-md text-brand-500 font-bold uppercase"># of words</p>
		</div>
		<div class="p-12 shadow-xl">
			<p class="text-9xl text-gray-300 font-bold">{accuracy}</p>
			<p class="text-md text-brand-500 font-bold uppercase">% Accuracy</p>
		</div>
	</div>
	<button on:click={() => startTyping()} class="bg-gray-700 shadow-md text-white font-bold py-2 px-4 rounded">Go again</button>
</div>
