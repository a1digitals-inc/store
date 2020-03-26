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
@import url('https://fonts.googleapis.com/css?family=Open+Sans&display=swap');

html {
    font-family: 'Open Sans', sans-serif;
}
a {
    color: black;
    text-decoration: none;
}

a:visited {
    color: black;
}
</style>
