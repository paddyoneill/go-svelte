<script lang="ts">
    import { APIClient } from "$lib/client";
    import type { components } from "$lib/api";
    import { goto } from "$app/navigation";

    let Entry = $state() as components["schemas"]["Entry"];

    async function createNewEntry(event: Event) {
        event.preventDefault();

        const response = await APIClient.POST("/entry", {
            body: Entry,
        });

        if (response.error) {
            return;
        }

        goto("/entry");
    }
</script>

<form
    onsubmit={(e) => {
        createNewEntry(e);
    }}
>
    <label>
        <span>ID</span>
        <input type="number" bind:value={Entry.id} />
    </label>

    <label>
        <span>Entry</span>
        <input type="text" bind:value={Entry.entry} placeholder="entry" />
    </label>

    <button type="submit">Create</button>
    <button
        type="button"
        onclick={() => {
            goto("/entry");
        }}>Cancel</button
    >
</form>
