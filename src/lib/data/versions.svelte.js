import { FetchVersions } from "$lib/adapter/versions.js";

export let Versions = CreateVersions();

export function CreateVersions() {
  /** @type {import("$lib/adapter/versions.js").VersionsConfig} */
  let versions = $state([]);

  /** @type {string} */
  let versionsException = $state("");

  return {
    get versions() { return versions },
    error: () => { return versionsException },
    latest: () => { return versions.length > 0 ? versions[0] : ""; },
    load: async () => {
      if (versions.length != 0) {
        return;
      }
      try {
        versions = await FetchVersions();
        versionsException = "";
      } catch (/** @type {any} */ err) {
        versionsException = err.message;
      }
    }
  }
}