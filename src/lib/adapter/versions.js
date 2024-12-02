import { AdapterError } from "$lib/adapter/error";

/**
 * @typedef {string[]} VersionsConfig
 */

/**
 * Fetches the VersionsConfig.
 * @returns {Promise<VersionsConfig>}
 * @throws {AdapterError}
 */
export const FetchVersions = async () => {
  const res = await fetch(`/api/versions.json`, {
    method: "GET",
  })
  if (res.ok) {
    return await res.json();
  } else {
    throw new AdapterError(await res.text(), res.status);
  }
}