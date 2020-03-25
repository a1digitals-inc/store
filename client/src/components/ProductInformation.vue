<template>
    <div v-if="loaded">
        <ImageCarousel :images="this.product.images" />
        <p>{{ this.product.name }}</p>
    </div>
</template>

<script>
import axios from "axios"

import ImageCarousel from "@/components/ImageCarousel.vue"

export default {
    name: "ProductInformation",
    props: ["identifier"],
    components: {
        ImageCarousel
    },
    data () {
        return {
            product: null,
            loaded: false
        }
    },
    created () {
        this.fetchProduct()
    },
    methods: {
        fetchProduct() {
            axios.get("http://localhost:8080/api/product/" + this.identifier)
            .then(response => {
                this.product = response.data.message
                this.loaded = true
            })
        }
    }
}
</script>
