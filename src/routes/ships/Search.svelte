<script>
  import { Ships } from "$lib/data/ships.svelte";
  import { cn } from "$lib/utils";
  import Icon from "@iconify/svelte";
  import { fade } from "svelte/transition";

  let { class: klass, shipsPerCycle=150 } = $props();

  let searchLoaderState = $state(false);

  let searchException = $state("");

  let searchQuery = $state("");

  /** @param {string} query */
  const searchShips = async (query) => {
    Ships.reset();
    if (!Ships.ships || !query) {
      return;
    }

    searchLoaderState = true;
    searchException = "";

    try {
      /** @type {import("$lib/adapter/ships/ships.js").ShipMap} */
      let matchingShips = {};

      let shipIndex = 0;
      for (const [key, value] of Object.entries(Ships.ships)) {
        shipIndex++;
        if (key.includes(query, 0) || JSON.stringify(value).includes(query, 0)) {
          matchingShips[key] = value;
        }

        // in maximum scale Ships.ships is assumed to hold 30000 ships (30 MB)
        // parsing and searching all of them takes 1-3 seconds. To avoid blocking
        // the eventloop we register a promise that allows us to yield control
        // to the eventloop. Doing this 30000 times means 30000 eventloop iterations until
        // the ships are filtered, which is a long.. very long time. To avoid this we take
        // a middleground, where control is yielded after every xth ship.
        // 'shipsPerCycle' define after how many ships are calculated before yielding control.
        // High number = fast ui & slow search; Low number = slow ui & fast search.
        if (!(shipIndex % shipsPerCycle)) await new Promise(resolve => setTimeout(resolve, 0));
      }
      Ships.replace(matchingShips);
    } catch (/** @type {any} */ err) {
      console.error(err.message);
      searchException = "ship query failed";
    }
    searchLoaderState = false;
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