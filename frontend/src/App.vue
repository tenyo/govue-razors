
<template>
  <div id="app">
    <button type="button" @click="pickRazor()">Get a new razor</button>
    <div v-if="id != null">
      <h1>{{ razors[id].name }}</h1>
      <p>{{ razors[id].summary }}</p>
    </div>
  </div>
</template>

<script>
import Vue from 'vue'

export default {
  name: 'App',
  components: {},
  data() {
    return {
      id: null,
      current: null,
      razors: [],
    }
  },
  methods: {
    pickRazor() {
      let id = 0
      do {
        id = this.getRandomInt(0, this.razors.length-1)
      } while (this.id === id)
      this.id = id
    },
    loadRazors() {
      Vue.axios.get('/razors')
        .then((response) => {
          console.debug(response.data)
          this.razors = response.data
        })
    },
    getRandomInt(min, max) {
      min = Math.ceil(min)
      max = Math.floor(max)
      return Math.floor(Math.random() * (max - min + 1) + min)
    },
  },
  mounted() {
    this.loadRazors()
  }
}
</script>
