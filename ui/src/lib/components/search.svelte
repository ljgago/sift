<script lang="ts">
  import HomeIcon from "@lucide/svelte/icons/home";
  import SearchIcon from "@lucide/svelte/icons/search";
  import * as InputGroup from "$lib/components/ui/input-group";
  import { Button } from "$lib/components/ui/button";

  let { value = $bindable(), onHome, onSearch, onKeyDown } = $props();
</script>

<div class="flex gap-2 place-content-center">
  <Button variant="outline" onclick={onHome}><HomeIcon /></Button>

  <InputGroup.Root class="max-with">
    <InputGroup.Input placeholder="Search packages" bind:value onkeydown={onKeyDown} />
    <InputGroup.Addon>
      <SearchIcon />
    </InputGroup.Addon>
  </InputGroup.Root>

  <Button variant="outline" onclick={onSearch}>Search</Button>
</div>
