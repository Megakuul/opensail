<script>

</script>

<svg viewBox="0 -10 868 560" id="ship" fill="none" xmlns="http://www.w3.org/2000/svg" {...$$props}>
  <defs>
    <linearGradient id="gradient" x1="0%" y1="0%" x2="100%" y2="0%" gradientUnits="userSpaceOnUse">
      <stop offset="0%" class="start" />
      <stop offset="100%" class="end" />
    </linearGradient>
  </defs>

  <!-- letters -->
  <path id="letter-p" class="letter" d="M574.001 469L574.001 292.321M574.001 292.321C574.001 292.321 634.5 291.895 632.501 248.321C630.909 213.614 584.241 212.411 571.946 212.699C570.301 212.737 569.084 214.143 569.188 215.786L574.001 292.321Z" />
  <path id="letter-e" class="letter" d="M746.03 282.459C746.03 282.459 705.512 300.111 691.176 279.224C685.231 270.564 684.489 263.685 684.452 253.181C684.409 241.155 684.943 232.501 693.038 223.607C706.101 209.254 733.18 211.551 742.7 228.558C745.559 233.665 746.029 242.059 746.067 246.392C746.081 247.961 744.818 249.206 743.249 249.248L693.038 250.604" />
  <path id="letter-n" class="letter" d="M795.91 288.321L795.91 219.956C795.91 218.247 797.335 216.884 799.042 216.959L849.59 219.194C851.194 219.265 852.457 220.586 852.457 222.191L852.458 289.246" />
  
  <!-- hull -->
  <path id="hull" d="M12.1866 497.185L621.528 506.4L617.742 543.5H47.1492L12.1866 497.185Z" />
  <!-- jibsail -->
  <path id="jib" d="M22.9542 458.5L294.723 33.184L299.433 458.5H22.9542Z" />
  <!-- mainsail -->
  <path id="mainsail" d="M381.722 1.72878C380.879 0.61731 379.238 0.774566 378.523 1.97169L95.1718 476.03C93.9857 478.014 95.3955 480.538 97.7073 480.569L614.636 487.403C616.675 487.43 617.601 484.542 615.962 483.329C596.951 469.259 537.651 418.065 504.799 306.847C494.633 313.88 482.297 318 469 318C434.206 318 406 289.794 406 255C406 220.206 434.206 192 469 192C471.19 192 473.355 192.112 475.487 192.33C441.087 83.3147 391.526 14.664 381.722 1.72878Z" />
</svg>

<style>
  #ship {
    transform: translateX(12%);
    animation: animate-ship 5s 1s forwards ease;
  }

  @keyframes animate-ship {
    100% {
      transform: translateX(0);
    }
  }

  #gradient .start {
    stop-color: var(--primary-bg-color);
    stop-opacity: 0.8;
  }

  #gradient .end {
    stop-color: var(--secondary-bg-color);
    stop-opacity: 0.3;
  }

  .letter {
    stroke-dasharray: 1000;
    stroke-dashoffset: 1000;
    filter: drop-shadow(3px 3px 2px rgba(0, 0, 0, 0.2));
    stroke-width: 22;
    stroke-linecap: round;
    stroke: url(#gradient);
  }

  #letter-p {
    animation: animate-letter 5s 1s forwards ease;
  }

  #letter-e {
    animation: animate-letter 5s 2s forwards ease;
  }

  #letter-n {
    animation: animate-letter 5s 3s forwards ease;
  }

  @keyframes animate-letter {
    100% {
      stroke-dashoffset: 0;
    }
  }

  #hull {
    filter: drop-shadow(3px 3px 2px rgba(0, 0, 0, 0.2));
    fill: url(#gradient);
    stroke: var(--secondary-bg-color);
    stroke-opacity: 0.2;
    stroke-width: 2;
    stroke-linecap: round;
    stroke-linejoin: round;
    transition: all ease 1s;
    animation: animate-hull 7s forwards ease;
  }

  #hull:hover {
    cursor: pointer;
    stroke-opacity: 1;
    stroke-width: 4;
  }

  @keyframes animate-hull {
    0% {
      fill-opacity: 0;
      stroke-dasharray: 3000;
      stroke-dashoffset: 3000;
    }
    100% {
      fill-opacity: 1;
      stroke-dashoffset: 0;
    }
  }

  #jib {
    filter: drop-shadow(3px 3px 2px rgba(0, 0, 0, 0.2));
    fill: url(#gradient);
    stroke: var(--secondary-bg-color);
    stroke-opacity: 0.2;
    stroke-width: 2;
    stroke-linecap: round;
    stroke-linejoin: round;
    transition: all ease 1s;
    animation: animate-sail 8s forwards ease-out;
  }

  #jib:hover {
    cursor: pointer;
    stroke-opacity: 1;
    stroke-width: 4;
  }

  #mainsail {
    filter: drop-shadow(3px 3px 2px rgba(0, 0, 0, 0.2));
    fill: url(#gradient);
    stroke: var(--secondary-bg-color);
    stroke-opacity: 0.2;
    stroke-width: 2;
    stroke-linecap: round;
    stroke-linejoin: round;
    transition: all ease 1s;
    animation: animate-sail 6s forwards ease-out;
  }

  #mainsail:hover {
    cursor: pointer;
    stroke-opacity: 1;
    stroke-width: 4;
  }

  @keyframes animate-sail {
    0% {
      fill-opacity: 0;
      stroke-dasharray: 4000;
      stroke-dashoffset: 4000;
    }
    100% {
      fill-opacity: 1;
      stroke-dashoffset: 0;
    }
  }
</style>