<template>
    <div v-if="loaded">
        <ImageCarousel :images="this.product.images" />
        <h2>{{ this.product.name }}</h2>
        <p>{{ this.product.description }}</p>
        <div v-if="product.discount < 1">
            <p class="strike">{{ this.product.price }} $</p>
            <p><strong>{{ this.product.price * this.product.discount }} $</strong></p>
        </div>
        <div v-else>
            <p>{{ this.product.price }} $</p>
        </div>
        <div v-if="product.options">
            <select v-model="option">
                <option v-for="option in product.options" :key="option">{{ option }}</option>
            </select>
            <input class="btn" type="button" value="Add to cart" @click="addToCart" />
        </div>
        <div v-else>
            <input class="btn" type="button" value="Soldout" disabled />
        </div>
    </div>
    <div v-else>
        <Spinner />
    </div>
</template>

<script>
import axios from "axios"

import ImageCarousel from "@/components/ImageCarousel.vue"
import Spinner from "@/components/Spinner.vue"

export default {
    name: "ProductInformation",
    props: ["identifier"],
    components: {
        ImageCarousel,
        Spinner
    },
    data () {
        return {
            product: null,
            loaded: false,
            option: "",
        }
    },
    created () {
        this.fetchProduct()
    },
    methods: {
        fetchProduct() {
            axios.get("/api/product/" + this.identifier)
            .then(response => {
                this.product = response.data.message
                this.loaded = true
            })
        },
        addToCart() {
            console.log(this.option)
        }
    }
}
</script>

<style scoped>
div {
    text-align: center;
}

p {
    white-space: pre;
    word-wrap: break-word;
}

.strike {
    text-decoration: line-through;
}

.btn {
    background-color: black;
    color: white;
    border: 1px solid black;
    padding: 10px;
    text-align: center;
}

.btn:hover {
    background-color: white;
    color: black;
}
</style>
