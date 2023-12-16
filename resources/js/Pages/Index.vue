<script setup>
import { Head } from '@inertiajs/vue3';

defineProps({
    greeting: String,
});
</script>

<template>
    <Head title="Welcome" />

    {{ greeting }}
</template>
