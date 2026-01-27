import { error } from "@sveltejs/kit";
import { PUBLIC_BASE_URL } from '$env/static/public';
import type { ResponsePackage } from "$lib/types/package";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params }) => {
  if (params.slug !== "") {
    const response = await fetch(
      `${PUBLIC_BASE_URL}/api/registry/package/${params.slug}`,
    );
    const result: ResponsePackage = await response.json();
    const content = base64ToUtf8(result.data);

    return { content };
  }

  error(404, "Not Found");
};

function base64ToUtf8(data: string) {
  const binaryString = atob(data);
  const bytes = new Uint8Array(binaryString.length);

  for (let i = 0; i < binaryString.length; i++) {
    bytes[i] = binaryString.charCodeAt(i);
  }

  const decoder = new TextDecoder("utf-8");
  return decoder.decode(bytes);
}
