<template>
  <div class="container">
    <h1 class="title">teeny-url</h1>
    <h2 class="subtitle">Shorten your URLs</h2>
    <div class="field is-grouped">
      <p class="control is-expanded">
        <input v-model="url" v-on:keyup.enter="createShortUrl" class="input" type="url" placeholder="Enter a URL" required>
      </p>
      <p class="control">
        <a v-on:click="createShortUrl" class="button is-primary">Submit</a>
      </p>
    </div>
    <div v-if="shortUrl !== null" class="container">
      <p>Your shortened URL is <a v-bind:href="shortUrl">{{shortUrl}}</a></p>
    </div>
  </div>
</template>

<script>
module.exports = {
  name: 'Home',
  data: function () {
    return {
      url: null,
      shortUrl: null
    }
  },
  methods: {
    createShortUrl: function () {
      this.$http.post('url/', {url: this.url}).then(response => {
        console.log(response)
        this.shortUrl = process.env.REDIRECT_URL + response.body.Key
      }, response => {
        console.error(response)
      })
    }
  }
}
</script>
