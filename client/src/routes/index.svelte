<script lang="ts">
	import Typing from "$modules/Typing/index.svelte"
	import Result from "$modules/Result/index.svelte"
	import { state } from '../stores/states';

	const options = {
		typing: {
			component: Typing
		},
		result: {
			component: Result
		}
	}

	let active

	$: active = options[$state]
</script>

<svelte:head>
	<title>@typing</title>
</svelte:head>

<div>
	<svelte:component this={active.component} />
</div>
