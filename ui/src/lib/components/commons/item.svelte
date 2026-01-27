<script lang="ts">
  import ScaleIcon from "@lucide/svelte/icons/scale";
  import BoxIcon from "@lucide/svelte/icons/box";
  import DownloadIcon from "@lucide/svelte/icons/download";
  import Badge from "$lib/components/ui/badge/badge.svelte";

  interface Props {
    name?: string;
    description?: string;
    keywords?: string[];
    updated?: string;
    version?: string;
    dependents?: string;
    license?: string;
    downloads?: string;
  }

  const props: Props = $props();
</script>

<section class="flex items-center">
  <div class="flex flex-col gap-2 grow">
    <a data-sveltekit-preload-data="hover" href={`/package/${props.name}`}>
      <h3 class="text-neutral-900 dark:text-neutral-300 font-bold text-xl">
        {props.name}
      </h3>
    </a>
    <p class="text-neutral-600 dark:text-neutral-400">
      {props.description}
    </p>
    {#if props.keywords && props.keywords.length > 0}
      <ul
        class="flex gap-2 flex-wrap text-neutral-900 dark:text-neutral-300 overflow-auto h-8"
      >
        {#each props.keywords as keyword}
          <li>
            <a href={`/search?q=keywords:${keyword}`}>
              <Badge class="rounded-sm" variant="secondary">{keyword}</Badge>
            </a>
          </li>
        {/each}
      </ul>
    {/if}
    <div>
      <span
        class="flex items-center-self gap-x-2 text-neutral-600 dark:text-neutral-400 font-mono"
      >
        {props.version}
        {#if props.updated}
          • {props.updated}
        {/if}
        • <BoxIcon size="20" />
        {props.dependents} dependents
        {#if props.license}
          • <ScaleIcon size={20} /> {props.license}
        {/if}
      </span>
    </div>
  </div>

  {#if props.downloads}
    <div
      class="flex items-center gap-x-2 text-neutral-900 dark:text-neutral-200"
    >
      <DownloadIcon size={20} />
      {props.downloads}
    </div>
  {/if}
</section>
