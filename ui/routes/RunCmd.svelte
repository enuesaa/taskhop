<script lang="ts">
	import { useRunCmd } from '$lib/graphql/runCmd'

	const runCmd = useRunCmd()

	async function handleClick() {
		await $runCmd.mutateAsync({input: { cmd: 'ls -la', memo: '' }})
	}
</script>

<button on:click|preventDefault={handleClick}>run</button>

{$runCmd.data}
