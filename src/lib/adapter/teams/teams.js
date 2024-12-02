import { AdapterError } from "$lib/adapter/error";

/**
 * @typedef {Object.<string, TeamConfig>} TeamMap
 */

/**
 * @typedef {Object} TeamConfig
 * @property {string} name
 * @property {TeamConfigMember[]} members
 */

/**
 * @typedef {Object} TeamConfigMember
 * @property {string} name
 * @property {string[]} roles
 */

/**
 * Fetches the full TeamMap.
 * @param {string} version
 * @returns {Promise<TeamMap>}
 * @throws {AdapterError}
 */
export const FetchTeams = async (version) => {
  const res = await fetch(`/api/${version}/teams.json`, {
    method: "GET",
  })
  if (res.ok) {
    return await res.json();
  } else {
    throw new AdapterError(await res.text(), res.status);
  }
}