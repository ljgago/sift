<script lang="ts">
  import "../app.css";
  import type { LayoutProps } from "./$types";
  import favicon from "$lib/assets/favicon.svg";
  import { ModeWatcher } from "mode-watcher";
  import Header from "$lib/components/commons/header.svelte";
  import Search from "$lib/components/search.svelte";
  import { goto } from "$app/navigation";
  import { page } from "$app/state";

  let { data, children }: LayoutProps = $props();

  let search = $state({
    name: "",
  });

  function doSearch(q: string) {
    if (q === "") {
      goto("/");
    } else {
      goto(`/search?q=${q}`);
    }
  }

    function handleHome() {
    goto("/");
  }

  function handleSearch() {
    doSearch(search.name);
  }

  function handleKeyDown(event: KeyboardEvent) {
    if (event.key === "Enter") {
      doSearch(search.name);
    }
  }

  function preLoad() {
    search.name = page.url.searchParams.get("q") ?? page.params.slug ?? "";
  }

  $effect(() => {
    preLoad();
  });

  if (!page.params.slug && !page.url.searchParams.get("q")) {
    goto("/");
  }
</script>

<svelte:head><link rel="icon" href={favicon} /></svelte:head>

<ModeWatcher />

<div class="flex flex-col gap-3 mx-12">
  <Header stars={data.stars_count} />
  <Search
    bind:value={search.name}
    onHome={handleHome}
    onSearch={handleSearch}
    onKeyDown={handleKeyDown}
  />

  {@render children()}
</div>
