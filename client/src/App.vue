<template>
  <div id="app">
    <StoreHeader />
    <router-view/>
  </div>
</template>

<script>
import StoreHeader from "@/components/StoreHeader"

export default {
    name: "App",
    components: {
        StoreHeader
    }
}
</script>

<style>
a {
    color: black;
    text-decoration: none;
}

a:visited {
    color: black;
}
</style>
