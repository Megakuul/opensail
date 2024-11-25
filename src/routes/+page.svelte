<script>
  import Logo from "$lib/components/Logo.svelte";
</script>

<svelte:head>
	<title>Opensail</title>
	<meta name="description" content="Free and Opensource sailing platform." />
  <meta property="og:site_name" content="Opensail" />
	<meta property="og:description" content="Free and Opensource sailing platform." />
	<meta property="og:title" content="Opensail - Home">
  <meta property="og:type" content="website">
	<meta property="og:image" content="https://opensail.battleshiper.dev/favicon.png" />
	<link rel="canonical" href="https://opensail.battleshiper.dev" />
</svelte:head>

<div class="flex flex-col items-center"> 
  <Logo class="w-8/12 mt-10"></Logo>

	<div class="wrapper">
		<div class="box w-96 h-96 ">

		</div>
	</div>

</div>

<style>
	.wrapper {
		perspective: 1000px;
	}
	.box {
		transform: rotateY(12deg);
	}
</style>