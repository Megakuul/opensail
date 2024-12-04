<script>
  import { Manifest } from "$lib/data/manifest.svelte";
  import { Ships } from "$lib/data/ships.svelte";
  import { Teams } from "$lib/data/teams.svelte";
  import { Versions } from "$lib/data/versions.svelte";
  import Icon from "@iconify/svelte";
  import { fade } from "svelte/transition";

  /** @type {boolean} */
  let showVersionDropdown = $state(false);

  /** @param {string} version */
  const loadVersion = async (version) => {
    Versions.setLatest(version);
    await Manifest.load();
    await Ships.load();
    await Teams.load();
  }
</script>

<div class="flex flex-row justify-between w-full">
  <a href="/" class=" relative m-2 py-2 px-4 flex items-center gap-1 rounded-lg text-slate-300/80  bg-slate-300/15 hover:bg-slate-300/5 transition-all duration-700">
    <span>Return</span>
    <Icon icon="icon-park-outline:return" width="18" height="18" />
  </a>

  <input type="checkbox" id="versionDropdownCheckbox" class="invisible" bind:checked={showVersionDropdown}>
  <label for="versionDropdownCheckbox" class="relative m-2 py-2 px-4 flex items-center gap-1 cursor-pointer rounded-lg text-slate-300/80  bg-slate-300/15 transition-all duration-700">
    <span>
      {#if Versions.getLatest()}
      {Versions.getLatest()}
      {:else}
      <Icon icon="svg-spinners:tadpole" width="18" height="18" />
      {/if}
    </span>
    <Icon icon="qlementine-icons:version-control-16" width="18" height="18" />

    {#if showVersionDropdown}
    <div transition:fade class="absolute top-full z-50 right-0 h-64 w-28 mt-2 flex flex-col rounded-xl overflow-scroll-hidden bg-slate-300/15">
      {#each Versions.versions as version, i}
        <button class="m-1 py-2 px-3 flex items-center justify-between cursor-pointer rounded-lg hover:bg-slate-800/10 transition-all duration-700"
        onclick={() => {showVersionDropdown = !showVersionDropdown; loadVersion(version)}}>
          {#if i===0}
          <span class="font-bold">{version}</span>
          <Icon icon="mingcute:git-commit-fill" width="24" height="24" />
          {:else}
          <span>{version}</span>
          <Icon icon="mingcute:git-commit-line" width="24" height="24" />
          {/if}
        </button>
      {/each}
    </div>
    {/if}
  </label>
</div>