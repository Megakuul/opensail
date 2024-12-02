import { Versions } from "$lib/data/versions.svelte.js";
import { FetchShips } from "$lib/adapter/ships/ships.js";

/** @type {import("$lib/adapter/ships/ships.js").ShipMap | null} */
export let Ships = $state(null);

/** @type {string} */
export let ShipsException = $state("");

/** @param {string} version @throws {AdapterError} */
export const LoadShips = async (version) => {
  try {
    Ships = await FetchShips(version);
    ShipsException = "";
  } catch (/** @type {any} */ err) {
    ShipsException = err.message;
  }
}

$effect(() => {
  if (Versions.latest()) {
    LoadShips(Versions.latest());
  }
})