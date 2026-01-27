import { PUBLIC_BASE_URL } from '$env/static/public';
import type { ResponseStars } from "$lib/types/github-stars";
import type { LayoutLoad } from "./$types";

export const load: LayoutLoad = async ({ fetch }) => {
  const response = await fetch(
    `${PUBLIC_BASE_URL}/api/github/repos/ljgago/sift/stars`,
  );
  const result: ResponseStars = await response.json();

  return result.data;
};

export const ssr = false;
