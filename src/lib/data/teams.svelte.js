import { FetchTeams } from "$lib/adapter/teams/teams.js";

export let Teams = CreateTeams();

export function CreateTeams() {
  /** @type {import("$lib/adapter/teams/teams.js").TeamMap | null} */
  let teamsOriginal = $state(null);

  /** @type {import("$lib/adapter/teams/teams.js").TeamMap | null} */
  let teamsBuffer = $state(null);

  /** @type {string} */
  let teamsException = $state("");

  return {
    get teams() { return teamsBuffer },
    error: () => { return teamsException },
    load: async (/** @type {string} */ version) => {
      if (!version) return;
      try {
        teamsException = "";
        teamsOriginal = null;
        teamsBuffer = null;
        teamsOriginal = await FetchTeams(version);
        teamsBuffer = JSON.parse(JSON.stringify(teamsOriginal));
      } catch (/** @type {any} */ err) {
        teamsException = err.message;
      }
    },
    replace: (/** @type {import("$lib/adapter/teams/teams.js").TeamMap} */ map) => {teamsBuffer = map},
    reset: () => {teamsBuffer = JSON.parse(JSON.stringify(teamsOriginal));}
  }
}