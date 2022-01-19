<script lang="ts">
	import { word, char } from '../../stores/words';

	const getCurrentPos = () => {
		const prevCharID = `c${$word}${$char - 1}`
		const currCharID = `c${$word}${$char}`
		const caret = document.getElementById("caret")

		if ($char === 0) {
			let offsets = document.getElementById(currCharID).getBoundingClientRect()
			return [offsets.left, offsets.top]
		}
		if ($char === 1 && $word === 0) {
			caret.classList.remove("animate-caret")
		}
		let offsets = document.getElementById(prevCharID).getBoundingClientRect()
		return [offsets.right, offsets.top]
	}

	const setCurrentPos = () => {
		const [left, top] = getCurrentPos()
		const caret = document.getElementById("caret")
		caret.style.left = left + "px"
		caret.style.top = top + "px"
	}

	$: if (!($word === 0 && $char === 0)) {
		setCurrentPos()
	}

</script>

<div class="absolute h-8 w-0.5 bg-brand-500 rounded-full animate-caret transition-all" id="caret"></div>