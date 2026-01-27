<script lang="ts">
  import * as Pagination from "$lib/components/ui/pagination/index.js";
  import { page } from "$app/state";
  import { goto, preloadData } from "$app/navigation";

  const { count, page: currentPage, perPage, ...restProps } = $props();

  function prevNextPage(step: "prev" | "next") {
    let url = new URL(page.url);

    if (step === "prev" && currentPage > 1) {
      url.searchParams.set("page", String(currentPage - 1));
    }

    if (step === "next" && currentPage !== count) {
      url.searchParams.set("page", String(currentPage + 1));
    }

    return url;
  }

  async function prevPagePreload() {
    const url = prevNextPage("prev");
    await preloadData(url.toString());
  }

  function prevPage() {
    const url = prevNextPage("prev");
    goto(url);
  }

  async function nextPagePreload() {
    const url = prevNextPage("next");
    await preloadData(url.toString());
  }

  function nextPage() {
    const url = prevNextPage("next");
    goto(url);
  }

  // async function goToIndexPagePreload(index: number) {
  //   let url = new URL(page.url);
  //   url.searchParams.set("page", String(index));

  //   await preloadData(url.toString());
  // }

  function goToIndexPage(index: number) {
    let url = new URL(page.url);

    url.searchParams.set("page", String(index));

    goto(url);
  }
</script>

<div {...restProps}>
  <Pagination.Root {count} page={currentPage} {perPage}>
    {#snippet children({ pages, currentPage })}
      <Pagination.Content>
        <Pagination.Item>
          <Pagination.Previous
            onclick={prevPage}
            onmouseover={prevPagePreload}
          />
        </Pagination.Item>
        {#each pages as page (page.key)}
          {#if page.type === "ellipsis"}
            <Pagination.Item>
              <Pagination.Ellipsis />
            </Pagination.Item>
          {:else}
            <Pagination.Item>
              <Pagination.Link
                {page}
                isActive={currentPage === page.value}
                onclick={() => goToIndexPage(page.value)}
              >
                {page.value}
              </Pagination.Link>
            </Pagination.Item>
          {/if}
        {/each}
        <Pagination.Item>
          <Pagination.Next onclick={nextPage} onmouseover={nextPagePreload} />
        </Pagination.Item>
      </Pagination.Content>
    {/snippet}
  </Pagination.Root>
</div>
