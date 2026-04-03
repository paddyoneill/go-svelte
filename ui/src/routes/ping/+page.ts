import type { PingResponse } from "$lib/types";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }): Promise<{ result: PingResponse }> => {
    const response = await fetch('/ping');

    if (!response.ok) {
        return { result: { "result": "failed" } };
    }

    const result = await response.json();

    return result;
};
