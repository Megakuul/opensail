<script>
  import { Teams } from "$lib/data/teams.svelte";
  import { cn } from "$lib/utils";
  import Icon from "@iconify/svelte";
  import { fade } from "svelte/transition";

  let { class: klass, teamsPerCycle=150 } = $props();

  let searchLoaderState = $state(false);

  let searchException = $state("");

  let searchQuery = $state("");

  /** @param {string} query */
  const searchTeams = async (query) => {
    Teams.reset();
    if (!Teams.teams || !query) {
      return;
    }

    searchLoaderState = true;
    searchException = "";
    searchQuery = query.toLowerCase();

    try {
      /** @type {import("$lib/adapter/teams/teams.js").TeamMap} */
      let matchingTeams = {};

      let teamIndex = 0;
      for (const [key, valueMap] of Object.entries(Teams.teams)) {
        teamIndex++;
        const value = JSON.stringify(valueMap).toLowerCase()
        if (key.toLowerCase().includes(searchQuery, 0) || value.includes(searchQuery, 0)) {
          matchingTeams[key] = valueMap;
        }

        // in maximum scale Teams.teams is assumed to hold >30000 teams (4 MB)
        // parsing and searching all of them takes 1 second. To avoid blocking
        // the eventloop we register a promise that allows us to yield control
        // to the eventloop. Doing this 30000 times means 30000 eventloop iterations until
        // the teams are filtered, which is a long.. very long time. To avoid this we take
        // a middleground, where control is yielded after every xth team.
        // 'teamsPerCycle' define after how many teams are calculated before yielding control.
        // High number = fast ui & slow search; Low number = slow ui & fast search.
        if (!(teamIndex % teamsPerCycle)) await new Promise(resolve => setTimeout(resolve, 0));
      }
      Teams.replace(matchingTeams);
    } catch (/** @type {any} */ err) {
      console.error(err.message);
      searchException = "team query failed";
    }
    searchLoaderState = false;
  }
</script>

<div class="{cn("w-full flex flex-col gap-1 items-center text-slate-50/70", klass)}">
  <form class="w-full flex flex-row justify-center items-center gap-2" onsubmit="{() => searchTeams(searchQuery)}">
    <input type="text" placeholder="Search for Team..." bind:value={searchQuery}
    class="{cn("search-box", searchException?"outline outline-red-700/80":"")}" />
    {#if searchLoaderState}
      <Icon icon="svg-spinners:ring-resize" width="24" height="24" />
    {:else}
      <Icon icon="line-md:search-twotone" width="24" height="24" />
    {/if}
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