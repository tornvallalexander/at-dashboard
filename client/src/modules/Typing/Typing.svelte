<script lang="ts">
	import { onMount } from 'svelte';
	import { word, char } from '../../stores/words';
	import { state } from "../../stores/states"
	import { Class } from "$utils/constants"
	import axios from "axios"
	import variables from '../../lib/variables';
	import Caret from "./Caret.svelte"
	import { States } from '$utils/constants';

	let words = [];

	interface RandomTopWordsRes {
		data: {
			words: string[]
		}
	}

	onMount(async () => {
		const response = await axios.get<RandomTopWordsRes>(`${variables.baseURL}/randomTopWords`)
		words = response.data.words
	})

	interface KeystrokeProps {
		key: string;
		altKey: boolean;
		ctrlKey: boolean;
		shiftKey: boolean;
	}

	const handleKeydown = ({ key, altKey, ctrlKey, shiftKey }: KeystrokeProps) => {
		if ($state !== States.typing) return
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
			case "Meta":
				break;
			case "Tab":
				break;
			case "Shift":
				break;
			default:
				// check special keys
				handleChar(key)
		}
	}

	const handleChar = (key: string) => {
		const currentChar = words[$word][$char]
		if (key === currentChar) {
			const DOMCurrentChar = getCurrentChar()
			DOMCurrentChar.classList.add(Class.CHAR_CORRECT)
			if (checkComplete()) {
				char.update(ch => ch + 1)
				handleComplete()
				return
			}
			char.update(ch => ch + 1)
			return
		}
		handleCharError(key)
	}

	const handleSpace = () => {
		word.update(wd => wd + 1)
		char.set(0)
	}

	const handleBack = () => {
		if ($char === 0 && $word === 0) return;
		if ($char > words[$word].length) {
			char.update(ch => ch - 1)
			removeCurrentChar()
			return
		}
		if ($char === 0) {
			removeWordError()
			word.update(wd => wd - 1)
			const DOMCurrentWord = getCurrentWord()
			$char = DOMCurrentWord.children.length ? DOMCurrentWord.children.length : words[$word].length
			return
		}
		char.update(ch => ch - 1)
		const DOMCurrentChar = getCurrentChar()
		removeCharClass()
		if (DOMCurrentChar.innerText === words[$word][$char]) {
			removeWordError()
		}
	}

	const handleCharError = (key: string) => {
		const DOMCurrentChar = getCurrentChar()
		const DOMCurrentWord = getCurrentWord()
		if (!DOMCurrentChar || $char > words[$word].length) {
			const span = document.createElement("span")
			span.id = `c${$word}${$char}`
			span.className = `${Class.COLOR_TRANSITION} ${Class.CHAR_ERROR}`
			span.innerText = key
			const DOMCurrentWord = getCurrentWord()
			DOMCurrentWord.appendChild(span)
			char.update(ch => ch + 1)
			return
		}
		DOMCurrentChar.className = `${Class.COLOR_TRANSITION} ${Class.CHAR_ERROR}`
		DOMCurrentWord.classList.add(Class.WORD_INCORRECT)
		char.update(ch => ch + 1)
	}

	const getCurrentChar = () => {
		const DOMRef = `c${$word}${$char}`
		return document.getElementById(DOMRef)
	}

	const getCurrentWord = () => {
		const DOMRef = `w${$word}`
		return document.getElementById(DOMRef)
	}

	const removeWordError = () => {
		const DOMCurrentWord = getCurrentWord()
		DOMCurrentWord.classList.remove(Class.WORD_INCORRECT)
	}

	const removeCharClass = () => {
		const DOMCurrentChar = getCurrentChar()
		DOMCurrentChar.classList.remove(Class.CHAR_ERROR)
		DOMCurrentChar.classList.remove(Class.CHAR_CORRECT)
	}

	const removeCurrentChar = () => {
		const DOMCurrentWord = getCurrentWord()
		const DOMCurrentChar = getCurrentChar()
		DOMCurrentWord.removeChild(DOMCurrentChar)
	}

	const checkComplete = () => {
		return ($word === words.length - 1 && $char === words[$word].length - 1)
	}

	const handleComplete = () => {
		state.set(States.result)
	}

</script>

<svelte:window on:keydown={handleKeydown}/>
<svlete:head>
	<link rel="preconnect" href="https://fonts.googleapis.com">
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
	<link href="https://fonts.googleapis.com/css2?family=Oxygen+Mono&display=swap" rel="stylesheet">
</svlete:head>

<style>
	/*keep these styles for later usage*/
	:global(.text-gray-300) {}
  :global(.text-red-400) {}
	:global(.underline-error) {
		text-decoration: underline;
		text-decoration-color: #f87171;
	}

	/*div {*/
  /*  font-family: 'Oxygen Mono', monospace;*/
	/*}*/
</style>

<div class="flex flex-wrap max-w-[50rem]">
	<Caret />
	{#if words.length}
		{#each words as word, i}
			<div class="mr-3 text-gray-500 text-3xl font-regular tracking-wide" id={`w${i}`}>
				{#each word as char, j}
					<span class="transition-colors" id={`c${i}${j}`}>
						{char}
					</span>
				{/each}
			</div>
		{/each}
	{:else}
		<p>Loading...</p>
	{/if}
</div>