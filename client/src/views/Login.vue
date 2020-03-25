<template>
    <div>
        <input type="password" @keyup.enter="submit" v-model="password" />
    </div>
</template>

<script>
import axios from "axios"

export default {
  name: 'Login',
  data () { 
      return {
        password: "",
      }
  },
  methods: {
      submit() {
          axios("http://localhost:8080/api/login", {method: "post", data: {password: this.password}, withCredentials: true})
          .then(response => {
              if (response.data.message == "Authenticated") {     
                this.$router.push("/dashboard")
              }
          })
          this.password = ""
      }
  }
}
</script>

<style scoped>
div {
    margin-top: 10%;
    text-align: center;
}
</style>
