<script>
  import { Teams } from "$lib/data/teams.svelte";
  import { cn } from "$lib/utils";
  import { fade } from "svelte/transition";

  let { class: klass } = $props();

  let searchException = $state("");

  let searchQuery = $state("");

  /** @param {string} query */
  const searchTeams = (query) => {
    Teams.reset();
    if (!Teams.teams || !query) {
      return;
    }
    try {
      searchException = "";

      /** @type {import("$lib/adapter/teams/teams.js").TeamMap} */
      let queriedTeams = {};
      
      // this is extremly inefficient in theory; in praxis this algorithm can easily handle
      // the applications maximum scale (limited by the global number of unique regatta boat teams).
      for (const [key, map] of Object.entries(Teams.teams)) {
        if (key.includes(query, 0) || JSON.stringify(map).includes(query, 0)) {
          queriedTeams[key] = map;
        }
      }

      Teams.replace(queriedTeams);
    } catch (err) {
      console.error(err);
      searchException = "ship query failed";
    }
  }
</script>

<div class="w-full flex flex-col gap-1 items-center">
  <input placeholder="Search for Team..." bind:value={searchQuery} onsubmit="{() => searchTeams(searchQuery)}"
  class="{cn("search-box", searchException?"outline outline-red-700/80":"", klass)}" />
  {#if searchException}
    <p transition:fade class="w-full text-xs text-start">{searchException}</p>
  {/if}
</div>


<style>
  .search-box {
    @apply w-3/4 sm:w-2/4 h-10 p-2 rounded-lg text-slate-50/70 bg-slate-50/20 focus:outline focus:outline-slate-100/60 focus:outline-2;
  }
</style>