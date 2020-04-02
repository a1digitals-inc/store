import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import axios from 'axios'

Vue.use(VueRouter)

function validateAndRefreshToken(to, from, next) {
    axios(process.env.VUE_APP_ROOT_API + "/api/refresh", {method:"post", withCredentials: true})
        .then(response => {
            if (response.status == 200) {
                next()
            } else {
                next("/login")
            }
        })
        .catch(() => {
            next("/login")
        })
}

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/about',
        name: 'About',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    },
    {
        path: "/webstore",
        name: "WebStore",
        component: () => import(/* webpackChunkName: "webstore" */ '../views/WebStore.vue')
    },
    {
        path: "/product/:name",
        name: "Product",
        component: () => import(/* webpackChunkName: "product" */ '../views/Product.vue')
    },
    {
        path: "/login",
        name: "Login",
        component: () => import(/* webpackChunkName: "login" */ '../views/Login.vue')
    },
    {
        path: "/cart",
        name: "Cart",
        component: () => import(/* webpackChunkName: "cart" */ '../views/Cart.vue')
    },
    {
        path: "/dashboard",
        name: "Dashboard",
        beforeEnter: validateAndRefreshToken,
        component: () => import(/* webpackChunkName: "dashboard" */ '../views/Dashboard.vue')
    },
    {
        path: "/dashboard/products",
        name: "DashboardProducts",
        beforeEnter: validateAndRefreshToken,
        component: () => import(/* webpackChunkName: "dashboardproducts" */ '../views/DashboardProducts.vue')
    },
    {
        path: "*",
        name: "NotFound",
        component: () => import(/* webpackChunkName: "notFound" */ '../views/NotFound.vue')
    }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
