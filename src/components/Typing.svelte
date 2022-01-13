<script>
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
</script>

<div class="flex flex-wrap">
	{#each words as word}
		<div class="mr-3 text-gray-500 text-2xl font-medium">
			{#each word as char}
				<span>
					{char}
				</span>
			{/each}
		</div>
	{/each}
</div>