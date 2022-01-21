<script>
	import '../app.css';
	import Header from "../components/Header.svelte"
</script>

<svelte:head>
	<svlete:head>
		<link rel="preconnect" href="https://fonts.googleapis.com">
		<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
		<link href="https://fonts.googleapis.com/css2?family=Chakra+Petch:wght@500&display=swap" rel="stylesheet">
	</svlete:head>
</svelte:head>

<style>
	:global(body) {
		max-width: 100rem;
		margin: auto;
    font-family: 'Chakra Petch', sans-serif;
	}
</style>

<Header />
<div class="max-w-screen-lg mx-auto">
	<slot></slot>
</div>

