<script>
  import { cn } from "$lib/utils";
import Icon from "@iconify/svelte";

  let { class: klass, title, message } = $props();

  /** @type {boolean} */
  let showVersionDropdown = $state(false);

  /** @type {HTMLLabelElement}*/
  let exceptionLabel;

  $effect(() => {
    if (!showVersionDropdown) {
      exceptionLabel.scrollTo({top: 0, behavior: "smooth"})
    }
  })
</script>

<input type="checkbox" id="exceptionDropdownCheckbox" class="invisible" bind:checked={showVersionDropdown}>
<label bind:this={exceptionLabel} for="exceptionDropdownCheckbox" class={cn("exception-box", showVersionDropdown?"h-44":"h-12", showVersionDropdown?"overflow-scroll-hidden":"overflow-hidden", klass)}>
  <div class="flex flex-row gap-4 items-center text-red-800">
    <span><Icon icon="mdi:error-outline" width="24" height="24" /></span>
    <h1 class="text-xl text-nowrap">{title}</h1>
  </div>
  <div class="m-1 p-2 rounded-lg text-slate-300/80 bg-red-800/30">
    <p class="text-xs lg:text-base hidden lg:block whitespace-break-spaces break-all">{message}</p>
  </div>
</label>

<style>
  .exception-box {
    @apply flex flex-col gap-2 m-2 p-2 transition-all duration-700 cursor-pointer rounded-lg border-2 border-red-800/80 bg-slate-950/80;
  }
</style>