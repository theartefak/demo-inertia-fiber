<script setup>
import { Head } from '@inertiajs/vue3';

defineProps({
    greeting: String,
    users: Object | Array,
});
</script>

<template>
    <Head title="Welcome" />

    {{ greeting }}
</template>
