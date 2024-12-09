<script>
  import ActionBar from "$lib/components/ActionBar.svelte";
  import ExceptionBar from "$lib/components/ExceptionBar.svelte";
  import Loader from "$lib/components/Loader.svelte";
  import { Manifest } from "$lib/data/manifest.svelte";
  import { Teams } from "$lib/data/teams.svelte";
  import { Versions } from "$lib/data/versions.svelte";
  import { page } from '$app/stores';
  import { onMount } from "svelte";
  import TeamOverlay from "./TeamOverlay.svelte";
  import Search from "./Search.svelte";
  import { goto } from "$app/navigation";

  /** @type {import("$lib/adapter/teams/teams.js").TeamConfig | null} */
  let mountedTeam = $state(null);

  /** @param {string} name @param {import("$lib/adapter/teams/teams.js").TeamConfig | null} team */
  const mountTeam = (name, team) => {
    goto(`${window.location.pathname}?team=${name}`);
    mountedTeam = team;
  }

  onMount(async () => {
    await Versions.load();

  });

  $effect(() => {
    Manifest.load(Versions.getLatest());
    Teams.load(Versions.getLatest());
  })

  $effect(() => {
    const requestedTeam = $page.url.searchParams.get("team");
    if (requestedTeam) {
      mountedTeam = Teams.teams?.[requestedTeam] ?? null;
    } else {
      mountedTeam = null;
    }
  })
</script>

<svelte:head>
	<title>Opensail | Teams</title>
	<meta name="description" content="Checkout registered Opensail teams!" />
  <meta property="og:site_name" content="Opensail" />
	<meta property="og:description" content="Checkout registered Opensail teams!" />
	<meta property="og:title" content="Opensail - Teams">
  <meta property="og:type" content="website">
	<meta property="og:image" content="https://opensail.battleshiper.dev/favicon.png" />
	<link rel="canonical" href="https://opensail.battleshiper.dev/teams" />
</svelte:head>

<div class="flex flex-col items-center min-h-[100vh]">
  <ActionBar></ActionBar>
  <h1 class="title text-7xl sm:text-9xl mt-12" title="Teams">Teams</h1>

  <Search class="my-10"></Search>

  <TeamOverlay mountedTeam={mountedTeam}></TeamOverlay>

  {#if Teams.teams}
    <div class="flex flex-col gap-6 items-center w-full py-5 max-h-[100vh] overflow-scroll-hidden">
      {#each Object.entries(Teams.teams).slice(0, 20) as [key, value]}
        <button onclick="{() => mountTeam(key, value)}"
          class="relative max-w-[1000px] w-9/12 p-5 flex flex-row gap-2 justify-start items-center cursor-pointer rounded-lg shadow-inner bg-slate-300/20 text-slate-200/70">
          <p class="font-bold text-nowrap overflow-hidden">{value.name}</p>
          <p class="text-nowrap overflow-hidden opacity-65 hidden sm:block"># {key}</p>
        </button>
      {/each}
      {#if Object.entries(Teams.teams).length < 1}
        <div class="w-9/12 flex justify-center text-slate-200/70">
          <p class="text-base sm:text-2xl text-center">Sorry, no teams match your search...</p>
        </div>
      {/if}
    </div>
  {:else if Versions.error()}
    <div class="mt-auto">
      <ExceptionBar class="min-w-[50vw]" title={"Error - Failed to load versions"} message={Versions.error()}></ExceptionBar>
    </div>
  {:else if Manifest.error()}
    <div class="mt-auto">
      <ExceptionBar class="min-w-[50vw]" title={"Error - Failed to load manifest"} message={Manifest.error()}></ExceptionBar>
    </div>
  {:else if Teams.error()}
    <div class="mt-auto">
      <ExceptionBar class="min-w-[50vw]" title={"Error - Failed to load teams"} message={Teams.error()}></ExceptionBar>
    </div>
  {:else}
    <Loader class="w-36 h-[70vh]"></Loader>
  {/if}
</div>

