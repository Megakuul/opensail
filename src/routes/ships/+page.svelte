<script>
  import ActionBar from "$lib/components/ActionBar.svelte";
    import ExceptionBar from "$lib/components/ExceptionBar.svelte";
  import Loader from "$lib/components/Loader.svelte";
  import { Manifest } from "$lib/data/manifest.svelte";
  import { Ships } from "$lib/data/ships.svelte";
  import { Versions } from "$lib/data/versions.svelte";
  import Icon from "@iconify/svelte";
  import { onMount } from "svelte";

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

  <input placeholder="Search for Ship..." class="w-3/4 sm:w-2/4 h-10 p-2 rounded-lg text-slate-50/70 bg-slate-50/20 focus:outline focus:outline-slate-100/60 focus:outline-2" />

  {#if Ships.ships}
    <p>Ships</p>
  {:else if Versions.error()}
  <div class="mt-auto">
    <ExceptionBar class="" title={"Error - Failed to load versions"} message={Versions.error()}></ExceptionBar>
  </div>
  {:else if Manifest.error()}
  <div class="mt-auto">
    <ExceptionBar class="" title={"Error - Failed to load manifest"} message={Manifest.error()}></ExceptionBar>
  </div>
  {:else if Ships.error()}
  <div class="mt-auto">
    <ExceptionBar class="" title={"Error - Failed to load ships"} message={Ships.error()}></ExceptionBar>
  </div>
  {:else}
  <Loader class="w-36 h-[70vh]"></Loader>
  {/if}
</div>
