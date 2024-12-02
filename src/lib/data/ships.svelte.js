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
    load: async () => {
      try {
        ships = await FetchShips(Versions.latest());
        shipsException = "";
      } catch (/** @type {any} */ err) {
        shipsException = err.message;
      }
    }
  }
}