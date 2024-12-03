<script>
  import ActionBar from "$lib/components/ActionBar.svelte";
import Loader from "$lib/components/Loader.svelte";
  import { Manifest } from "$lib/data/manifest.svelte";
  import { Ships } from "$lib/data/ships.svelte";
  import { Versions } from "$lib/data/versions.svelte";
  import Icon from "@iconify/svelte";
  import { onMount } from "svelte";

  onMount(async () => {
    Versions.load();
  })

  $effect(() => {
    Manifest.load(Versions.getLatest());
    Ships.load(Versions.getLatest());
  })
</script>

<div class="flex flex-col items-center min-h-[80vh]">
  <ActionBar></ActionBar>
  <h1 class="title text-7xl sm:text-9xl mt-12" title="Ships">Ships</h1>

  {#if false && Ships.ships}
    <p>Team</p>
  {:else if Versions.error() || Ships.error() || Manifest.error()}
  <p>
    {Ships.error()}
  </p>
  {:else}
  <Loader class="w-36 h-[80vh]"></Loader>
  {/if}
</div>
