<script>
  import { cn } from "$lib/utils";
import Icon from "@iconify/svelte";

  let { class: klass, title, message } = $props();

  /** @type {boolean} */
  let showVersionDropdown = $state(false);
</script>

<input type="checkbox" id="exceptionDropdownCheckbox" class="invisible" bind:checked={showVersionDropdown}>
<label for="exceptionDropdownCheckbox" class={cn("exception-box", showVersionDropdown?"h-44":"h-14", showVersionDropdown?"overflow-scroll-hidden":"overflow-hidden", klass)}>
  <div class="flex flex-row gap-4 items-center">
    <span><Icon icon="mdi:error-outline" width="24" height="24" /></span>
    <h1 class="text-2xl">{title}</h1>
  </div>
  <p>{message}</p>
</label>

<style>
  .exception-box {
    @apply flex flex-col gap-2 m-2 p-2 transition-all duration-700 cursor-pointer rounded-lg border-2 border-red-800/80 text-red-800 bg-slate-950/80;
  }
</style>