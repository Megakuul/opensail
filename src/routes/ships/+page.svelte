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
  import ShipOverlay from "./ShipOverlay.svelte";
  import { getRatingColor } from "./rating";
  import { page } from '$app/stores';

  /** @type {import("$lib/adapter/ships/ships.js").ShipConfig | null} */
  let mountedShip = $state(null);

  onMount(async () => {
    await Versions.load();
    const requestedShip = $page.url.searchParams.get("ship");
    if (requestedShip) {
      await Ships.load(Versions.getLatest());
      mountedShip = Ships.ships?.[requestedShip] ?? null;
    }
  });

  $effect(() => {
    Manifest.load(Versions.getLatest());
    Ships.load(Versions.getLatest());
  })
</script>

<div class="static flex flex-col items-center min-h-[100vh]">
  <ActionBar></ActionBar>
  <h1 class="title text-7xl sm:text-9xl mt-12" title="Ships">Ships</h1>

  <Search class="my-10"></Search>

  <ShipOverlay bind:mountedShip={mountedShip}></ShipOverlay>

  {#if Ships.ships}
    <div class="flex flex-col gap-6 items-center w-full py-5 max-h-[100vh] overflow-scroll-hidden">
    {#each Object.entries(Ships.ships).slice(0, 20) as [key, value]}
      <button onclick="{() => mountedShip = value}"
        class="relative max-w-[1000px] w-9/12 p-5 flex flex-row gap-2 justify-start items-center cursor-pointer rounded-lg shadow-inner bg-slate-300/20 text-slate-200/70">
        <p class="font-bold text-nowrap overflow-hidden">{value.boat_info.name}</p>
        <p class="text-nowrap overflow-hidden opacity-65 hidden sm:block"># {key.toUpperCase().replace("_", " ")}</p>
        <div class="bg-slate-300/80 rounded-lg px-3 py-1 overflow-hidden ml-auto">
          <p class="font-bold text-nowrap" style="color: {getRatingColor(value.boat_rating.tcc, 3)}">
            {value.boat_rating.tcc?.toFixed(2)}
          </p>
        </div>

        {#if value.boat_spec.source === "orc"}
          <div class="absolute bottom-[-3px] right-[-3px] h-3 flex flex-row items-center gap-1 p-2 rounded-md bg-emerald-600/90">
            <span class="text-sm font-bold">orc</span>
            <Icon icon="ic:twotone-verified" width="12" height="12" />
          </div>
        {/if}
      </button>
    {/each}
    {#if Object.entries(Ships.ships).length < 1}
      <div class="w-9/12 flex justify-center text-slate-200/70">
        <p class="text-base sm:text-2xl text-center">Sorry, no ships match your search...</p>
      </div>
    {/if}
    </div>
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
