export const state = () => ({

  items: [],
  relatedItems: [],
  video: {},
  meta: {}
})

export const actions = {

  async fetchPopularVideos({commit}, params ={}) {
    console.log('fetching videos...')
    const res = await this.$axios.get('/api/popular', {params: params})
    commit('mutatePopularVideos', res.data)
  },

  async fetchVideo({commit}, payload) {
    const res = await this.$axios.get(`/api/video/${payload}`)
    commit('mutateVideo', res.data)
  },

  async fetchRelatedVideos({commit}, payload) {
    const res = await  this.$axios.get(`/api/related/${payload}`)
    commit('mutateRelatedVideos', res.data)
  }
}


export const mutations = {

  mutatePopularVideos(state, payload) {
    state.items =  payload.items ? state.items.concat(payload.items) : []
    state.meta = {
      pageInfo:  payload.pageInfo,
      nextPageToken: payload.nextPageToken
    }
  },

  mutateVideo(state, payload) {
    state.video = payload.items[0]
  },

  mutateRelatedVideos(state, payload) {
    state.relatedItems = payload.items
  }
}

export const getters = {
  getPopularVideos(state) {
    return state.items
  },

  getMeta(state) {
    return state.meta
  },

  getVideo(state) {
    return state.video
  },

  getRelated(state) {
    return state.relatedItems
  }

}
