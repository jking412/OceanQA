import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const routes = [
    {
        path: '/',
        redirect: {
            name: 'Show',
        },
        name: 'Home',
        component: () => import('../components/MainHome.vue'),
        children: [
            {
                path: '/',
                name: 'Show',
                component: () => import('../components/ShowQuestion.vue')
            },
            {
                path: '/add',
                name: 'Add',
                component: () => import('../components/AddQuestion.vue')
            },
            {
                path: '/question/:id',
                name: 'Detail',
                component: () => import('../components/DetailQuestion.vue')
            }
        ],
    },
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes,
});

export default router