<script>
  import { Teams } from "$lib/data/teams.svelte";
  import { cn } from "$lib/utils";
  import { fade } from "svelte/transition";

  let { class: klass, teamsPerCycle=150 } = $props();

  let searchException = $state("");

  let searchQuery = $state("");

  /** @param {string} query */
  const searchTeams = async (query) => {
    Teams.reset();
    if (!Teams.teams || !query) {
      return;
    }
    try {
      searchException = "";

      /** @type {import("$lib/adapter/teams/teams.js").TeamMap} */
      let matchingTeams = {};
      
      let teamIndex = 0;
      for (const [key, map] of Object.entries(Teams.teams)) {
        teamIndex++;
        if (key.includes(query, 0) || JSON.stringify(map).includes(query, 0)) {
          matchingTeams[key] = map;
        }

        // in maximum scale Teams.teams is assumed to hold >30000 teams (>1 MB)
        // parsing and searching all of them takes 0.5 seconds. To avoid blocking
        // the eventloop we register a promise that allows us to yield control
        // to the eventloop. Doing this 30000 times means 30000 eventloop iterations until
        // the ships are filtered, which is a long.. very long time. To avoid this we take
        // a middleground, where control is yielded after every xth team.
        // 'teamsPerCycle' define after how many ships are calculated before yielding control.
        // High number = fast ui & slow search; Low number = slow ui & fast search.
        if (!(teamIndex % teamsPerCycle)) await new Promise(resolve => setTimeout(resolve, 0));
      }
      Teams.replace(matchingTeams);
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