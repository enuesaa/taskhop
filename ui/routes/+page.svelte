<script lang="ts">
	import { getHealth } from '$lib/graphql/health'

	const health = getHealth()
</script>

a
{#if $health.isSuccess}
	{$health.data.code}
{/if}
