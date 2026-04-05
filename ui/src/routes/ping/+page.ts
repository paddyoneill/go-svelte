import type { PageLoad } from "./$types";
import createClient from "openapi-fetch";
import type { paths } from "$lib/api";

const client = createClient<paths>({ baseUrl: "/api" });

export const load: PageLoad = async ({ fetch }) => {
    const response = await client.GET("/ping", {
        fetch: fetch,
    })
    return { pong: response.data }
};
