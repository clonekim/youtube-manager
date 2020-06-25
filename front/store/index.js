export const state = () => ({

  items: [],
  meta: {}
})

export const actions = {

  async fetchPopularVideos({commit}, params ={}) {
    console.log('fetching videos...')
    const res = await this.$axios.get('/api/popular', {params: params})
    commit('mutatePopularVideos', res.data)
  }
}


export const mutations = {

  mutatePopularVideos(state, payload) {
    state.items = payload.items ? state.items.concat(payload.items) : []
    state.meta = {
      pageInfo:  payload.pageInfo,
      nextPageToken: payload.nextPageToken
    }
  }
}

export const getters = {
  getPopularVideos(state) {
    return state.items
  },

  getMeta(state) {
    return state.meta
  }
}
