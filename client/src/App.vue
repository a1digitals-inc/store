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
@import url('https://fonts.googleapis.com/css?family=Libre+Baskerville|Open+Sans&display=swap');
@import "./main.css";

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

.text-center {
    text-align: center;
}

.m-3 {
    margin: 1em;
}

.m-5 {
    margin: 3em;
}
</style>
