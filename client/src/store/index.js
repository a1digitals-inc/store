import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        cart: []
    },
    mutations: {
        initStore(state) {
            if (localStorage.getItem("cart")) {
                state.cart = JSON.parse(localStorage.getItem("cart"))
            }
        },
        setCart(state, cartProducts) {
            if (cartProducts != null) {
                state.cart = cartProducts
            } else {
                state.cart = []
            }
            localStorage.setItem("cart", JSON.stringify(state.cart))    
        },
        addToCart(state, product) {
            var exists = false
            for (let i = 0; i < state.cart.length; i++) {
                if (state.cart[i].identifier == product.identifier && state.cart[i].option == product.option) {
                    exists = true
                    state.cart[i].quantity += product.quantity
                }
            }
            if (!exists) {
                state.cart.push(product)
            }
            localStorage.setItem("cart", JSON.stringify(state.cart))
        }
    },
    getters: {
        getCartQuantity(state) {
            var count = 0
            if (state.cart) {
                for (let i = 0; i < state.cart.length; i++) {
                    count += state.cart[i].quantity
                }
            }
            return count
        },
        getCart(state) {
            return state.cart
        }
    }
})
