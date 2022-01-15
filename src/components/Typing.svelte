<script lang="ts">
	import { onMount } from 'svelte';
	import { getWords } from "$utils/words"
	import * as CLASS from "$utils/constants"

	const words = getWords()

	onMount(() => {
		window.document.onkeydown = (e) => handleKeystroke(e)
	})

	interface KeystrokeProps {
		key: string;
		altKey: boolean;
		ctrlKey: boolean;
		shiftKey: boolean;
	}

	const handleKeystroke = ({ key, altKey, ctrlKey, shiftKey }: KeystrokeProps) => {
		if (current.word === words.length) return
		switch (key) {
			case "Backspace":
				// check special keys
				// handleBackCtrl()
				handleBack()
				break;
			case " ":
				handleSpace()
				break;
			case "Enter":
				break;
			case "Control":
				break;
			case "Alt":
				break;
			default:
				// check special keys
				handleChar(key)
		}
	}

	const handleChar = (key: string) => {
		const currentChar = words[current.word][current.char]
		if (key === currentChar) {
			const DOMCurrentChar = getCurrentChar()
			DOMCurrentChar.classList.add("text-gray-100")
			current.char += 1
			return
		}
		handleCharError(key)
	}

	const handleSpace = () => {
		current.word += 1
		current.char = 0
	}

	const handleBack = () => {
		if (current.char === 0 && current.word === 0) return;
		if (current.char > words[current.word].length) {
			removeCurrentChar()
			current.char -= 1
			return
		}
		if (current.char === 0) {
			removeWordError()
			current.word -= 1
			const DOMCurrentWord = getCurrentWord()
			current.char = DOMCurrentWord.children.length ? DOMCurrentWord.children.length : words[current.word].length
			return
		}
		current.char -= 1
		const DOMCurrentChar = getCurrentChar()
		removeCharClass()
		if (DOMCurrentChar.innerText === words[current.word][current.char]) {
			removeWordError()
		}
	}

	const getCurrentChar = () => {
		const DOMRef = `c${current.word}${current.char}`
		return document.getElementById(DOMRef)
	}

	const getPrevChar = () => {
		const currentId = `${current.word}${current.char}`
		const prevId = parseInt(currentId) - 1
		const DOMRef = `c${prevId}`
		return document.getElementById(DOMRef)
	}

	const getCurrentWord = () => {
		const DOMRef = `w${current.word}`
		return document.getElementById(DOMRef)
	}

	const removeWordError = () => {
		const DOMCurrentWord = getCurrentWord()
		DOMCurrentWord.classList.remove("underline")
	}

	const handleCharError = (key: string) => {
		const DOMCurrentChar = getCurrentChar()
		if (!DOMCurrentChar || current.char > words[current.word].length) {
			current.char += 1
			const span = document.createElement("span")
			span.id = `c${current.word}${current.char}`
			span.className = CLASS.CHAR_ERROR
			span.innerText = key
			const DOMCurrentWord = getCurrentWord()
			DOMCurrentWord.appendChild(span)
			return
		}
		DOMCurrentChar.className = CLASS.CHAR_ERROR
		current.char += 1
	}

	const removeCharClass = () => {
		const DOMCurrentChar = getCurrentChar()
		DOMCurrentChar.classList.remove(CLASS.CHAR_ERROR)
		DOMCurrentChar.classList.remove("text-gray-100")
	}

	const removeCurrentChar = () => {
		const DOMCurrentWord = getCurrentWord()
		const DOMCurrentChar = getCurrentChar()
		DOMCurrentWord.removeChild(DOMCurrentChar)
	}

	$: current = {
		word: 0,
		char: 0,
	}
</script>

<div class="flex flex-wrap">
	{#each words as word, i}
		<div class="mr-3 text-gray-500 text-2xl font-medium" id={`w${i}`}>
			{#each word as char, j}
				<span class="transition-colors" id={`c${i}${j}`}>
					{char}
				</span>
			{/each}
		</div>
	{/each}
</div>