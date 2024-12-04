import { FetchShips } from "$lib/adapter/ships/ships.js";

export let Ships = CreateShips();

export function CreateShips() {
  /** @type {import("$lib/adapter/ships/ships.js").ShipMap | null} */
  let shipsOriginal = $state(null);

  /** @type {import("$lib/adapter/ships/ships.js").ShipMap | null} */
  let shipsBuffer = $state(null);

  /** @type {string} */
  let shipsException = $state("");

  return {
    get ships() { return shipsBuffer },
    error: () => { return shipsException },
    load: async (/** @type {string} */ version) => {
      if (!version) return;
      try {
        shipsException = "";
        shipsOriginal = null;
        shipsBuffer = null;
        shipsOriginal = await FetchShips(version);
        shipsBuffer = shipsOriginal;
      } catch (/** @type {any} */ err) {
        shipsException = err.message;
      }
    },
    replace: (/** @type {import("$lib/adapter/ships/ships.js").ShipMap} */ map) => {shipsBuffer = map},
    reset: () => {shipsBuffer = shipsOriginal}
  }
}