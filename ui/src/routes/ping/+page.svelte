<script lang="ts">
    import type { PageData } from "./$types";

    let { data }: { data: PageData } = $props();
</script>

<div>
    <p>Got {data.pong.result} response from ping</p>
</div>
