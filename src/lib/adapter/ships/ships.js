import { AdapterError } from "$lib/adapter/error";

/**
 * @typedef {Object.<string, ShipConfig>} ShipMap
 */

/**
 * @typedef {Object} ShipConfig
 * @property {string} team
 * @property {ShipConfigInfo} boatInfo
 * @property {ShipConfigSpec} boatSpec
 * @property {ShipConfigRating} boatRating
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
 * @property {ShipConfigSailArea} sailArea
 * @property {ShipConfigMisc} misc
 */

/**
 * @typedef {Object} ShipConfigDimension
 * @property {number} lengthOverAll
 * @property {number} draft
 * @property {number} beam
 * @property {number} forestrayHeight
 * @property {number} wettedSurfaceArea
 */

/**
 * @typedef {Object} ShipConfigSailArea
 * @property {number} main
 * @property {number} jib
 * @property {number} asymmetricSpinnaker
 * @property {number} symmetricSpinnaker
 */

/**
 * @typedef {Object} ShipConfigMisc
 * @property {number} stabilityIndex
 * @property {number} sailingDisplacement
 * @property {number} measuredDisplacement
 * @property {number} maxCrewWeight
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