<script>
  import Icon from "@iconify/svelte";
  import { onMount } from "svelte";
  import Ship from "$lib/components/Ship.svelte";
  import { getRatingColor } from "./rating";
  import { goto } from "$app/navigation";
  import { fade } from "svelte/transition";

  /**
   * @typedef Props
   * @property {import("$lib/adapter/ships/ships.js").ShipConfig | null} mountedShip
  */

  /** @type {Props} */
  let { 
    mountedShip
  } = $props();

  const unmountShip = () => {
    selectedShipComponent = null;
    displayedMaterial = null;
    goto(`${window.location.pathname}`);
  }

  /** @param {KeyboardEvent} e */
  const escapeMountedShip = e => {
    if (e.key === "Escape") unmountShip();
  }

  onMount(async () => {
    if (document) document.addEventListener("keydown", escapeMountedShip);
  });

  /** @type {"spinnaker" | "jib" | "main" | "rigging" | "hull" | null} */
  let selectedShipComponent = $state(null);

  /**
   * @typedef Material
   * @property {string} name
   * @property {string} fg
   * @property {string} bg
   * @property {string} icon
   * @property {number | undefined} percent
  */

  /** @type {Material[]}*/
  let composition = $derived([
    { 
      name: "Ballast", fg: "rgb(255, 255, 255)", bg: "rgb(2 6 23)", icon:"game-icons:weight", 
      percent: mountedShip?.boat_extra_spec.composition.ballast_percentage 
    }, { 
      name: "Carbon", fg: "rgb(255, 255, 255)", bg: "rgb(68 64 60)", icon:"carbon:carbon",
      percent: mountedShip?.boat_extra_spec.composition.cfk_percentage 
    }, { 
      name: "Aluminium", fg: "rgb(0, 0, 0)", bg: "rgb(113 113 122)", icon:"icon-park-outline:heavy-metal",
      percent: mountedShip?.boat_extra_spec.composition.alu_percentage 
    }, { 
      name: "Fiberglass", fg: "rgb(0, 0, 0)", bg: "rgb(212 212 216)", icon:"fa6-solid:sheet-plastic",
      percent: mountedShip?.boat_extra_spec.composition.gfk_percentage 
    }, { 
      name: "Wood", fg: "rgb(255, 255, 255)", bg: "rgb(69 26 3)", icon:"game-icons:wood-beam",
      percent: mountedShip?.boat_extra_spec.composition.wood_percentage 
    }, { 
      name: "Engine", fg: "rgb(255, 255, 255)", bg: "rgb(41 37 36)", icon:"mdi:engine-outline",
      percent: mountedShip?.boat_extra_spec.composition.engine_percentage 
    }, { 
      name: "Amenities", fg: "rgb(255, 255, 255)", bg: "rgb(219, 89, 9)", icon:"mdi:fireplace",
      percent: mountedShip?.boat_extra_spec.composition.amenity_percentage 
    }
  ])

  /** @type {Material | null} */
  let displayedMaterial = $state(null);
</script>

{#if mountedShip}
<div class="fixed inset-0 z-40 flex items-center justify-center bg-slate-900/70 overscroll-contain">
  <div class="relative w-full sm:w-[90%] h-[100dvh] sm:h-[90%] flex flex-col gap-4 p-4 sm:p-16 rounded-none sm:rounded-lg overflow-scroll-hidden overscroll-contain bg-slate-900/90">
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
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_base_spec.sail_area.symmetric_spinnaker} m²</p>
            </div>
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Asymmetric Apex Sail:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_base_spec.sail_area.asymmetric_spinnaker} m²</p>
            </div>
          {:else if selectedShipComponent === "jib"}
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Apex Sail:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_base_spec.sail_area.jib} m²</p>
            </div>
          {:else if selectedShipComponent === "main"}
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Apex Sail:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_base_spec.sail_area.main} m²</p>
            </div>
          {:else if selectedShipComponent === "rigging"}
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Forestay Height:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_base_spec.dimension.forestay_height} m</p>
            </div>
          {:else if selectedShipComponent === "hull"}
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Length overall:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_base_spec.dimension.length_over_all} m</p>
            </div>
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Beam:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_base_spec.dimension.beam} m</p>
            </div>
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Draft:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_base_spec.dimension.draft} m</p>
            </div>
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Wetted area:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_base_spec.dimension.wetted_surface_area} m²</p>
            </div>
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Displacement:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_base_spec.dimension.displacement} kg</p>
            </div>
            <div class="flex flex-row justify-between text-base sm:text-2xl">
              <p>Crew Weight:</p>
              <p class="bg-slate-600/40 px-2 py-1 rounded-lg">{mountedShip.boat_base_spec.dimension.crew_weight} kg</p>
            </div>
          {/if}
        {:else}
          <h1 class="p-2 mb-2 text-2xl sm:text-4xl font-bold text-center bg-slate-600/30 rounded-lg">Misc</h1>
          <div class="flex flex-row justify-between text-base sm:text-2xl">
            <p>Sailing Mode:</p>
            <p class="bg-slate-600/40 px-3 py-1 rounded-lg font-bold">{mountedShip.boat_extra_spec.design.mode}</p>
          </div>
          <div class="flex flex-row justify-between text-base sm:text-2xl">
            <p>Stabilization:</p>
            <p class="bg-slate-600/40 px-3 py-1 rounded-lg font-bold">{mountedShip.boat_extra_spec.design.stabilization}</p>
          </div>
          <div class="flex flex-row justify-between text-base sm:text-2xl">
            <p>Hull Design:</p>
            <p class="bg-slate-600/40 px-3 py-1 rounded-lg font-bold">{mountedShip.boat_extra_spec.design.hull}</p>
          </div>
        {/if} 
      </div>
    </section>
    <div class="flex flex-row mt-5 w-full h-4 flex-shrink-0 rounded-full bg-slate-200/10 overflow-hidden">
      {#each composition.sort((a, b) => Number(b.percent) - Number(a.percent)) as material}
      <button aria-label="{material.name}" class="transition-all duration-700 brightness-75 hover:brightness-110 cursor-pointer" 
        style="width: {material.percent}%; background-color: {material.bg};"
        onclick={() => { 
          displayedMaterial = displayedMaterial?.name===material.name ? null : material;
        }}>
      </button>
      {/each}
    </div>
    <div class="flex flex-row justify-between items-center gap-2 min-h-8 w-full">
      {#if displayedMaterial}
        <div transition:fade class="w-72 h-full p-1 transition-all duration-700 rounded-lg text-center font-bold overflow-hidden" 
          style="background-color: {displayedMaterial.bg}; color: {displayedMaterial.fg}">
          {displayedMaterial.name}
        </div>
        <div transition:fade class="flex justify-center items-center w-10 h-full p1 transition-all duration-700 rounded-lg font-bold"
          style="background-color: {displayedMaterial.bg}; color: {displayedMaterial.fg}">
          <Icon icon={displayedMaterial.icon} class="animate-pulse"></Icon>
        </div>
      {/if}
    </div>
    <hr class="border-slate-200/20 border-2 rounded-xl mt-auto">
    <section class="w-full flex flex-row gap-2 items-center justify-start text-slate-100/70 text-lg sm:text-xl">
      <p class="font-bold">
        Speed Factor (<span class="text-slate-400/60">
          x{mountedShip.boat_rating.speed_influence}
        </span>):
      </p>
      <div class="bg-slate-300/80 rounded-lg px-3 py-1 cursor-pointer ml-auto" 
        title="Drag: {mountedShip.boat_rating.speed_drag_points}; Upwind: {mountedShip.boat_rating.speed_upwind_points}; Downwind: {mountedShip.boat_rating.speed_downwind_points}">
        <p class="font-bold" style="color: {getRatingColor(mountedShip.boat_rating.speed_factor, 3)}">
          {mountedShip.boat_rating.speed_factor?.toFixed(2)}
        </p>
      </div>
    </section>
    <section class="w-full flex flex-row gap-2 items-center justify-start text-slate-100/70 text-lg sm:text-xl">
      <p class="font-bold">        
        Stabilization Factor (<span class="text-slate-400/60">
          x{mountedShip.boat_rating.stabilization_influence}
        </span>):
      </p>
      <div class="bg-slate-300/80 rounded-lg px-3 py-1 cursor-pointer ml-auto"
        title="Stability: {mountedShip.boat_rating.stabilization_points}">
        <p class="font-bold" style="color: {getRatingColor(mountedShip.boat_rating.stabilization_factor, 3)}">
          {mountedShip.boat_rating.stabilization_factor?.toFixed(2)}
        </p>
      </div>
    </section>
    <section class="w-full flex flex-row gap-2 items-center justify-start text-slate-100/70 text-lg sm:text-xl">
      <p class="font-bold">        
        Agility Factor (<span class="text-slate-400/60">
          x{mountedShip.boat_rating.stabilization_influence}
        </span>):
      </p>
      <div class="bg-slate-300/80 rounded-lg px-3 py-1 cursor-pointer ml-auto"
        title="Agility: {mountedShip.boat_rating.agility_points}">
        <p class="font-bold" style="color: {getRatingColor(mountedShip.boat_rating.agility_factor, 3)}">
          {mountedShip.boat_rating.agility_factor?.toFixed(2)}
        </p>
      </div>
    </section>
    <section class="w-full flex flex-row gap-2 items-center justify-between text-slate-100/70 text-xl sm:text-2xl">
      <p class="font-bold">
        Openfactor (<span class="text-slate-400/60">
          {mountedShip.boat_rating.version}
        </span>):
      </p>
      <div class="bg-slate-300/80 rounded-lg px-5 py-3 cursor-pointer">
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

  .material-box {
    cursor: pointer;
    filter: brightness(80%);
    transition: all ease .7s;
  }

  .material-box:hover {
    filter: brightness(120%);
  }
</style>