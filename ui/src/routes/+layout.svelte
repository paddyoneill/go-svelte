<script lang="ts">
	import favicon from "$lib/assets/favicon.svg";
	import "@picocss/pico";

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<nav>
	<ul>
		<li><strong>Go Svelte</strong></li>
	</ul>
	<ul>
		<li><a class="contrast" href="/">Home</a></li>
		<li><a class="contrast" href="/ping">Ping</a></li>
		<li><a class="contrast" href="/entry">Entry</a></li>
	</ul>
</nav>

{@render children()}
