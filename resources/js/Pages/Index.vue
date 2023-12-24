<script setup>
import { Head, useForm } from '@inertiajs/vue3';
import route from '@/Plugins/ziggy';

defineProps({
    greeting: String,
    users: Object | Array,
});

const form = useForm({
    email: '',
    password: '',
    name: '',
});

const submit = () => {
    form.post(route('create.dummy.user'), {
        onFinish: () => form.reset('password'),
    });
};
</script>

<template>
    <Head title="Welcome" />

    {{ greeting }}

    <form @submit.prevent="submit">
        <input type="text" v-model="form.name" /><br />
        <p v-text="form.errors.name" />
        <input type="text" v-model="form.email" /><br />
        <p v-text="form.errors.email" />
        <input type="text" v-model="form.password" /><br />
        <p v-text="form.errors.password" />
        <button>Save</button>
    </form>
</template>
