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
        <input type="text" v-model="form.email" /><br />
        <input type="text" v-model="form.password" /><br />
        <button>Save</button>
    </form>
</template>
