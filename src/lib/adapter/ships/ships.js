import { AdapterError } from "$lib/adapter/error";

/**
 * @typedef {Object.<string, ShipConfig>} ShipMap
 */

/**
 * @typedef {Object} ShipConfig
 * @property {string} team
 * @property {ShipConfigInfo} boat_info
 * @property {ShipConfigSpec} boat_spec
 * @property {ShipConfigRating} boat_rating
 */

/**
 * @typedef {Object} ShipConfigInfo
 * @property {string} source
 * @property {string} name
 * @property {string} class
 * @property {string} age
 * @property {string} builder
 * @property {string} designer
 */

/**
 * @typedef {Object} ShipConfigSpec
 * @property {string} source
 * @property {ShipConfigDimension} dimension
 * @property {ShipConfigSailArea} sail_area
 * @property {ShipConfigMisc} misc
 */

/**
 * @typedef {Object} ShipConfigDimension
 * @property {number} length_over_all
 * @property {number} draft
 * @property {number} beam
 * @property {number} forestay_height
 * @property {number} wetted_surface_area
 */

/**
 * @typedef {Object} ShipConfigSailArea
 * @property {number} main
 * @property {number} jib
 * @property {number} asymmetric_spinnaker
 * @property {number} symmetric_spinnaker
 */

/**
 * @typedef {Object} ShipConfigMisc
 * @property {number} stability_index
 * @property {number} sailing_displacement
 * @property {number} measured_displacement
 * @property {number} max_crew_weight
 */

/**
 * @typedef {Object} ShipConfigRating
 * @property {string} version
 * @property {number} tcc
 */

/**
 * Fetches the full ShipMap.
 * @param {string} version
 * @returns {Promise<ShipMap>}
 * @throws {AdapterError}
 */
export const FetchShips = async (version) => {
  const res = await fetch(`/api/${version}/ships.json`, {
    method: "GET",
  })
  if (res.ok) {
    return await res.json();
  } else {
    throw new AdapterError(await res.text(), res.status);
  }
}