<template>
  <div class="submisson-step2">
       <h4>基本信息</h4>
       <el-divider></el-divider>
       <el-form :model="articleForm" ref="articleForm"  :rules="ruleForm"  
        label-width="75px" 
        label-position="top"
        label-suffix=":"
        size="medium"
        status-icon
        inline-message>
        <el-form-item label="标题" prop="name">
          <el-input type="type" v-model="articleForm.name" placeholder="你的文章标题" ></el-input>
        </el-form-item>
        <el-form-item label="摘要" prop="abstract">
          <el-input type="textarea" v-model="articleForm.abstract" rows="6" placeholder="文章摘要" ></el-input>
        </el-form-item>
        <el-form-item label="keyword" prop="keyword">
          <el-input type="type" v-model="articleForm.keyword" placeholder="关键词" ></el-input>
        </el-form-item>
        
      </el-form>
  </div>
</template>

<script>


export default {
  name: 'ArticleInfo',
  props:{
    value:{
        type:Object,//在props接受父组件传递数据
        default:{
          name : "",
          abstract:"",
          keyword:""
        }
    }
  },
  data () {
    return {
      articleForm: this.value,
      ruleForm: { 
        name: [{  required: true, message: '文章标题不能为空', trigger: ['change'] }],
        abstract:[{  required: true, message: '摘要不能为空', trigger: ['change'] }],
        keyword: [{  required: false,  trigger: ['change'] }]
      }
    }
  },
  methods: {
     nextCheck() {
       var ret = false;
       this.$refs.articleForm.validate((valid) => {
          if (valid) {
            this.$emit('getChildData',this.articleForm)
            ret = true
          } else {
            this.$message.error('填写信息不符合要求');
            ret = false;
          }
        });
       return ret;
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


</style>
