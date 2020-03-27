<template>
    <div>
        <router-link to="/cart"><p>Cart ({{ $store.getters.getCartQuantity }})</p></router-link>
    </div>
</template>

<script>
export default {
    name: "Cart"
}
</script>

<style scoped>
div {
    text-align: center;
}
</style>
