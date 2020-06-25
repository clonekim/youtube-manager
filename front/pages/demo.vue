<template>
  <div>
    <button @click="$fetch">
      Refresh
    </button>
    <template v-if="$fetchState.pending">
      <h1>Loading...</h1>
    </template>
    <template v-else-if="$fetchState.error">
      <h1>
      Error
      </h1>
    </template>
    <template v-else>
      {{posts}}
    </template>
    <p>
      <n-link to="/">
        Home
      </n-link>
    </p>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        posts: []
      }
    },
    async fetch () {
      this.posts = await this.$axios.$get('https://jsonplaceholder.typicode.com/posts')
    },
    fetchOnServer: false
  }
</script>
