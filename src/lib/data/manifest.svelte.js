import { Versions } from "$lib/data/versions.svelte.js";
import { FetchManifest } from "$lib/adapter/manifest/manifest.js";

/** @type {import("$lib/adapter/manifest/manifest.js").ManifestConfig | null} */
export let Manifest = $state(null);

/** @type {string} */
export let ManifestException = $state("");

/** @param {string} version @throws {AdapterError} */
export const LoadManifest = async (version) => {
  try {
    Manifest = await FetchManifest(version);
    ManifestException = "";
  } catch (/** @type {any} */ err) {
    ManifestException = err.message;
  }
}

$effect(() => {
  if (Versions.latest()) {
    LoadManifest(Versions.latest());
  }
})