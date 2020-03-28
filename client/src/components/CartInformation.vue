<template>
    <div v-if="loaded">
        <div v-if="cart.cartProducts" class="w-80 text-center m-auto">
            <div class="text-center m-5">
                <h2>Your Cart</h2>
                <router-link to="/webstore">Back to store</router-link>
            </div>
            <div class="row j-space-around a-center" v-for="product in cart.cartProducts" :key="product.identifier">
                <img :src="product.thumbnail" />
                <p><strong>{{ product.name }}</strong> - {{ product.option }}</p>
                <p>{{ toPrice(product.price) }} $</p>
                <div class="col-1">
                    <input type="number" v-model.number="product.quantity" min="0" step="1" />
                </div>
                <p><strong>{{ toPrice(product.price * product.quantity) }} $</strong></p>
            </div>
            <input class="btn" type="button" value="Update cart" @click="verifyCart" />
        </div>
        <div v-else class="text-center m-5">
            <h2>Your cart is currently empty</h2>
            <router-link to="/webstore">Back to store</router-link>
        </div>
    </div>
    <div v-else>
        <Spinner />
    </div>
</template>


<script>
import axios from "axios"
import Spinner from "@/components/Spinner.vue"

export default {
    name: "CartInformation",
    components: {
        Spinner
    },
    created () {
        this.verifyCart()
    },
    data () {
        return {
            cart: null,
            loaded: false
        }
    },
    methods: {
        verifyCart() {
            axios.post("/api/cart", {"cartProducts": this.$store.getters.getCart})
                .then(response => {
                    this.cart = response.data.message
                    this.$store.commit("setCart", this.cart.cartProducts)
                    this.loaded = true
                })
        }, 
        toPrice(price) {
            return price/100
        }
    }
}
</script>

<style scoped>
img {
    max-width: 100px;
    width: 100%;
    height: auto;
}
</style>
