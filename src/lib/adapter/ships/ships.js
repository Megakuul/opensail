import { AdapterError } from "$lib/adapter/error";

/**
 * @typedef {Object.<string, ShipConfig>} ShipMap
 */

/**
 * @typedef {Object} ShipConfig
 * @property {string} team
 * @property {ShipConfigInfo} boat_info
 * @property {ShipConfigBaseSpec} boat_base_spec
 * @property {ShipConfigExtraSpec} boat_extra_spec
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
 * @typedef {Object} ShipConfigBaseSpec
 * @property {string} source
 * @property {ShipConfigBaseSpecDimension} dimension
 * @property {ShipConfigBaseSpecSailArea} sail_area
 */

/**
 * @typedef {Object} ShipConfigBaseSpecDimension
 * @property {number} length_over_all
 * @property {number} draft
 * @property {number} beam
 * @property {number} forestay_height
 * @property {number} wetted_surface_area
 * @property {number} displacement
 * @property {number} crew_weight
 */

/**
 * @typedef {Object} ShipConfigBaseSpecSailArea
 * @property {number} main
 * @property {number} jib
 * @property {number} asymmetric_spinnaker
 * @property {number} symmetric_spinnaker
 */

/**
 * @typedef {Object} ShipConfigExtraSpec
 * @property {string} source
 * @property {ShipConfigExtraSpecDesign} design
 * @property {ShipConfigExtraSpecComposition} composition
 */

/**
 * @typedef {Object} ShipConfigExtraSpecDesign
 * @property {string} mode
 * @property {string} stabilization
 * @property {string} hull
 */

/**
 * @typedef {Object} ShipConfigExtraSpecComposition
 * @property {number} ballast_percentage
 * @property {number} cfk_percentage
 * @property {number} alu_percentage
 * @property {number} gfk_percentage
 * @property {number} wood_percentage
 * @property {number} engine_percentage
 * @property {number} amenity_percentage
 */

/**
 * @typedef {Object} ShipConfigRating
 * @property {string} version
 * @property {number} tcc
 * @property {number} speed_factor
 * @property {number} speed_influence
 * @property {number} speed_drag_points
 * @property {number} speed_upwind_points
 * @property {number} speed_downwind_points
 * @property {number} stabilization_factor
 * @property {number} stabilization_influence
 * @property {number} stabilization_points
 * @property {number} agility_factor
 * @property {number} agility_influence
 * @property {number} agility_points
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