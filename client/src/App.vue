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
<link href="https://fonts.googleapis.com/css?family=Libre+Baskerville|Open+Sans&display=swap" rel="stylesheet">

html {
    font-family: 'Libre Baskerville', serif;
}

p {
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
