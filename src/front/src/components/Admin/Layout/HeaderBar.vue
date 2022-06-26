<template>
  <div class="el-header-bar">
    <el-row>
      <el-col :span="6">
        <div class="header-logo">
          <router-link to='/'>华东师范大学软件期刊</router-link>
        </div>
      </el-col>
      <el-col :span="4" :push="14" class="header-right">
        <el-dropdown >
          <span class="header-right-user">
            <i class="el-icon-user-solid"></i>
          </span>
          <el-dropdown-menu slot="dropdown" placement="bottom-start">
              <el-dropdown-item><router-link to="/admin/profile">修改资料</router-link></el-dropdown-item>
              <el-dropdown-item @click.native="logOut()">退出</el-dropdown-item>
          </el-dropdown-menu>
      </el-dropdown>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import { mapState,mapMutations } from 'vuex'

export default {
  name: 'HeaderBar',
  data () {
    return {}
  },
  computed: {
    ...mapState({
      userInfo : state => state.userInfo
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
    },
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.el-header-bar {
    background-color: rgb(84, 92, 100);
}

.header-logo {
  padding: 10px;
}

.header-logo a {
    margin-left: 0.5rem;
    margin-right: 0.5rem;
    padding: 0.5rem;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    font-size: 20px;
    font-weight: 700;
    line-height: 1;
    color: #fff;
    text-decoration: none;
    border: 1px solid transparent;
    border-radius: 2px;
}

.header-right {
  text-align: right;
}

.header-right .header-right-user{
  text-align: center;
  margin: 2px 10px;
  width: 50px;
  height: 50px;
}
.header-right .header-right-user i{
  color: #fff;
  font-size: 36px;
}

a{
  text-align: left;
  text-decoration:none;
}
</style>
