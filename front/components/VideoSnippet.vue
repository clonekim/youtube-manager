<template>
 <sui-card>

   <sui-card-content>
     <sui-image :src="item.thumbnails.medium.url" />
     <sui-card-header>
       <nuxt-link :to="`/video/${id}`">{{item.title}}</nuxt-link> </sui-card-header>
     <sui-card-meta>
       <span>채널</span> {{item.channelTitle}}

     </sui-card-meta>
     <sui-card-description>
         {{item.description|summary}}
     </sui-card-description>

     <sui-card-content extra>
       <span slot="right">{{item.publishedAt}}</span>
     </sui-card-content>

   </sui-card-content>


 </sui-card>
</template>

<script>
  export default {
    name: "VideoSnippet",
    props: {
      item: {
        type: Object
      },
      id: {
        type: String
      }
    },

    filters: {
      summary: (value) => {
        if(!value)
          return ""

        if (value.length > 80)
          return `${value.substr(0,80)}...`

        return value
      }
    }
  }
</script>

