<template>
    <div>
        <CartPreview />
        <StoreItems /> 
    </div>
</template>

<script>
import StoreItems from "@/components/StoreItems.vue"
import CartPreview from "@/components/CartPreview.vue"

export default {
    name: 'Webstore',
    components: { 
        StoreItems,
        CartPreview 
    }
}
</script>
