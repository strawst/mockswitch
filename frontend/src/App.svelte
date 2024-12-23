<script lang="ts">
  import { Route, Router } from 'svelte-navigator'
  import './style/style.scss'
  import Splash from "$/page/Splash.svelte";
  import Status from "$/page/status/Status.svelte";
</script>

<main>
  <Router>
    <Route path="/" component={Splash} />
    <Route path="/status" component={Status} />
  </Router>
</main>
