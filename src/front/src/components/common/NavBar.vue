<template>
  <div class="navBar">
      <div class="header-logo" >
        <span>华东师范大学软件期刊</span>
      </div>
      <el-menu 
        :default-active="activeIndex"
        @select="handleSelect"
        router
        mode="horizontal"
        background-color="#1e6292"
        text-color="#fff"
        active-text-color="#ffd04b">
        <!-- 导航条 -->
        <el-menu-item v-for="(item,index) in navBarItem" :key="index" 
          v-bind:index="item.path">
          {{item.title}}
        </el-menu-item>
        <!-- 下拉导航条 -->
        <el-submenu index="/about">
          <template slot="title">关于我们 </template>
          <el-menu-item v-for="(item,index) in subMenuItem" :key="index" 
            v-bind:index="item.path">
            {{item.title}}
          </el-menu-item>
        </el-submenu>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: 'navBar',
  data () {
    return {
      activeIndex: 'home',
      navBarItem: [
        { title:'主页',path:'home' },
        { title:'最新一期',path:'current' },
        { title:'归档',path:'archives' },
        { title:'公告',path:'announcements' }, 
        { title:'搜索',path:'search' }, 
      ],
      subMenuItem: [
        { title:'关于期刊',path:'about' },
        { title:'投稿',path:'submission' },
        { title:'道德规范',path:'privacy' },
        { title:'编辑团队',path:'editorial' },
        { title:'联系我们',path:'contact' }
      ]
      
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    },
    flushNav: function() {
       this.navBarItem.forEach((item,index) => {
            //检测当前路由路径  然后跳转到对应的item
            if(decodeURI(this.$route.path).search(item.path) != -1) {
                this.activeIndex = item.path;
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
.navBar {
    color: #000;
    font-family: "微软雅黑";
    width: 100%;
    background: #1e6292;
    opacity: 0.9;
}

.header-logo {
  margin: 20px 0px 10px 20px;
  font-weight: 700;
  font-size: 28px;
  color: #fff;
}

.header-logo a{
  margin: 20px 0px 0px 0px;
  font-weight: 700;
  font-size: 28px;
  color: #fff;
}
</style>
