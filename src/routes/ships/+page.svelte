<script>
  import { FetchManifest } from "$lib/adapter/manifest/manifest";
  import { onMount } from "svelte";

  /** @type {import("$lib/adapter/manifest/manifest.js").ManifestConfig} */
  let manifest;

  onMount(async () => {
    manifest = await FetchManifest()
  })
</script>

<h1>Ships</h1>

<p>
  {#if manifest}
  {manifest.engineVersion}
  {/if}
</p>