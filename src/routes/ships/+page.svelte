<script>
    import { Ships } from "$lib/data/ships.svelte";
  import { Versions } from "$lib/data/versions.svelte";
  import { onMount } from "svelte";

  onMount(async () => {
    await Versions.load();
  })
</script>

<h1>Ships</h1>

{#if Ships.ships}
  <p>Team</p>
{:else if Ships.error()}
<p>
  {Ships.error()}
</p>
{:else}
<div class="flex items-center justify-center">

  <p>Loading...</p>
</div>
{/if}