<template>
    <div class="nav-user">
            <span class="dropdown-nv">
              <!-- 未登录 -->
              <el-dropdown v-if="isLogin">
                  <span class="el-dropdown-link">
                    {{userInfo.Username}}<i class="el-icon-caret-bottom el-icon--right"></i>
                  </span>
                  <el-dropdown-menu slot="dropdown">
                      <el-dropdown-item><router-link to="/admin/profile">个人资料</router-link></el-dropdown-item>
                      <el-dropdown-item><router-link to="/admin/submission">操纵面板</router-link></el-dropdown-item>
                      <el-dropdown-item @click.native="logOut()">退出</el-dropdown-item>
                  </el-dropdown-menu>
              </el-dropdown>
              <!-- 登录状态 -->
              <div class="login-nav" v-else>
                <span><router-link to="/main/register">注册</router-link></span>
                <span><router-link to="/main/login">登录</router-link></span>
              </div>
            </span>
    </div>
</template>

<script>
import { mapGetters, mapState,mapMutations } from 'vuex'

export default {
  name: 'NavUser',
  data () {
    return{
      
    }
  },
  computed: {
    ...mapState({
      userInfo : state => state.userInfo,
      isLogin: state => state.isLogin,
    })
  },
  methods: {
    ...mapMutations(['setLogin','delToken','delUserInfo','delUserId']),
    logOut () {
      this.setLogin(false);
      this.delToken();
      this.delUserInfo();
      this.delUserId();
      this.$message.success('用户：' + this.userInfo.userName + ' 退出成功')
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
a{
  color: #000;
  text-decoration:none;
}

.nav-user {
  height: 30px;
}

.dropdown-nv {
  padding: 5px;
  float: right;
}

.el-dropdown-link {
    cursor: pointer;
    color: #fff;
}
.el-icon-arrow-down {
  font-size: 12px;
  color: #fff;
}

.login-nav a{
  text-align: left;
  padding-top: 5px;
  margin: 0 10px;
  color: #fff;
  text-decoration:none;
}
</style>
