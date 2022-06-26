<template>
  <div class="el-header-bar">
    <div class="header-main">
      <nav-user></nav-user>
      <nav-bar></nav-bar>
    </div>
  </div>
</template>

<script>
import NavBar from './NavBar'
import NavUser from './NavUser'

export default {
  name: 'HeaderBar',
  compoents: {
    NavBar
  },
  components: { NavBar,NavUser }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.el-header-bar {
    background-color: #1e6292;
    border-bottom: 1px solid #ddd;
}

.header-main {
  width: 1200px;
  margin: auto;
}
</style>
