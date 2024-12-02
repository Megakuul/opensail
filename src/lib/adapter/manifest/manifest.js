import { AdapterError } from "$lib/adapter/error";

/**
 * @typedef {Object} ManifestConfig
 * @property {string} engineVersion
 * @property {number} timestamp
 */

/**
 * Fetches the ManifestConfig.
 * @param {string} version
 * @returns {Promise<ManifestConfig>}
 * @throws {AdapterError}
 */
export const FetchManifest = async (version) => {
  const res = await fetch(`/api/${version}/manifest.json`, {
    method: "GET",
  })
  if (res.ok) {
    return await res.json();
  } else {
    throw new AdapterError(await res.text(), res.status);
  }
}