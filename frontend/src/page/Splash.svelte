<script lang="ts">
    import { Events } from '@wailsio/runtime'
    import { navigate } from 'svelte-navigator'

    Events.On('app', () => {
        navigate('/detail', { replace: true })
    })
</script>

<main class="h-full flex flex-col justify-between items-center">
    <div></div>
    <div class="flex flex-col items-center">
        <img src="/appicon.png" alt="Runnr" width="96px" height="96px">
        <h1 class="mt-2 text-2xl focus-visible:outline-none">
            Mockswitch
        </h1>
    </div>
    <div class="flex items-center mb-8 gap-2">
        <img src="/connectedtech.png" alt="Connected Tech" width="32px">
        <div class="border-r border-gray-300 h-8"></div>
        <img src="/pixcee.png" alt="Connected Tech" width="32px">
    </div>
</main>
