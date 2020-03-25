<template>
    <div>
        <ProductInformation :identifier="$route.params.name"/> 
    </div>
</template>

<script>
import ProductInformation from "@/components/ProductInformation.vue"

export default {
  name: 'Webstore',
  components: { 
    ProductInformation
  }
}
</script>
