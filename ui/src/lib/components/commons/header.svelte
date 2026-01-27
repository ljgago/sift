<script lang="ts">
  import { Separator } from "$lib/components/ui/separator";
  import ToggleTheme from "$lib/components/commons/toggle-theme.svelte";
  import GithubButton from "$lib/components/commons/github-button.svelte";

  const { stars } = $props<{ stars: number }>();

  function format(num: number) {
    const formatter = new Intl.NumberFormat("en-US", {
      notation: "compact",
      compactDisplay: "short",
    });

    return formatter.format(num);
   }

</script>

<div class="flex h-5 items-center gap-4 text-sm my-3 place-content-center">
  <GithubButton stars={format(stars)} />
  <Separator orientation="vertical" />
  <ToggleTheme />
</div>
