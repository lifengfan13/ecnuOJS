<template>
  <div>
    <header-bar/>
    <div class="main clearfix">
      <div class="main-content">
        <router-view></router-view>
      </div>
      <div class="sidebar">
        <sidebar-page></sidebar-page>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>

import HeaderBar from '../components/common/HeaderBar'
import Footer from '../components/common/Footer'
import SidebarPage from './Page/SidebarPage'

export default {
  name: 'MainPage',
  components: {
    HeaderBar,
    Footer,
    SidebarPage
  },
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.clearfix:after{
  content: "";
  display: block;
  height: 0;
  clear:both;
  visibility: hidden;
}
.clearfix{
    *zoom: 1;
}

.main {
    width: 1200px;
    background-color: #fff;
    margin: auto;
}
.main-content {
  width: 860px;
  min-height: 900px;
  padding: 20px;
  float: left;
  border-left: 1px solid #D8D8D8;
  border-right: 1px solid #D8D8D8;
}
.sidebar{
  width: 250px;
  padding: 20px;
  float: left;
}
</style>
