<script>
  import ActionBar from "$lib/components/ActionBar.svelte";
    import ExceptionBar from "$lib/components/ExceptionBar.svelte";
  import Loader from "$lib/components/Loader.svelte";
  import { Manifest } from "$lib/data/manifest.svelte";
  import { Ships } from "$lib/data/ships.svelte";
  import { Versions } from "$lib/data/versions.svelte";
  import Icon from "@iconify/svelte";
  import { onMount } from "svelte";
    import Search from "./Search.svelte";

  onMount(async () => {
    Versions.load();
  })

  $effect(() => {
    Manifest.load(Versions.getLatest());
    Ships.load(Versions.getLatest());
  })
</script>

<div class="flex flex-col items-center min-h-[100vh]">
  <ActionBar></ActionBar>
  <h1 class="title text-7xl sm:text-9xl mt-12" title="Ships">Ships</h1>

  <Search class="my-10"></Search>

  {#if Ships.ships}
    {#each Object.entries(Ships.ships).slice(0, 20) as [key, value]}
      <p>{key} {value.boat_info.builder}</p>
    {/each}
  {:else if Versions.error()}
  <div class="mt-auto">
    <ExceptionBar class="min-w-[50vw]" title={"Error - Failed to load versions"} message={Versions.error()}></ExceptionBar>
  </div>
  {:else if Manifest.error()}
  <div class="mt-auto">
    <ExceptionBar class="min-w-[50vw]" title={"Error - Failed to load manifest"} message={Manifest.error()}></ExceptionBar>
  </div>
  {:else if Ships.error()}
  <div class="mt-auto">
    <ExceptionBar class="min-w-[50vw]" title={"Error - Failed to load ships"} message={Ships.error()}></ExceptionBar>
  </div>
  {:else}
  <Loader class="w-36 h-[70vh]"></Loader>
  {/if}
</div>
