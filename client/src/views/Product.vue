<template>
    <div>
        <CartPreview />
        <ProductInformation :identifier="$route.params.name"/> 
    </div>
</template>

<script>
import ProductInformation from "@/components/ProductInformation.vue"
import CartPreview from "@/components/CartPreview.vue"

export default {
    name: 'Webstore',
    components: { 
        ProductInformation,
        CartPreview
    }
}
</script>
