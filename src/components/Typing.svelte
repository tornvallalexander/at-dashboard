<script lang="ts">
	import { onMount } from 'svelte';

	const words = [
		"rack",
		"taxi",
		"dish",
		"microphone",
		"sharp",
		"make",
		"talented",
		"solo",
		"race",
		"multimedia",
		"red",
		"proclaim",
		"name",
		"tycoon",
		"evolution",
		"effect",
		"linger",
		"learn",
		"lighter",
		"redundancy",
		"drift",
		"lesson",
		"cutting",
		"section",
		"house",
		"feminine",
		"future",
		"census",
		"exceed",
		"bathtub",
		"topple",
		"center",
		"world",
		"overeat",
		"surprise",
		"posture",
		"verdict",
		"leg",
		"gossip",
		"opposite",
		"operation",
		"goal",
		"traction",
		"recruit",
		"shelter",
		"reflect",
		"banner",
		"loss",
		"introduce",
		"dry",
	]

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

		console.log(key)
	}

	const handleChar = (key: string) => {
		const DOMCurrentWord = getCurrentWord()
		if (DOMCurrentWord.className.includes("underline")) {
			return
		}
		const currentChar = words[current.word][current.char]
		if (key === currentChar) {
			const DOMCurrentChar = getCurrentChar()
			DOMCurrentChar.classList.add("text-gray-100")
			current.char += 1
			return
		}
		DOMCurrentWord.classList.add("underline")
	}

	const handleSpace = () => {
		current.word += 1
		current.char = 0
	}

	const handleBack = () => {
		if (current.char === 0) {
			if (current.word === 0) {
				return
			}
			removeWordError()
			current.word -= 1
			current.char = words[current.word].length
			return
		}
		const DOMPrevChar = getPrevChar()
		DOMPrevChar.classList.remove("text-gray-100")
		current.char -= 1
		if (DOMPrevChar.innerText === words[current.word][current.char]) {
			removeWordError()
		}
	}

	const getCurrentChar = () => {
		const DOMRef = `c${current.word}${current.char}`
		return document.getElementById(DOMRef)
	}

	const getPrevChar = () => {
		const DOMRef = `c${current.word}${current.char - 1}`
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

	$: current = {
		word: 0,
		char: 0,
	}
</script>

<div class="flex flex-wrap">
	{#each words as word, i}
		<div class="mr-3 text-gray-500 text-2xl font-medium" id={`w${i}`}>
			{#each word as char, j}
				<span id={`c${i}${j}`}>
					{char}
				</span>
			{/each}
		</div>
	{/each}
</div>