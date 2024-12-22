<script lang="ts">
	import * as icons from 'lucide-svelte'
	import { XIcon } from 'lucide-svelte'

	export let iconName: string

	let LucideIcon = icons[iconName]
</script>

{#if LucideIcon}
	<LucideIcon {...$$props} />
{:else}
	<XIcon {...$$props} />
{/if}
