<template>
    <ProductForm :isNew=true />
</template>

<script>
import ProductForm from "@/components/dashboard/ProductForm.vue"

export default {
    name: "DashboardProduct",
    components: {
        ProductForm
    }
}
</script>
