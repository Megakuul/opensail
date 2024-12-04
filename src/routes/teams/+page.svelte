<script>
    import ActionBar from "$lib/components/ActionBar.svelte";
    import ExceptionBar from "$lib/components/ExceptionBar.svelte";
    import Loader from "$lib/components/Loader.svelte";
  import { Manifest } from "$lib/data/manifest.svelte";
  import { Teams } from "$lib/data/teams.svelte";
  import { Versions } from "$lib/data/versions.svelte";
  import { onMount } from "svelte";

  onMount(async () => {
    Versions.load();
  })

  $effect(() => {
    Manifest.load(Versions.getLatest());
    Teams.load(Versions.getLatest());
  })
</script>

<div class="flex flex-col items-center min-h-[100vh]">
  <ActionBar></ActionBar>
  <h1 class="title text-7xl sm:text-9xl mt-12" title="Teams">Teams</h1>

  {#if false && Teams.teams}
    <p>Team</p>
    {:else if Versions.error()}
    <div class="mt-auto">
      <ExceptionBar class="" title={"Error - Failed to load versions"} message={Versions.error()}></ExceptionBar>
    </div>
    {:else if Manifest.error()}
    <div class="mt-auto">
      <ExceptionBar class="" title={"Error - Failed to load manifest"} message={Manifest.error()}></ExceptionBar>
    </div>
    {:else if Teams.error()}
    <div class="mt-auto">
      <ExceptionBar class="" title={"Error - Failed to load teams"} message={Teams.error()}></ExceptionBar>
    </div>
    {:else}
  <Loader class="w-36 h-[70vh]"></Loader>
  {/if}
</div>

