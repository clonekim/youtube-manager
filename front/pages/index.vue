<template>
  <div>

    <template v-if="$fetchState.pending">
      <h1>로딩....</h1>
    </template>

    <template v-else-if="$fetchState.error">
      <h1>Error</h1>
    </template>

    <template v-else>
      <sui-card-group :items-per-row="3">
        <VideoSnippet :item="item.snippet" v-for="(item, index) in items" :key="index"/>
      </sui-card-group>
      <sui-button primary @click="getMore">More</sui-button>
    </template>


  </div>


</template>

<script>
  import VideoSnippet from "../components/VideoSnippet";

  export default {
    components: {VideoSnippet},
    computed: {
      items() {
        return this.$store.getters.getPopularVideos
      },
      nextPageToken() {
        return this.$store.getters.getMeta.nextPageToken
      }
    },


    async fetch() {
      await this.$store.dispatch('fetchPopularVideos')
    },

    methods: {
      getMore() {
        this.$store.dispatch('fetchPopularVideos', {pageToken: this.nextPageToken})
      }
    }
  }
</script>

