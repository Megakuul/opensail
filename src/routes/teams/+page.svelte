<script>
  import { Teams, TeamsException } from "$lib/data/teams.svelte";
  import { Versions } from "$lib/data/versions.svelte";
  import { onMount } from "svelte";

  onMount(async () => {
    await Versions.load();
  })
</script>

<h1>Teams</h1>

{#if Teams}
  <p>Team</p>
{:else if TeamsException}
<p>
  {TeamsException}
</p>
{:else}
<div class="flex items-center justify-center">

  <p>Loading...</p>
</div>
{/if}
