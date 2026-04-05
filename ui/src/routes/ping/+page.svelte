<script lang="ts">
    import type { PageProps } from "./$types";

    let { data }: PageProps = $props();
</script>

{#if data.pong}
    <div>
        <p>Got {data.pong.result} response from ping</p>
    </div>
{/if}
