import { FetchVersions } from "$lib/adapter/versions.js";

export let Versions = CreateVersions();

export function CreateVersions() {
  /** @type {string} */
  let latest = $state("");

  /** @type {import("$lib/adapter/versions.js").VersionsConfig} */
  let versions = $state([]);

  /** @type {string} */
  let versionsException = $state("");

  return {
    get versions() { return versions },
    error: () => { return versionsException },
    load: async () => {
      if (versions.length != 0) {
        return;
      }
      try {
        versionsException = "";
        versions = (await FetchVersions()).reverse();
        latest = versions.length > 0 ? versions[0] : "";
      } catch (/** @type {any} */ err) {
        versionsException = err.message;
      }
    },
    setLatest: (/** @type {string} */ version) => { latest = version },
    getLatest: () => { return latest; },
  }
}