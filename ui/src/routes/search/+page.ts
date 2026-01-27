import { PUBLIC_BASE_URL } from '$env/static/public';
import type { ResponseList } from "$lib/types/package";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, url }) => {
  const response = await fetch(
    `${PUBLIC_BASE_URL}/api/registry/search?${url.searchParams}`,
  );
  const result: ResponseList = await response.json();

  return result.data;
};
