<template>
    <div v-if="loaded">
        <StoreItem v-for="item in items" :key="item.identifier" :item="item" />
    </div>
    <div v-else>
        <Spinner />
    </div>
</template>

<script>
import axios from "axios";
import StoreItem from "@/components/StoreItem.vue"
import Spinner from "@/components/Spinner.vue"

export default {
    name: "StoreItems",
    components: {
        StoreItem,
        Spinner
    },
    data () {
        return {
            items: [],
            loaded: false
        }
    },
    created () {
        this.fetchItems()
    },
    methods: {
        fetchItems() {
            axios.get("/api/products")
            .then(response => {
                this.items = response.data.message
                this.loaded = true
            })
            .catch(error => console.log(error))
        }
    }
}

</script>

<style scoped>
div {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}
</style>
