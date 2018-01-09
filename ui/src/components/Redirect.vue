<template>
  <p>Redirecting you to {{fullUrl}}</p>
</template>

<script>
  module.exports = {
    name: 'Redirect',
    data: function () {
      return {
        fullUrl: ''
      }
    },
    created: function () {
      var id = this.$route.params.id

      this.$http.get('url/' + id).then(response => {
        this.fullUrl = response.body.Value
        window.location.replace(this.fullUrl)
      }, response => {
        console.error(response)
      })
    }
  }
</script>