import { Versions } from "$lib/data/versions.svelte.js";
import { FetchTeams } from "$lib/adapter/teams/teams.js";

/** @type {import("$lib/adapter/teams/teams.js").TeamMap | null} */
export let Teams = $state(null);

/** @type {string} */
export let TeamsException = $state("");

/** @param {string} version @throws {AdapterError} */
export const LoadTeams = async (version) => {
  try {
    Teams = await FetchTeams(version);
    TeamsException = "";
  } catch (/** @type {any} */ err) {
    TeamsException = err.message;
  }
}

$effect(() => {
  if (Versions.latest()) {
    LoadTeams(Versions.latest());
  }
})