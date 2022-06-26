<template>
  <div class="article-info">
       <h4>基本信息</h4>
       <el-divider></el-divider>
       <el-form :model="authorForm" ref="authorForm"  :rules="ruleForm"  
        label-width="75px" 
        label-position="top"
        label-suffix=":"
        size="medium"
        status-icon
        inline-message>
        <el-form-item label="作者" prop="author">
          <el-input type="text" v-model="authorForm.author" placeholder="姓名" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="通讯作者" prop="cauthor">
          <el-input type="text" v-model="authorForm.cauthor" placeholder="姓名" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="第二作者">
          <el-input type="text" v-model="authorForm.sauthor" placeholder="姓名" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="第三作者">
          <el-input type="text" v-model="authorForm.tauthor" placeholder="姓名" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="稿件领域">
          <el-radio v-for="(item,index) in topicList" :key="index" v-model="authorForm.topic" :label="index">{{item}}</el-radio>
        </el-form-item>
        <el-form-item label="作者单位" prop="organization">
          <el-input type="text" v-model="authorForm.organization" placeholder="单位" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="通讯地址" prop="address">
          <el-input type="text" v-model="authorForm.address" placeholder="通讯地址" autocomplete="off"></el-input>
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
          author:"",
          cauthor:"",
          sauthor:"",
          tauthor:"",
          topic: 1,
          organization:"",
          address:""
      }
    }
  },
  data () {
    return {
      topicList: ["区块链","大数据与社会计算","云计算与服务计算","群智协同计算","中文信息处理与检索","数据挖掘",
        "大数据","数据科学与工程","边缘计算","智能服务","机器学习","智能移动终端","网络空间安全","社交网络","模式识别",
        "隐私保护","视频处理","计算机视觉","虚拟现实","物联网关键理论与技术","先进计算与系统","硬件安全","自然语言处理"
      ],
      authorForm: this.value,
      ruleForm: { 
        author: [{  required: true, message: '作者不能为空', trigger: ['change'] }],
        cauthor: [{  required: true, message: '通讯作者不能为空', trigger: ['change'] }],
        organization:[{  required: true, message: '单位不能为空', trigger: ['change'] }],
        address:[{  required: false, trigger: ['change'] }],
      }
    }
  },
  methods: {
     nextCheck() {
       var ret = false;
       this.$refs.authorForm.validate((valid) => {
          if (valid) {
            this.$emit('getChildData',this.authorForm)
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