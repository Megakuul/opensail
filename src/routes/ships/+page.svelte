<script>
  import ActionBar from "$lib/components/ActionBar.svelte";
  import ExceptionBar from "$lib/components/ExceptionBar.svelte";
  import Loader from "$lib/components/Loader.svelte";
  import { Manifest } from "$lib/data/manifest.svelte";
  import { Ships } from "$lib/data/ships.svelte";
  import { Versions } from "$lib/data/versions.svelte";
  import Icon from "@iconify/svelte";
  import { mount, onMount } from "svelte";
  import Search from "./Search.svelte";

  /** @type {import("$lib/adapter/ships/ships.js").ShipConfig | null} */
  let mountedShip = $state(null);

  /** @param {KeyboardEvent} e */
  const escapeMountedShip = e => {
    if (e.key === "Escape") mountedShip = null;
  }

  onMount(async () => {
    if (document) document.addEventListener("keydown", escapeMountedShip);
    Versions.load();
  });

  $effect(() => {
    Manifest.load(Versions.getLatest());
    Ships.load(Versions.getLatest());
  })

  /** @param {number} tcc @param {number} boost @returns {string} */
  const getRatingColor = (tcc, boost) => {
    // the tcc anchor is 2.00 because ship tcc's are likely in range 0.00-2.00.
    const tccAnchor = 2.00;
    // move tcc to wrapped range to apply the boost then convert it back to original range (0.00-2.00).
    const boostedTcc = ((tcc - (tccAnchor / 2)) * boost) + (tccAnchor / 2);
    // the color anchor is 150 because hsl 0-150 is red to green.
    const colorAnchor = 150; 
    // move tcc to color range and flip it because red is the lowest and we want lowest to be green.
    const relativeTcc = colorAnchor - ((boostedTcc / tccAnchor) * colorAnchor);
    // generate color value by clamping the color range.
    const colorValue = Math.max(0, Math.min(relativeTcc, colorAnchor)); // clamp(anchor - relativeTcc)
    return `hsl(${Math.round(colorValue).toString()} 100% 28%)`;
  }
</script>

<div class="static flex flex-col items-center min-h-[100vh]">
  <ActionBar></ActionBar>
  <h1 class="title text-7xl sm:text-9xl mt-12" title="Ships">Ships</h1>

  <Search class="my-10"></Search>
  {#if mountedShip}
    <div class="absolute inset-0 z-40 flex items-center justify-center bg-slate-900/70">
      <div class="relative w-full sm:w-[90%] h-full sm:h-[90%] flex flex-col gap-4 p-16 bg-slate-900/90 rounded-lg">
        <button class="absolute top-2 right-2 p-2 rounded-lg transition-all duration-700 hover:bg-slate-600/20" onclick="{() => mountedShip = null}">
          <Icon icon="fluent-emoji-flat:cross-mark" width="24" height="24" />
        </button>
        <section class="w-full flex flex-row gap-2 items-center justify-start">
          <h1 class="font-bold text-xl sm:text-4xl text-slate-100/80">{mountedShip.boat_info.name}</h1>
          <h2 class="text-slate-200/40 text-base sm:text-3xl">({mountedShip.boat_info.class})</h2>
        </section>
        <section class="w-full flex flex-row gap-2 items-center justify-start text-slate-100/70">
          <p>Sailed by Team
            <a href="/teams?team={mountedShip.team}" class="font-bold underline">{mountedShip.team}</a>.
          </p>
        </section>
        <section class="w-full flex flex-row gap-2 items-center justify-start text-slate-100/70">
          <p>Built in
            <span class="bg-slate-600/40 p-1 rounded-lg">{mountedShip.boat_info.age}</span> by 
            <span class="bg-slate-600/40 p-1 rounded-lg">{mountedShip.boat_info.builder}</span>.
          </p>
          <p>Designed by
            <span class="bg-slate-600/40 p-1 rounded-lg">{mountedShip.boat_info.designer}</span>. 
          </p>
        </section>
        <hr class="border-slate-200/20 border-2 rounded-xl mt-auto">
        <section class="w-full flex flex-row gap-2 items-center justify-between text-slate-100/70">
          <p class="font-bold text-2xl">
            Openfactor (<span class="text-slate-400/60">
              {mountedShip.boat_rating.version}
            </span>):
          </p>
          <div class="bg-slate-300/80 rounded-lg px-5 py-3 overflow-hidden">
            <p class="font-bold text-2xl text-nowrap" style="color: {getRatingColor(mountedShip.boat_rating.tcc, 3)}">
              {mountedShip.boat_rating.tcc?.toFixed(2)}
            </p>
          </div>
        </section>
      </div>
    </div>  
  {/if}

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
