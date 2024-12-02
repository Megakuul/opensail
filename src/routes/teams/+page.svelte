<script>
  import { Teams } from "$lib/data/teams.svelte";
  import { Versions } from "$lib/data/versions.svelte";
  import { onMount } from "svelte";

  onMount(async () => {
    await Versions.load();
    await Teams.load();
  })
</script>

<h1>Teams</h1>

{#if Teams.teams}
  <p>Team</p>
{:else if Teams.error()}
<p>
  {Teams.error()}
</p>
{:else}
<div class="flex items-center justify-center">

  <p>Loading...</p>
</div>
{/if}
