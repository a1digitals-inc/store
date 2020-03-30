<template>
    <div v-if="loaded">
        <div v-if="cart.cartProducts" class="w-80 m-auto text-center">
            <div class="text-center m-5">
                <h2>Your Cart</h2>
            </div>
            <div class="row j-space-around a-center" v-for="product in cart.cartProducts" :key="product.identifier">
                <img :src="product.thumbnail" />
                <div class="col-2">
                    <p><strong>{{ product.name }}</strong> - {{ product.option }}</p>
                </div>
                <div>
                    <p>Price</p>
                    <p>{{ toPrice(product.price) }} $</p>
                </div>
                <div class="col-1">
                    <p>Quantity</p>
                    <input class="m-2" type="number" v-model.number="product.quantity" min="0" step="1" />
                </div>
                <div>
                    <p>Total</p>
                    <p><strong>{{ toPrice(product.price * product.quantity) }} $</strong></p>
                </div>
            </div>
            <p class="m-5"><strong>Total: {{ toPrice(cart.total) }} $</strong></p>
            <div>
                <input class="btn" type="button" value="Update cart" @click="verifyCart" />
                <router-link class="btn" tag="button" to="/checkout">Checkout</router-link>
            </div>
            <div class="m-5">
                <router-link to="/webstore">Back to store</router-link>
            </div>
        </div>
        <div v-else class="text-center m-5">
            <h2>Your cart is currently empty</h2>
            <router-link class="m-5" to="/webstore">Back to store</router-link>
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
                    console.log(response.data.message)
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

input[type=number] {
    width: 30px;
}
</style>
