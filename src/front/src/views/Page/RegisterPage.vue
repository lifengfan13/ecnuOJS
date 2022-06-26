<template>
  <div class="register-page">
      <h1>注册</h1>
      <div class="register-box" style="width: 400px;">
        <el-form :model="registerForm" ref="registerForm" :rules="rules" 
          label-width="75px" 
          label-position="top" 
          label-suffix=":" 
          size="small"
          inline-message>
        <el-form-item label="用户名" prop="username">
          <el-input v-model="registerForm.username" prefix-icon="el-icon-user-solid" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="registerForm.email" prefix-icon="el-icon-message" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="手机" prop="telephone">
          <el-input v-model="registerForm.telephone" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="单位" prop="organization">
          <el-input v-model="registerForm.organization" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="地址" prop="address">
          <el-input v-model="registerForm.address" autocomplete="off"></el-input>
        </el-form-item>
         <el-form-item label="邮政编码" prop="postcode">
          <el-input v-model="registerForm.postcode" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input type="password" v-model="registerForm.password" prefix-icon="el-icon-lock" autocomplete="new-password"></el-input>
        </el-form-item>
        <el-form-item label="验证密码" prop="checkPassword">
          <el-input type="password" v-model="registerForm.checkPassword" prefix-icon="el-icon-lock" autocomplete="new-password"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('registerForm')">立即注册</el-button>
          <el-button @click="resetForm('registerForm')">重置</el-button>
        </el-form-item>
      </el-form>
      </div>
      
  </div>
</template>

<script>

import { register } from '@/request/front/api';


export default {
  name: 'RegisterPage',
   data() {
      // 重复密码验证
      const checkPassword = async(rule, value, callback) => {
        if (value.length < 1) {
          return callback(new Error('重复密码不能为空！'));
        } else if(this.registerForm.password != this.registerForm.checkPassword){
          return callback(new Error('两次输入密码不一致！'));
        }else{
          callback()
        }
      };
      return {
        registerForm: {
          username: '',
          email: '',
          telephone:'',
          organization:'',
          address:'',
          postcode:'',
          password: '',
          checkPassword: ''
        },
        rules: {
          username: [
            { required: true, message: '用户名不能为空', trigger: 'change'},
            { pattern:/^[a-zA-Z0-9;',()_./~!@#$%^&*-]{4,100}$/, message: '名称需要在4-100位，支持数字、字母、英文标点（半角）', trigger: 'change'}
          ],
          email: [
            { required: true, message: '邮箱不能为空', trigger: 'change' },
            { pattern:/^([a-zA-Z0-9]+[-_\.]?)+@[a-zA-Z0-9]+\.[a-z]+$/, message: '请输入合法的邮箱地址', trigger: 'change' }
          ],
          telephone: [
            { required: true, message: '手机不能为空', trigger: 'change' },
            { pattern:/^1[3-9]\d{9}$/, message: '请输入正确的手机号', trigger: 'change' }
          ],
          organization: [
            { required: false,  trigger: 'change' }
          ],
          address: [
            { required: false,  trigger: 'change' }
          ],
          postcode: [
            { required: false,  trigger: 'change' }
          ],
          password: [
            {  required: true, message: '密码不能为空', trigger: 'change' },
            { 
              pattern: /^(?![\d]+$)(?![a-z]+$)(?![A-Z]+$)(?![^\da-zA-Z]+$)([^\u4e00-\u9fa5\s]){6,20}$/, 
              message: '请输入6-20位英文字母、数字或者符号（除空格），且必须包含大小写、数字三种',
              trigger: 'blur'
            }
          ],
          checkPassword: [
            {  required: true, validator: checkPassword, trigger: 'blur' },
          ]
        }
      };
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            register(this.registerForm).then(res => {
              if(res.status == 201) {
                this.$message.success("注册成功!")
              }else {
                console.error(res.error.msg)
                this.$message.error("注册失败！")
              }
            }).catch(err =>{
              if(err) {
                console.error(err.error.msg)
              }
              this.$message.error("注册失败！")
            })     
          } else {
            this.$message.error("你填写的消息有误！")
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
