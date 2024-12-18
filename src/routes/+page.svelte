<script>
  import Logo from "$lib/components/Logo.svelte";
  import Icon from "@iconify/svelte";
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

<div class="flex flex-col gap-20 items-center"> 
  <Logo class="w-11/12 lg:w-8/12 mt-10"></Logo>

	<section class="flex flex-col items-start mx-6">
		<h1 class="text-4xl sm:text-8xl font-bold text-slate-100/70 select-none">Opensail</h1>
		<h2 class="text-xl sm:text-4xl text-slate-100/50 ml-4 sm:ml-32 select-none">Your Comprehensive Regatta Rating System</h2>
	</section>

	<section class="flex flex-row gap-8 justify-around items-center w-10/12 sm:w-6/12">
		<a href="/ships" class="flex flex-row items-center justify-center gap-2 w-1/2 p-2 rounded-lg transition-all duration-700 text-slate-100/60 bg-slate-400/20 hover:bg-slate-400/10">
			<span class="font-bold text-lg sm:text-xl">Ships</span>
			<Icon icon="line-md:external-link" width="24" height="24" />
		</a>
		<a href="/teams" class="flex flex-row items-center justify-center gap-2 w-1/2 p-2 rounded-lg transition-all duration-700 text-slate-100/60 bg-slate-400/20 hover:bg-slate-400/10">
			<span class="font-bold text-lg sm:text-xl">Teams</span>
			<Icon icon="line-md:external-link" width="24" height="24" />
		</a>
	</section>

	<section class="relative flex flex-col gap-4 mx-4 p-4 mt-0 lg:mt-20 max-w-[1100px] text-slate-100/60 text-xs sm:text-xl text-justify">
		<div class="absolute hidden xl:block top-[-30px] left-[-40px] text-9xl font-serif text-slate-100/30">“</div>
		<p>
			Opensail is an <a class="font-bold underline" href="https://github.com/megakuul/opensail">opensource</a>
			ship database that offers a <b>refined</b> and <b>sophisticated</b> handicap system for mixed regattas.
		</p>
		<p>
			Let’s be honest... Most mixed inland regattas rely on complex, opaque and often unfair handicap systems.
			Building a system that everyone agrees on? That’s about as easy as assembling a nuclear submarine in your backyard.
		</p>
		<p>
			To end this madness, Openfactor uses <b>generalized</b> but <b>predictable</b> and <b>deterministic</b> methods to calculate handicaps, 
			avoiding overly complex velocity predictions or rule-of-thumb estimations.
		</p>
		<p class="text-center font-bold text-lg sm:text-4xl my-8">«We can't do better, but we can do transparent»</p>
	</section>

	<section class="w-full flex flex-col lg:flex-row-reverse items-center justify-around my-10 sm:my-40 text-slate-100/60">
		<h1 class="text-4xl sm:text-6xl font-bold text-center text-nowrap my-10"><a href="/ships" class="hover:underline">Registered Ships</a></h1>
		<div class="wrapper">
			<div class="box-right w-96 h-96 rounded-lg p-6 bg-slate-900/30 border-[1px] border-slate-100 shadow-xl shadow-slate-100/20">
				<h2 class="text-xl sm:text-2xl font-bold mb-3">Sailing Ship</h2>
				<p class="mb-3">Find <a class="font-bold underline" href="/ships">registered vessels</a> and browse their specifications.</p>
				<center class="my-3 animate-pulse">
					<Icon icon="icon-park-twotone:sailing" width="180" height="180" />
				</center>
				<a href="https://github.com/Megakuul/opensail/blob/main/REGISTER.md" class="w-full block p-2 rounded-lg bg-slate-400/20 hover:bg-slate-400/10 transition-all duration-700">
					Register your Ship
				</a>
			</div>
		</div>
	</section>

	<section class="w-full flex flex-col lg:flex-row items-center justify-around my-10 sm:my-40 text-slate-100/60">
		<h1 class="text-4xl sm:text-6xl font-bold text-center text-nowrap my-10"><a href="/teams" class="hover:underline">Registered Teams</a></h1>
		<div class="wrapper">
			<div class="box-left w-96 h-96 rounded-lg p-6 bg-slate-900/30 border-[1px] border-slate-100 shadow-xl shadow-slate-100/20">
				<h2 class="text-xl sm:text-2xl font-bold mb-3">Crew</h2>
				<p class="mb-3">Find <a class="font-bold underline" href="/teams">registered teams</a> and browse their roles.</p>
				<center class="my-3 animate-pulse">
					<Icon icon="fluent:people-team-16-regular" width="180" height="180" />
				</center>
				<a href="https://github.com/Megakuul/opensail/blob/main/REGISTER.md" class="w-full block p-2 rounded-lg bg-slate-400/20 hover:bg-slate-400/10 transition-all duration-700">
					Register your Crew
				</a>
			</div>
		</div>
	</section>
</div>

<style>
	.wrapper {
		perspective: 1000px;
	}
	.box-right {
		transform: rotateY(6deg);
	}
	.box-left {
		transform: rotateY(-6deg);
	}

	@media (max-width: 640px) {
		.wrapper {
			perspective: none;
		}

		.box-right,
		.box-left {
			transform: rotateY(0deg);
		}
	} 
</style>