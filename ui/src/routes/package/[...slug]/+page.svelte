<script lang="ts">
  import { onMount } from "svelte";
  import type { PageProps } from "./$types";
  import { codeToHtml } from "shiki";
  import "$lib/assets/github-markdown.css";

  const { data }: PageProps = $props();

  async function highlight() {
    const snippets = document.querySelectorAll<HTMLElement>("pre code");

    for (const node of snippets) {
      const lang = node.className.replace("language-", "") || "text";

      const html = await codeToHtml(node.textContent, {
        lang,
        themes: {
          light: "github-light",
          dark: "github-dark",
        },
      });

      if (!!node && !!node.parentElement) {
        node.parentElement.outerHTML = html;
      }
    }
  }

  onMount(highlight);
</script>

<div class="my-6">
  <article class="markdown-body dark:markdown-body">
    {@html data.content}
  </article>
</div>
