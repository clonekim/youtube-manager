<template>
  <sui-item-group>
    <sui-item v-for="item in items">
      <nuxt-link :to="`/video/${item.id.videoId}`">
        <sui-item-image size="tiny" :src="item.snippet.thumbnails.medium.url"/>
        <sui-item-content vertical-align="middle">
          <sui-item-header>{{item.snippet.title}}</sui-item-header>
        </sui-item-content>
      </nuxt-link>
    </sui-item>
  </sui-item-group>
</template>

<script>
  export default {
    name: 'RelatedVideo',
    props: {
      items: {
        type: Array
      }
    }
  }
</script>

