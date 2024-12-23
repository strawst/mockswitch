<script lang="ts">
    import {Events} from '@wailsio/runtime'
    import {navigate} from "svelte-navigator";

    Events.On('app', (state: any) => {
        navigate('/status', {
                state: state.data[0],
                replace: true
            }
        )
    })
</script>

<main class="h-full flex flex-col justify-between items-center">
    <div></div>
    <div class="flex flex-col items-center">
        <img alt="Runnr" height="96px" src="/appicon.png" width="96px">
        <h1 class="mt-2 text-2xl focus-visible:outline-none">
            Mockswitch
        </h1>
    </div>
    <div class="flex items-center mb-8 gap-2">
        <img alt="Connected Tech" src="/connectedtech.png" width="32px">
        <div class="border-r border-gray-300 h-8"></div>
        <img alt="Connected Tech" src="/pixcee.png" width="32px">
    </div>
</main>
