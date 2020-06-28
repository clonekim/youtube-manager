<template>
  <div>

    <sui-header>
      YouTube
    </sui-header>

    <sui-card-group :items-per-row="3">
      <VideoSnippet :item="item.snippet"
                    :id="item.id"
                    v-for="(item, index) in items" :key="index"/>
    </sui-card-group>
    <sui-button primary @click="getMore">More</sui-button>

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
      if (this.$store.getters.getPopularVideos && this.$store.getters.getPopularVideos.length > 0)
        return

      await this.$store.dispatch('fetchPopularVideos', {count: 100})
    },

    methods: {
      getMore() {
        this.$store.dispatch('fetchPopularVideos', {pageToken: this.nextPageToken, count: 100})
      }
    }
  }
</script>

