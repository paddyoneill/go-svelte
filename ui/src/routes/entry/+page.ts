import { APIClient } from "$lib/client";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
    const response = await APIClient.GET("/entry", { fetch: fetch })

    return { entryList: response.data }
}
