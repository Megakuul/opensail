<script>
  import { cn } from "$lib/utils";
  import { fade } from "svelte/transition";

  let { class: klass } = $props();

  let loadingDots = $state(0);

  setInterval(() => {
    loadingDots = loadingDots>2 ? loadingDots = 0 : loadingDots+1;
  }, 800)
</script>

<div class="{cn("flex flex-col gap-6 justify-center items-center", klass)}">
  <svg viewBox="0 -10 630 560" fill="none" xmlns="http://www.w3.org/2000/svg">
    <defs>
      <linearGradient id="gradient" x1="0%" y1="0%" x2="100%" y2="0%" gradientUnits="userSpaceOnUse">
        <stop offset="0%" class="start" />
        <stop offset="100%" class="end" />
      </linearGradient>
    </defs>
  
    <!-- hull -->
    <path id="hull" d="M12.1866 497.185L621.528 506.4L617.742 543.5H47.1492L12.1866 497.185Z" />
    <!-- jibsail -->
    <path id="jib" d="M22.9542 458.5L294.723 33.184L299.433 458.5H22.9542Z" />
    <!-- mainsail -->
    <path id="mainsail" d="M381.722 1.72878C380.879 0.61731 379.238 0.774566 378.523 1.97169L95.1718 476.03C93.9857 478.014 95.3955 480.538 97.7073 480.569L614.636 487.403C616.675 487.43 617.601 484.542 615.962 483.329C596.951 469.259 537.651 418.065 504.799 306.847C494.633 313.88 482.297 318 469 318C434.206 318 406 289.794 406 255C406 220.206 434.206 192 469 192C471.19 192 473.355 192.112 475.487 192.33C441.087 83.3147 391.526 14.664 381.722 1.72878Z" />
  </svg>
  <p class="text-slate-200/60 text-2xl font-bold min-w-32">Loading
    {#each Array(loadingDots) as _}
      <span transition:fade class="mx-[1px]"> . </span>
    {/each}
  </p>
</div>


<style>
  #gradient .start {
    stop-color: var(--primary-bg-color);
    stop-opacity: 1;
  }

  #gradient .end {
    stop-color: var(--secondary-bg-color);
    stop-opacity: 0.3;
  }

  #hull, 
  #jib,
  #mainsail {
    filter: drop-shadow(3px 3px 2px rgba(0, 0, 0, 0.2));
    fill: url(#gradient);
    fill-opacity: 0;

    stroke: rgba(255, 255, 255, 0.3);
    stroke-opacity: 0.8;
    stroke-width: 3;
    stroke-linecap: round;
    stroke-linejoin: round;
    stroke-dasharray: 2000;
    stroke-dashoffset: 2000;

    transition: all ease 1s;
    animation: animate-component 10s 1s infinite ease-in-out;
  }

  #hull:hover {
    cursor: pointer;
    stroke-opacity: 1;
    stroke-width: 4;
  }

  @keyframes animate-component {
    0%, 100% {
      fill-opacity: 0;
      stroke-dasharray: 2000;
      stroke-dashoffset: 2000;
    }
    50% {
      fill-opacity: 1;
      stroke-dashoffset: 0;
    }
  }
</style>