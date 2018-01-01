<template>
  <div class="container">
    <h1 class="title">teeny-url</h1>
    <h2 class="subtitle">Shorten your URLs</h2>
    <div class="field is-grouped">
      <p class="control is-expanded">
        <input v-model="url" class="input" type="url" placeholder="Enter a URL" required>
      </p>
      <p class="control">
        <a v-on:click="createShortUrl" class="button is-primary">Submit</a>
      </p>
    </div>
    <div v-if="shortUrl !== null" class="container">
      <p>Your shortened URL is {{shortUrl}}</p>
    </div>
  </div>
</template>

<script>
module.exports = {
  name: 'Home',
  data: {
    url: null,
    shortUrl: null
  },
  methods: {
    createShortUrl: function () {
      console.log(this.url)
      this.shortUrl = this.url
    }
  }
}
</script>
