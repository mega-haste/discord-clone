<script lang="ts">
    import Icon from "@iconify/svelte";

    export let icon: string;
    // export let text: string = "";
    export let href: string = "/";
</script>

<a {href} class="btn text-left justify-start">
    <span>
        <!-- svelte-ignore missing-declaration -->
        <Icon {icon} />
    </span>
    <span>
        <slot />
    </span>
</a>
