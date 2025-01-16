<script lang="ts">
    import {setContext} from 'svelte'
    import {writable} from 'svelte/store'
    import {useLocation} from "svelte-navigator";

    const location = useLocation();

    let initial = writable<any>({
        state: $location.state,
        selected: 0,
    })
    setContext('initial', initial)

    $: isWindows = navigator.userAgent.includes('Win64')
    $: margin = !isWindows ? 'mt-10' : ''
</script>

<main class="h-full flex flex-col justify-center items-center">
    {#if !isWindows}
        <div class="fixed top-0 w-full h-10 bg-white z-10 flex self-stretch justify-center items-center shadow">
            <img src="/appicon.png" alt="Mockswitch" width="18px" height="18px"/>
            <h1 class="text-sm text-gray-500 ml-1">
                {$initial.state.workspace?.name}
            </h1>
        </div>
    {/if}
    <div class="w-full h-full flex-1 flex flex-row overflow-clip {margin}">
        <slot/>
    </div>
</main>