<template>
    <div>
        <router-link :to="'/product/' + item.identifier">
            <img :src="item.thumbnail" loading="lazy" />
            <p>{{ item.name }}</p>
            <p v-if="item.soldout"><strong>Soldout</strong></p>
        </router-link>
    </div>
</template>

<script>
export default {
    name: "StoreItem",
    props: ["item"]
}
</script>

<style scoped>
div {
    max-width: 100px;
    text-align: center;
    margin: 10px;
}

img {
    max-width: 100%;
}
</style>
