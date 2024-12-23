<script lang="ts">
    import {Route, Router} from 'svelte-navigator'
    import './style/style.scss'
    import Splash from "$/page/Splash.svelte";
    import Status from "$/page/status/Status.svelte";
    import MainLayout from "$/component/layout/MainLayout.svelte";
</script>

<main>
    <Router>
        <Route component={Splash} path="/"/>
        <Route>
            <MainLayout>
                <Route component={Status} path="/status"/>
            </MainLayout>
        </Route>
    </Router>
</main>
