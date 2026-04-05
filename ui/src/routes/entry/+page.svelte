<script lang="ts">
    import type { PageProps } from "./$types";

    let { data }: PageProps = $props();
</script>

<table>
    <thead>
        <tr>
            <th>ID</th>
            <th>Entry</th>
        </tr>
    </thead>
    {#if data.entryList}
        <tbody>
            {#each data.entryList as entry}
                <tr>
                    <th>{entry.id}</th>
                    <th>{entry.entry}</th>
                </tr>
            {/each}
        </tbody>
    {/if}
</table>

<div><a href="/entry/create">Create new entry</a></div>
