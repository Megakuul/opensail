import { AdapterError } from "$lib/adapter/error";

/**
 * @typedef {Object} ManifestConfig
 * @property {string} engineVersion
 * @property {number} timestamp
 */

/**
 * Fetches the ManifestConfig.
 * @returns {Promise<ManifestConfig>}
 * @throws {AdapterError}
 */
export const FetchManifest = async () => {
  const res = await fetch(`${import.meta.env.VITE_OPENSAIL_DATA_ENDPOINT}/manifest.json`, {
    method: "GET",
  })
  if (res.ok) {
    return await res.json();
  } else {
    throw new AdapterError(await res.text(), res.status);
  }
}