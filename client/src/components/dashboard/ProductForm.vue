<template>
    <div v-if="isLoaded">
        <form class="w-80 m-auto">
            <div v-if="isNew">
                <h2>New Product</h2>
            </div>
            <input type="text" />
        </form>
    </div>
    <div v-else>
        <Spinner />
    </div>
</template>

<script>
import axios from "axios"
import Spinner from "@/components/Spinner.vue"

export default {
    name: "ProductForm",
    props: ["isNew", "identifier"],
    components: {
        Spinner
    },
    data () {
        return {
            product: {
            
            },
            isLoaded: true
        }
    },
    created () {
        if (!this.isNew) {
            this.isLoaded = false
            this.fetchProduct()
        }
    },
    methods: {
        fetchProduct() {
            axios(
                process.env.VUE_APP_ROOT_API + "/api/admin/products/" + this.identifier,
                {method: "get", withCredentials: true}
            ).then(response => {
                console.log(response.data.message)
            })
        }
    }
}
</script>
