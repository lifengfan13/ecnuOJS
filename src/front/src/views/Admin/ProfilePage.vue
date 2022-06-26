<template>
  <div class="profile-page">
    <title-header title="个人主页"></title-header>
    <div>
      <h3>用户名：{{UserInfo.Username}}</h3>
      <br>
      <ul>
        <li>电话：{{UserInfo.Telephone}}</li>
        <li>邮箱：{{UserInfo.Email}}</li>
        <li>单位：{{UserInfo.Organization}}</li>
        <li>地址：{{UserInfo.Address}}</li>
        <li>邮政编码：{{UserInfo.Postcode}}</li>
      </ul>
      <br>
      <el-button type="success" size="small" @click="changeData">修改信息</el-button>
      <el-dialog
        title="修改信息"
        :visible.sync="showChange"
        width="30%" >
        <div>
          <el-form :model="tempUserInfo" ref="tempUserInfo"  :rules="ruleForm"  
            label-width="75px" 
            label-position="top"
            label-suffix=":"
            size="medium"
            status-icon
            inline-message>
            <el-form-item label="用户名" prop="username">
              <el-input type="text" v-model="tempUserInfo.username" placeholder="用户名" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="邮箱" prop="email">
              <el-input type="text" v-model="tempUserInfo.email" placeholder="邮箱" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="单位" prop="organization">
              <el-input type="text" v-model="tempUserInfo.organization" placeholder="单位" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="地址" prop="address">
              <el-input type="text" v-model="tempUserInfo.address" placeholder="地址" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item label="邮政编码" prop="postcode">
              <el-input type="text" v-model="tempUserInfo.postcode" placeholder="邮政编码" autocomplete="off"></el-input>
            </el-form-item>
          </el-form>
          
        </div>
        <span slot="footer" class="dialog-footer">
          <el-button @click="showChange = false">取 消</el-button>
          <el-button type="primary" @click="assgindata('tempUserInfo')">确 定</el-button>
        </span>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import {getUserInfo,changeUserInfo} from '@/request/front/api'
import TitleHeader from '@/components/common/TitleHeader.vue'
export default {
  name: 'ProfilePage',
  data(){
    return {
      showChange:false,
      UserInfo:{},
      tempUserInfo:{
        id:'',
        username:'',
        email:'',
        organization:'',
        address:'',
        postcode:''
      },
      ruleForm: { 
        username: [{  required: true, message: '用户名不能为空', trigger: ['change'] }],
        email: [{  required: true, message: '邮箱不能为空', trigger: ['change'] }],
        organization:[{  required: true, message: '单位不能为空', trigger: ['change'] }],
        address:[{  required: true, message: '地址不能为空', trigger: ['change'] }],
        postcode:[{  required: true, message: '邮政编码不能为空', trigger: ['change'] }],
      }
    }
  },
  components:{
    TitleHeader
  },
  methods:{
    changeData(){
      this.showChange=true
    },
    getInfo(){
      this.tempUserInfo.id=this.$store.state.userId
      getUserInfo({'id':this.$store.state.userId}).then(res => {
        if(res.status == 200) {
          this.UserInfo=res.data 
          this.tempUserInfo.username=this.UserInfo.Username
          this.tempUserInfo.email=this.UserInfo.Email
          this.tempUserInfo.organization=this.UserInfo.Organization
          this.tempUserInfo.address=this.UserInfo.Address
          this.tempUserInfo.postcode=this.UserInfo.Postcode
        }else {
          this.$message.error(res.error.msg)
        }
      }).catch(err =>{
        console.error(err)
      })
    },
    assgindata(formName){
      this.$refs[formName].validate((valid) => {
          if (valid) {
            changeUserInfo(this.tempUserInfo).then(res => {
              if(res.status == 200) {
                this.$message.success("修改成功!")
                
                this.UserInfo.Username=this.tempUserInfo.username
                this.UserInfo.Email=this.tempUserInfo.email
                this.UserInfo.Organization=this.tempUserInfo.organization
                this.UserInfo.Address=this.tempUserInfo.address
                this.UserInfo.Postcode=this.tempUserInfo.postcode
                this.showChange = false
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
    }
  },
  created(){
    this.getInfo()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,h2,h3,h4,h5 {
  margin: 5px 0;
  color: #006798;
}
</style>
