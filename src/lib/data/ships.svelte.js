import { Versions } from "$lib/data/versions.svelte.js";
import { FetchShips } from "$lib/adapter/ships/ships.js";

export let Ships = CreateShips();

export function CreateShips() {
  /** @type {import("$lib/adapter/ships/ships.js").ShipMap | null} */
  let ships = $state(null);

  /** @type {string} */
  let shipsException = $state("");

  return {
    get ships() { return ships },
    error: () => { return shipsException },
    load: async (/** @type {string} */ version) => {
      if (!version) return;
      try {
        shipsException = "";
        ships = null;
        ships = await FetchShips(version);
      } catch (/** @type {any} */ err) {
        shipsException = err.message;
      }
    }
  }
}