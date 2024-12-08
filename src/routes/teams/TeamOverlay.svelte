<script>
  import Icon from "@iconify/svelte";
  import { onMount } from "svelte";

  /**
   * @typedef Props
   * @property {import("$lib/adapter/teams/teams.js").TeamConfig | null} mountedTeam
  */

  /** @type {Props} */
  let { 
    mountedTeam = $bindable()
  } = $props();

  const unmountTeam = () => {
    window.history.pushState({}, "", `${window.location.pathname}`);
    mountedTeam = null;
  }

  /** @param {KeyboardEvent} e */
  const escapeMountedTeam = e => {
    if (e.key === "Escape") unmountTeam();
  }

  onMount(async () => {
    if (document) document.addEventListener("keydown", escapeMountedTeam);
  });
</script>

{#if mountedTeam}
<div class="absolute inset-0 z-40 flex items-center justify-center bg-slate-900/70">
  <div class="relative w-full sm:w-[90%] h-full sm:h-[90%] flex flex-col gap-10 p-4 sm:p-16 rounded-none sm:rounded-lg overflow-scroll-hidden bg-slate-900/90">
    <button class="absolute top-2 right-2 p-2 rounded-lg transition-all duration-700 hover:bg-slate-600/20" onclick="{() => unmountTeam()}">
      <Icon icon="fluent-emoji-flat:cross-mark" width="24" height="24" />
    </button>
    <section class="w-full flex flex-row gap-2 items-center justify-start">
      <h1 class="font-bold text-2xl sm:text-5xl text-slate-100/80">{mountedTeam.name}</h1>
    </section>
    {#each mountedTeam.members as member}
      <section class="w-full flex flex-col gap-2 text-slate-100/70">
        <h1 class="text-lg sm:text-2xl">{member.name}</h1>
        <hr class="border-slate-200/20 border-2 rounded-xl">
        <div class="w-full flex flex-row gap-4">
          {#each member.roles as role}
            {#if role === "bowman"}
              <div class="flex flex-row gap-1 items-center p-2 rounded-lg bg-orange-600/30">
                <Icon icon="streamline:bow-solid" width="18" height="18" />
                <p class="font-bold ">{role}</p>
              </div>
            {:else if role === "trimmer"}
              <div class="flex flex-row gap-1 items-center p-2 rounded-lg bg-emerald-600/30">
                <Icon icon="game-icons:rope-coil" width="22" height="22" />
                <p class="font-bold ">{role}</p>
              </div>
            {:else if role === "skipper"}
              <div class="flex flex-row gap-1 items-center p-2 rounded-lg bg-indigo-600/30">
                <Icon icon="mdi:helm" width="24" height="24" />
                <p class="font-bold ">{role}</p>
              </div>
            {:else}
              <div class="flex flex-row gap-1 items-center p-2 rounded-lg bg-slate-500/20">
                <Icon icon="carbon:person" width="22" height="22" />
                <p class="font-bold ">{role}</p>
              </div>
            {/if}
          {/each}
        </div>
      </section>
    {/each}
  </div>
</div>
{/if}

<style>
  h1, h2, p {
    @apply text-nowrap overflow-hidden;
  }
</style>