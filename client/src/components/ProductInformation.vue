<template>
    <div v-if="loaded">
        <div class="row j-center my-5">
            <ImageCarousel :images="this.product.images" />
            <div>
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
        </div>
        <div class="m-5">
            <router-link to="/webstore">Back to store</router-link>
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
            option: null,
        }
    },
    created () {
        this.fetchProduct()
    },
    methods: {
        fetchProduct() {
            axios.get(process.env.VUE_APP_ROOT_API + "/api/product/" + this.identifier)
            .then(response => {
                this.product = response.data.message
                if (this.product.options) { 
                    this.option = this.product.options[0]
                }
                this.loaded = true
            })
        },
        addToCart() {
            const product = {
                "identifier": this.identifier,
                "option": this.option,
                "quantity": 1,
            }
            this.$store.commit("addToCart", product)
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

select {
    display: block;
    margin: auto;
    border: 1px solid black;
    background-color: white;
    color: black;
    padding: 10px;
}
</style>
