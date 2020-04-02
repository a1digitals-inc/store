<template>
  <nav>
    <li><router-link to="/webstore">WEBSTORE</router-link></li>
    <li><router-link to="/media">MEDIA</router-link></li>
    <li><router-link to="/dashboard">DASHBOARD</router-link></li>
  </nav>
</template>

<script>
export default {
  name: 'Home',
}
</script>

<style scoped>
nav {
    text-align: center;
    list-style-type: none;
}

li {
    margin: 10px;
}
</style>
