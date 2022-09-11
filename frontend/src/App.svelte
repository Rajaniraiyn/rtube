<script lang="ts">
  import { Route, router } from "tinro";
  import Downoad from "./views/Downoad.svelte";
  import Home from "./views/Home.svelte";

  router.mode.hash();
</script>

<Route path="/">
  <Home />
</Route>
<Route path="/:id" let:meta>
  <Downoad id={meta.params.id} />
</Route>
