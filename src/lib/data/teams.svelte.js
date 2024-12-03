import { Versions } from "$lib/data/versions.svelte.js";
import { FetchTeams } from "$lib/adapter/teams/teams.js";

export let Teams = CreateTeams();

export function CreateTeams() {
  /** @type {import("$lib/adapter/teams/teams.js").TeamMap | null} */
  let teams = $state(null);

  /** @type {string} */
  let teamsException = $state("");

  return {
    get teams() { return teams },
    error: () => { return teamsException },
    load: async (/** @type {string} */ version) => {
      if (!version) return;
      try {
        teamsException = "";
        teams = null;
        teams = await FetchTeams(version);
      } catch (/** @type {any} */ err) {
        teamsException = err.message;
      }
    }
  }
}