<script lang="ts">
    import { APIClient } from "$lib/client";
    import type { components } from "$lib/api";
    import { goto } from "$app/navigation";

    let newEntryFormData: components["schemas"]["Entry"] = $state({
        entry: "",
        id: 0,
    });

    async function createNewEntry(event: Event) {
        event.preventDefault();

        const response = await APIClient.POST("/entry", {
            body: newEntryFormData,
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
        <input type="number" bind:value={newEntryFormData.id} />
    </label>

    <label>
        <span>Entry</span>
        <input
            type="text"
            bind:value={newEntryFormData.entry}
            placeholder="entry"
        />
    </label>

    <button type="submit">Create</button>
    <button
        type="button"
        onclick={() => {
            goto("/entry");
        }}>Cancel</button
    >
</form>
