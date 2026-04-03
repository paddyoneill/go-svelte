import type { PageLoad } from "./$types";
import createClient from "openapi-fetch";
import type { paths } from "$lib/api";

const client = createClient<paths>();

export const load: PageLoad = async ({ fetch }) => {
    const resp = await client.GET("/ping")
    if (resp.data === undefined) {
        return {
            pong: {
                result: "",
            }
        }
    }

    return {
        pong: {
            result: resp.data.result,
        }
    }

};
