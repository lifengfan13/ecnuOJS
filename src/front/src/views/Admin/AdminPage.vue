<template>
  <div class="admin-page">
    <el-container>
      <el-header>
        <Header-bar></Header-bar>
      </el-header>
      <el-container>
        <el-aside width="200px">
           <Admin-Aside></Admin-Aside>
        </el-aside>
        <el-main>
          <Admin-Main></Admin-Main>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import HeaderBar from '@/components/Admin/Layout/HeaderBar.vue'
import AdminAside from '@/components/Admin/Layout/AdminAside.vue'
import AdminMain from '@/components/Admin/Layout/AdminMain.vue'

export default {
  name:'AdminPage',
  data () {
    return {

    }
  },
  components:{
    HeaderBar,
    AdminAside,
    AdminMain
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

.el-header {
    padding: 0;
    margin: 0;
}

.el-aside {
  overflow: hidden;
}
</style>
