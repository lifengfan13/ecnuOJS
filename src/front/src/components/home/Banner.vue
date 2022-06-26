<!--  -->
<template>
  <div class="banner">
    <div class="swiper">
      <el-carousel indicator-position="outside">
        <el-carousel-item v-for="item in 4" :key="item">
          <h3>{{ item }}</h3>
        </el-carousel-item>
      </el-carousel>
    </div>

    <div class="account-body">
      <el-tabs v-model="activeName" stretch>
        <el-tab-pane label="登录" name="login">
          <el-form ref="loginForm" 
            :model="loginForm" 
            :rules="loginRules"
            label-width="auto" 
            status-icon>
            <el-form-item label="账号" prop="username">
              <el-input v-model="loginForm.username" placeholder="邮箱" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
              <el-input type="password" v-model="loginForm.password" placeholder="密码" autocomplete="off"></el-input>
            </el-form-item>
             <el-form-item>
              <el-button type="primary" @click="submitForm('loginForm')">登录</el-button>
              <el-button  @click="resetForm('loginForm')">重置</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="注册" name="register">注册</el-tab-pane>
      </el-tabs>
    </div>
  </div>
</template>

<script>
export default {
  name: "Banner",
  data () {
    return {
      activeName: 'login',
      loginForm: {
        username:"",
        password: "",
      },
      registerForm: {

      },
      loginRules: {
          username: [
            { required: true, message: '请输入邮箱地址', trigger: 'blur' },
            { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }
          ],
          password: [
            { required: true, message: '密码不能为空'},
            {
              pattern: /(?=.*[a-zA-Z0-9]).{6,16}/,  
              message: '请输入6-20位英文字母、数字', 
              trigger: 'blur'
            }
          ]
        }
    };
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  }
}

</script>
<style scoped>
.banner {
  margin: 25px 50px;
  display: flex;
  justify-content:space-between;
}

.swiper {
  width: 800px;
}

.account-body {
  width: 270px;
  padding: 10px 30px;
  margin-left: 50px;
  background-color: #fff;
  border: 1px solid #D8D8D8;
}

.el-carousel__item h3 {
    color: #475669;
    font-size: 18px;
    opacity: 0.75;
    line-height: 300px;
    margin: 0;
  }
  
  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }
  
  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }
</style>