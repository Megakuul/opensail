<script>
  import Icon from "@iconify/svelte";
  import { onMount } from "svelte";
  import Ship from "$lib/components/Ship.svelte";
  import { getRatingColor } from "./rating";

  /**
   * @typedef Props
   * @property {import("$lib/adapter/ships/ships.js").ShipConfig | null} mountedShip
  */

  /** @type {Props} */
  let { 
    mountedShip = $bindable()
  } = $props();

  const unmountShip = () => {
    window.history.pushState({}, "", `${window.location.pathname}`);
    selectedShipComponent = null;
    mountedShip = null;
  }

  /** @type {"spinnaker" | "jib" | "main" | "rigging" | "hull" | null} */
  let selectedShipComponent = $state(null);

  /** @param {KeyboardEvent} e */
  const escapeMountedShip = e => {
    if (e.key === "Escape") unmountShip();
  }

  onMount(async () => {
    if (document) document.addEventListener("keydown", escapeMountedShip);
  });
</script>

{#if mountedShip}
<div class="absolute inset-0 z-40 flex items-center justify-center bg-slate-900/70">
  <div class="relative w-full sm:w-[90%] h-[100dvh] sm:h-[90%] flex flex-col gap-4 p-4 sm:p-16 rounded-none sm:rounded-lg overflow-scroll-hidden bg-slate-900/90">
    <button class="absolute top-2 right-2 p-2 rounded-lg transition-all duration-700 hover:bg-slate-600/20" onclick="{() => unmountShip()}">
      <Icon icon="fluent-emoji-flat:cross-mark" width="24" height="24" />
    </button>
    <section class="w-full flex flex-row gap-2 items-center justify-start">
      <h1 class="font-bold text-xl sm:text-4xl text-slate-100/80">{mountedShip.boat_info.name}</h1>
      <h2 class="text-slate-200/40 text-base sm:text-3xl">({mountedShip.boat_info.class})</h2>
    </section>
    <section class="w-full flex flex-row gap-2 items-center justify-start text-slate-100/70 sm:text-base">
      <p>Sailed by Team
        <a href="/teams?team={mountedShip.team}" class="font-bold underline">{mountedShip.team}</a>.
      </p>
    </section>
    <section class="w-full flex flex-col sm:flex-row gap-2 items-start sm:items-center justify-start text-slate-100/70 text-xs sm:text-base">
      <p>Built in
        <span class="bg-slate-600/40 p-1 rounded-lg">{mountedShip.boat_info.age}</span> by 
        <span class="bg-slate-600/40 p-1 rounded-lg">{mountedShip.boat_info.builder}</span>.
      </p>
      <p>Designed by
        <span class="bg-slate-600/40 p-1 rounded-lg">{mountedShip.boat_info.designer}</span>. 
      </p>
    </section>
    <section class="w-full flex flex-col lg:flex-row gap-4 justify-between">
      <Ship class="w-full h-[256px] sm:h-[512px]" bind:selected="{selectedShipComponent}"></Ship>
      <div class="w-full h-[256px] sm:h-[512px] p-6 flex flex-col gap-4 rounded-lg overflow-scroll-hidden bg-slate-600/40 text-slate-100/70">
        {#if selectedShipComponent}
          <h1 class=" p-2 mb-2 text-2xl sm:text-4xl font-bold text-center rounded-lg bg-slate-600/30 flex-shrink-0">
            {selectedShipComponent.charAt(0).toUpperCase() + selectedShipComponent.slice(1).toLowerCase()}
          </h1>
          {#if selectedShipComponent === "spinnaker"}
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Symmetric Apex Sail:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.sail_area.symmetric_spinnaker} m²</p>
            </div>
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Asymmetric Apex Sail:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.sail_area.asymmetric_spinnaker} m²</p>
            </div>
          {:else if selectedShipComponent === "jib"}
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Apex Sail:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.sail_area.jib} m²</p>
            </div>
          {:else if selectedShipComponent === "main"}
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Apex Sail:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.sail_area.main} m²</p>
            </div>
          {:else if selectedShipComponent === "rigging"}
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Forestay Height:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.dimension.forestay_height} m</p>
            </div>
          {:else if selectedShipComponent === "hull"}
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Length overall:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.dimension.length_over_all} m</p>
            </div>
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Beam:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.dimension.beam} m</p>
            </div>
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Draft:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.dimension.draft} m</p>
            </div>
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Wetted area:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.dimension.wetted_surface_area} m²</p>
            </div>
          {/if}
        {:else}
          <h1 class="p-2 mb-2 text-2xl sm:text-4xl font-bold text-center bg-slate-600/30 rounded-lg">Misc</h1>
          <div class="flex flex-row justify-between text-base sm:text-2xl">
            <p>Crew Weight:</p>
            <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.misc.max_crew_weight} kg</p>
          </div>
          <div class="flex flex-row justify-between text-base sm:text-2xl">
            <p>Displacement:</p>
            <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.misc.sailing_displacement} kg</p>
          </div>
          <div class="flex flex-row justify-between text-base sm:text-2xl">
            <p>Stability Index:</p>
            <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_spec.misc.stability_index}</p>
          </div>
        {/if} 
      </div>
    </section>
    <hr class="border-slate-200/20 border-2 rounded-xl mt-auto">
    <section class="w-full flex flex-row gap-2 items-center justify-between text-slate-100/70 text-xl sm:text-2xl">
      <p class="font-bold">
        Openfactor (<span class="text-slate-400/60">
          {mountedShip.boat_rating.version}
        </span>):
      </p>
      <div class="bg-slate-300/80 rounded-lg px-5 py-3 ">
        <p class="font-bold" style="color: {getRatingColor(mountedShip.boat_rating.tcc, 3)}">
          {mountedShip.boat_rating.tcc?.toFixed(2)}
        </p>
      </div>
    </section>
  </div>
</div>
{/if}

<style>
  h1, h2, p {
    @apply text-nowrap overflow-hidden;
  }
</style>