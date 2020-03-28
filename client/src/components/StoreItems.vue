<template>
    <div v-if="loaded">
        <div v-if="products.length > 0">
            <StoreItem v-for="product in products" :key="product.identifier" :item="product" />
        </div>
        <div v-else>
            <h2>There are no products to show</h2>
        </div>
    </div>
    <div v-else>
        <Spinner />
    </div>
</template>

<script>
import axios from "axios";
import StoreItem from "@/components/StoreItem.vue"
import Spinner from "@/components/Spinner.vue"

export default {
    name: "StoreItems",
    components: {
        StoreItem,
        Spinner
    },
    data () {
        return {
            products: [],
            loaded: false
        }
    },
    created () {
        this.fetchItems()
    },
    methods: {
        fetchItems() {
            axios.get("/api/products")
            .then(response => {
                this.products = response.data.message
                this.loaded = true
            })
            .catch(error => console.log(error))
        }
    }
}

</script>

<style scoped>
div {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}
</style>
