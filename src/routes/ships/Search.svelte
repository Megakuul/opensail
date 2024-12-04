<script>
  import { Ships } from "$lib/data/ships.svelte";
  import { cn } from "$lib/utils";
  import Icon from "@iconify/svelte";
    import { onDestroy, onMount } from "svelte";
  import { fade } from "svelte/transition";

  let { class: klass } = $props();

  /** @type {Worker | null} */
  let searchWorker = null;

  let searchLoaderState = $state(false);

  let searchException = $state("");

  let searchQuery = $state("");

  onMount(() => {
    searchWorker = new Worker("SearchWorker.js");
  })

  onDestroy(() => {
    if (searchWorker) searchWorker.terminate();
  })

  /** @param {string} query */
  const searchShips = (query) => {
    Ships.reset();
    if (!Ships.ships || !query || !searchWorker) {
      return;
    }

    searchLoaderState = true;
    searchException = "";

    searchWorker.postMessage({components: JSON.stringify(Ships.ships), query: query})
    searchWorker.onmessage = function(e) {
      const { status, matching } = e.data;
      if (status) {
        console.error(status);
        searchException = "ship query failed";
      } else {
        Ships.replace(matching);
      }
      searchLoaderState = false;
    }
  }
</script>

<div class="{cn("w-full flex flex-col gap-1 items-center text-slate-50/70", klass)}">
  <form class="w-full flex flex-row justify-center items-center gap-2" onsubmit="{() => searchShips(searchQuery)}">
    <input type="text" placeholder="Search for Ship..." bind:value={searchQuery}
    class="{cn("search-box", searchException?"outline outline-red-700/80":"")}" />
    <Icon icon="svg-spinners:ring-resize" class={searchLoaderState?"opacity-100":"opacity-0"} width="24" height="24" />
  </form>
  {#if searchException}
    <p transition:fade class="text-xs font-bold text-red-800/80">{searchException}</p>
  {/if}
</div>

<style>
  .search-box {
    @apply w-3/4 sm:w-2/4 h-10 p-2 rounded-lg bg-slate-50/20 focus:outline focus:outline-slate-100/60 focus:outline-2;
  }
</style>