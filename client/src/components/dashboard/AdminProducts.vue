<template>
    <div class="text-center w-80 m-auto">
        <h2>Store Products</h2>
        <div v-if="loaded">
            <ul>
                <li v-for="product in products" :key="product.identifier">{{ product.name }}</li>
            </ul>
            <router-link to="/dashboard/product"><button class="btn">New Product</button></router-link>
        </div>
        <div v-else>
            <Spinner />
        </div>
    </div>
</template>

<script>
import axios from "axios"
import Spinner from "@/components/Spinner.vue"

export default {
    name: "AdminProducts",
    components: {
        Spinner
    },
    data () {
        return {
            products: [],
            loaded: false
        }
    },
    created () {
        this.fetchProducts()
    },
    methods: {
        fetchProducts() {
            axios(
                process.env.VUE_APP_ROOT_API + "/api/admin/products",
                {method: "get", withCredentials: true}
            ).then(response => {
                this.products = response.data.message
                this.loaded = true
            })
        }
    }
}
</script>
