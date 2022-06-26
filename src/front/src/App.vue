<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import MainPage from './views/MainPage.vue'
import storage from '@/model/storage'
import { mapGetters, mapState,mapMutations } from 'vuex'

import {getUserInfo} from '@/request/front/api'

export default {
  name: 'App',
  components: {
    MainPage
  },
  methods: {
    ...mapMutations(['setToken','setLogin','setUserId','setUserInfo']),
    flushState() {
      let token = storage.get("token");
      let isLogin = storage.get("isLogin");
      let userId = storage.get("userId");
      this.setToken(token);
      this.setLogin(isLogin);
      this.setUserId(userId);

      if(!isLogin || userId === undefined || userId === ''){
        return;
      }

      let params = {
        id:userId
      }

      getUserInfo(params).then(res => {
        if(res.status == 200) {
          this.setUserInfo(res.data) 
        }else {
          this.$message.error(res.error.msg)
        }
      }).catch(err =>{
        console.error(err)
      })
    }
  },
  created() {
    this.flushState();

    //在页面刷新时将vuex里的信息保存到sessionStorage里
    window.addEventListener("beforeunload",()=>{
      this.flushState();
    })
  }
}
</script>

<style>
#app {
  font-family: "Noto Sans",-apple-system,BlinkMacSystemFont,"Segoe UI","Roboto","Oxygen-Sans","Ubuntu","Cantarell","Helvetica Neue",sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  font-size: 14px;
  line-height: 1.43rem;
  color: rgba(0,0,0,0.87);
  background: #fff;
  min-width: 1200px;
}

body {
  padding: 0;
  margin: 0;
}
img {
    max-width: 100%;
    width: auto;
    height: auto;
}
</style>
