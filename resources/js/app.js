import { createApp, h } from 'vue';
import { createInertiaApp } from '@inertiajs/vue3';
import { resolvePageComponent } from '@/Plugins/vite/inertia-helpers';
import { ZiggyVue } from '@/Plugins/ziggy/vue';
import DefaultLayout from '@/Layouts/SimpleLayout.vue';

const appName = import.meta.env.VITE_APP_NAME || 'Artefak';

createInertiaApp({
    title: (title) => `${title} - ${appName}`,
    resolve: async (name) => {
        const page = await resolvePageComponent(`./Pages/${name}.vue`, import.meta.glob('./Pages/**/*.vue'));
        page.default.layout ??= DefaultLayout;
        return page;
    },
    setup({ el, App, props, plugin }) {
        return createApp({ render: () => h(App, props) })
            .use(plugin)
            .use(ZiggyVue)
            .mount(el);
    },
    progress: {
        color: '#4B5563',
    },
});
