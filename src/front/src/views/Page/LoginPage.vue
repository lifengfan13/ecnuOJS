<template>
  <div class="register-page">
      <h1>登录</h1>
      <div class="register-box" style="width: 400px;">
        <el-form :model="loginForm" ref="loginForm" :rules="ruleForm" 
          label-width="75px" 
          label-position="top" 
          label-suffix=":" 
          size="small"
          inline-message>
        <el-form-item label="手机号" prop="telephone">
          <el-input v-model="loginForm.telephone" prefix-icon="el-icon-user-solid" ></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="loginForm.password" prefix-icon="el-icon-lock"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('loginForm')">登录</el-button>
          <el-button @click="resetForm('loginForm')">重置</el-button>
        </el-form-item>
      </el-form>
      </div>
      
  </div>
</template>

<script>
import { login,getUserInfo } from '@/request/front/api';
import { mapMutations } from 'vuex'

export default {
  name: 'RegisterPage',
    data() {
      return {
        loginForm: {
          telephone: '',
          password: '',
        },
        ruleForm: {
          telephone: [
            { required: true, message: '手机不能为空', trigger: 'change' },
            { pattern:/^1[3-9]\d{9}$/, message: '请输入正确的手机号', trigger: 'change' }
          ],
          password: [
            {  required: true, message: '密码不能为空', trigger: 'change' },
            { 
              pattern: /^(?![\d]+$)(?![a-zA-Z]+$)(?![^\da-zA-Z]+$)([^\u4e00-\u9fa5\s]){6,20}$/, 
              message: '请输入6-20位英文字母、数字或者符号（除空格），且字母、数字和标点符号至少包含两种',
              trigger: 'blur'
            }
          ]
        }
      };
    },
    methods: {
      ...mapMutations(['setToken','setLogin','setUserInfo','setUserId']),
      setLoginState(data) {
        this.setToken(data.Token);
        this.setLogin(true);
        this.setUserId(data.UserID)
        //更新用户信息
        let params = {
          id:data.UserID
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
      },
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            login(this.loginForm).then(res => {
              //登录成功
              if(res.status == 200) {
                this.setLoginState(res.data) //设置登录状态
                this.$message.success("登录成功")
                this.$router.push('/') // 跳转到首页
              }else {
                this.$message.error("登录失败")
              }
            }).catch(err =>{
              console.error(err)
            })
          } else {
            this.$message.error('登录失败');
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }
    },
    created() {
     
    }
  }
</script>

<style scoped>
</style>
