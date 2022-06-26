<template>
  <div class="admin-aside">
      <el-menu router
        :default-active="activeIndex"
        @select="handleSelect"
        background-color="#545c64"
        text-color="#fff">
        <!-- 导航条 -->

        <el-submenu index="article">
          <template slot="title">我的投稿</template>
          <el-menu-item index="submission"><template slot="title">新建投稿</template></el-menu-item>
          <el-menu-item index="historySubmission">历史提交</el-menu-item>
        </el-submenu>

        <el-submenu index="review" v-show="userInfo.Role > 1">
          <template slot="title">审稿</template>
          <el-menu-item index="historyReview">审稿队列</el-menu-item>
        </el-submenu>

        <el-submenu index="management" v-show="userInfo.Role > 2">
          <template slot="title">期刊管理</template>
          <el-menu-item index="submissonboard">投稿管理</el-menu-item>
          <el-menu-item index="issue">期次管理</el-menu-item>
        </el-submenu>
        <el-submenu index="user" v-show="userInfo.Role > 2">
          <template slot="title">用户管理</template>
          <el-menu-item index="allusers">用户列表</el-menu-item>
        </el-submenu>
        <el-menu-item index="profile">
          <template slot="title">个人信息</template>
        </el-menu-item>

      </el-menu>
  </div>
</template>

<script>
import {  mapState} from 'vuex'

export default {
  name: 'AdminAside',
  data () {
    return {
      activeIndex: 'submission',
      navBarItem: ['review','submission','article','historySubmission','historyReview','management' ,'issue','submissonboard' ,'profile']
    }
  },
  computed: {
    ...mapState({
      userInfo : state => state.userInfo
    })
  },
  methods: {
    handleSelect (key, keyPath) {
    },
    flushNav: function() {
       this.navBarItem.forEach((item,index) => {
            //检测当前路由路径  然后跳转到对应的item
            if(decodeURI(this.$route.path).search(item) != -1) {
                this.activeIndex = item;
            }
        })
    },
    activeNav: function(index) {
      console.log(index)
      this.activeIndex = index
    }
  },
  created() {
    this.flushNav()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
