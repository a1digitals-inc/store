<template>
    <div>
        <Cart />
        <ProductInformation :identifier="$route.params.name"/> 
    </div>
</template>

<script>
import ProductInformation from "@/components/ProductInformation.vue"
import Cart from "@/components/Cart.vue"

export default {
    name: 'Webstore',
    components: { 
        ProductInformation,
        Cart
    }
}
</script>
