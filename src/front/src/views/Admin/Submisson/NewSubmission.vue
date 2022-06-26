<template>
  <div class="admin-submission" >
      <h2>作者投稿</h2>
      <el-divider></el-divider>

      <el-steps :active="active" align-center process-status="success" finish-status="success">
        <el-step title="开始" description="基础信息"></el-step>
        <el-step title="步骤2" description="文章信息"></el-step>
        <el-step title="步骤3" description="作者信息"></el-step>
        <el-step title="步骤4" description="文件上传"></el-step>
        <el-step title="结束" description="提交"></el-step>
      </el-steps>

      <div class="content">
        <!-- 步骤一 -->
        <Submisson-Policy v-show="active == 0" ref="step1"></Submisson-Policy>
        <!-- 步骤二 -->
        <Author-Info v-show="active == 1" :value='authorForm' @getChildData='getChildData' ref="step2"></Author-Info>
        <!-- 步骤三 -->
        <Article-Info v-show="active == 2" :value='articleForm' @getChildData='getChildData' ref="step3"></Article-Info>
        <!-- 步骤四 -->
        <File-Upload v-show="active == 3" :value='fileList' @getChildData='getChildData' ref="step4"></File-Upload>
      </div>
      <div class="next">
        <el-button type="success" plain @click="next">
          <i class="el-icon-check"></i>&nbsp;&nbsp;保存并继续
        </el-button>
        <el-button type="primary" plain @click="pre">
          上一步
        </el-button>
      </div>
    </div>
</template>

<script>

import SubmissonPolicy from "@/components/Admin/Submisson/SubmissonPolicy.vue"
import ArticleInfo from "@/components/Admin/Submisson/ArticleInfo.vue"
import AuthorInfo from "@/components/Admin/Submisson/AuthorInfo.vue"
import FileUpload from "@/components/Admin/Submisson/FileUpload.vue"

import { submitArticle } from '@/request/admin/api';

export default {
  name: 'AdminSubmission',
  data () {
    return {
      active: 0,
      articleForm:{
        name: "",
        abstract:"",
        keyword:""
      },
      authorForm: {
        author:"",
        cauthor:"",
        sauthor:"",
        tauthor:"",
        topic:"",
        organization:"",
        address:""
      },
      fileList:[],
    }
  },
  components: {
    SubmissonPolicy,
    ArticleInfo,
    AuthorInfo,
    FileUpload,
  },
  methods: {
    pre() {
      if(this.active > 0) {
        this.active--;
      }
    },
    next() {
      if (this.active < 4)  {
        var step = "step" + (this.active + 1)
        if(this.$refs[step].nextCheck()){
          this.active++;
        }
      }else{
        let submitObj = {
          ...this.articleForm,
          ...this.authorForm,
          file:this.fileList[0].raw,
        }
        this.submit(submitObj)
      }
    },
    getChildData(data){
      switch(this.active) {
        case 1:
          this.authorForm = data;
          console.log(this.authorForm)
          break;
        case 2:
          this.articleForm = data;
          console.log(this.articleForm)
          break;
        case 3:
          this.fileList = data;
          console.log(this.fileList)
          break;

        default:
          console.error('not match data');
      }
    },
    submit(submitObj) {//上传投稿信息
      console.log(submitObj)
      submitArticle(submitObj).then(res => {
          if(res.status == 201) {
            this.$message.success("投稿成功")
          }else {
            this.$message.error(res.error.msg)
          }
        }).catch(err =>{
          console.error(err)
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
p {
    margin: 0;
    font-size: 16px;
    line-height: 1.6;
    text-rendering: optimizeLegibility;
}

h1,h2,h3,h4,h5 {
  margin: 5px 0;
  color: #006798;
}

.content {
 padding: 10px 0;
}

.next {
  height: 50px;
  margin: 10px 0;
}

.next button {
  float: right;
  margin: 0 10px;
}







.upload {
  position: relative;
  padding: 0 20px;
  background: #fff;
}

.upload h2{
  padding: 10px 0;
  margin: 0;
  font-size: 1rem;
  line-height: 1.2;
  color: rgba(0,0,0,0.84);
}

.upload-button {
  padding: 10px 0;
}
.upload-button button {
  float: right;
  margin: 0 10px;
}

</style>
