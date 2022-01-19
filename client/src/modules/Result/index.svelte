<script lang="ts">

	import { state } from "../../stores/states";
	import { States} from "$utils/constants";
</script>

<div>Test completed!</div>
<button on:click={() => state.set(States.typing)}>Go again</button>