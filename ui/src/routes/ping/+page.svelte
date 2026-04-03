<script lang="ts">
    import type { PageData } from "./$types";

    let { data }: { data: PageData } = $props();
</script>

<div>
    <p>Got ping response: {data.result}</p>
</div>
