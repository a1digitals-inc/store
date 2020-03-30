<template>
    <nav>
        <li><router-link class="m-3" to="/dashboard/products">PRODUCTS</router-link></li>
        <li><router-link class="m-3" to="/dashboard/orders">ORDERS</router-link></li>
        <li><router-link class="m-3" to="/dashboard/settings">SETTINGS</router-link></li>
    </nav>
</template>

<script>
export default {
    name: "DashboardPage"
}
</script>

<style scoped>
nav {
    text-align: center;
    list-style-type: none;
}
</style>
