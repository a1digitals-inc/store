<template>
    <div>
        <StoreItem v-for="item in items" :key="item.identifier" :item="item" />
    </div>
</template>

<script>
import axios from "axios";
import StoreItem from "@/components/StoreItem.vue"

export default {
    name: "StoreItems",
    components: {
        StoreItem   
    },
    data () {
        return {
            items: []
        }
    },
    created () {
        this.fetchItems()
    },
    methods: {
        fetchItems() {
            axios.get("http://localhost:8080/api/products")
            .then(response => {
                this.items = response.data.message
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
