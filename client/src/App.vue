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

.btn {
    background-color: black;
    color: white;
    border: 1px solid black;
    padding: 10px;
    text-align: center;
    margin: 20px;
}

.btn:hover {
    background-color: white;
    color: black;
} 

.btn:disabled,.btn[disabled],.btn[disabled]:hover {
    color: white;
    background-color: grey;
    border: 1px solid grey;
}
</style>
