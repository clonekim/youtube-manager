<template>

  <sui-grid :columns="3">
    <sui-grid-row>
      <sui-grid-column :width="10">
        <youtube :video-id="this.$route.params.id" ref="youtube"/>
      </sui-grid-column>

      <sui-grid-column>
        <RelatedVideos :items="relatedItems" />
      </sui-grid-column>
    </sui-grid-row>

  </sui-grid>
</template>

<script>
  import RelatedVideos from "@/components/RelatedVideos.vue"
  export default {

    components: {
      RelatedVideos
    },

    computed: {

      items() {
        return this.$store.getters.getVideo
      },

      relatedItems() {
        return this.$store.getters.getRelated
      }
    },

    async fetch() {
      await this.$store.dispatch('fetchVideo', this.$route.params.id)
      await this.$store.dispatch('fetchRelatedVideos', this.$route.params.id)
    }
  }
</script>

