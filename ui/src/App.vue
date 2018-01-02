<template>
  <div id="app" class="site">
    <section class="site-content">
      <router-view/>
    </section>
    <footer-bar/>
  </div>
</template>

<script>
import FooterBar from './components/FooterBar'

export default {
  name: 'app',
  components: {
    FooterBar
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.site {
  display: flex;
  min-height: 100vh;
  flex-direction: column;
}

.site-content {
  flex: 1;
}
</style>
