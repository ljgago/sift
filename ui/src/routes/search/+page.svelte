<script lang="ts">
  import { Separator } from "$lib/components/ui/separator/index.js";
  import Item from "$lib/components/commons/item.svelte";
  import type { PageProps } from "./$types";
  import dayjs from "dayjs";
  import relativeTime from "dayjs/plugin/relativeTime";
  import Pagination from "$lib/components/pagination.svelte";

  const { data }: PageProps = $props();

  let paginate = $state<{ query: string; page: number; size: number }>({
    query: "",
    page: 1,
    size: 20,
  });

  dayjs.extend(relativeTime);

  $effect(() => {
    paginate.page = data.page;
    paginate.size = data.size;
  });

  const items = $derived(
    data?.items?.map((item) => ({
      name: item?.package?.name,
      description: item?.package?.description,
      keywords: item?.package?.keywords,
      updated: dayjs(item?.updated).fromNow(),
      version: item?.package?.version,
      dependents: item?.dependents,
      license: item?.package?.license,
      downloads: item?.downloads?.monthly?.toLocaleString("en-US"),
    })),
  );
</script>

{#each items as item}
  <div>
    <Separator class="my-2" />
    <Item {...item} />
  </div>
{/each}

<Pagination
  class="my-6"
  count={data.total_items}
  page={data.page}
  perPage={data.size}
/>
