<script lang="ts">
	import RunCmd from './RunCmd.svelte'
	import { getHealth } from '$lib/graphql/health'

	const health = getHealth()
</script>

a
{#if $health.isSuccess}
	{$health.data.code}
{/if}

<RunCmd />
