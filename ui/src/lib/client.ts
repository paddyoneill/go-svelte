import createClient from "openapi-fetch";
import type { paths } from "$lib/api";

export const APIClient = createClient<paths>({ baseUrl: "/api" })
