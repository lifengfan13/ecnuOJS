<template>
  <div class="file-upload">
      <h4>文件上传</h4>
      <el-divider></el-divider>
      <el-upload drag=""
        ref="upload" 
        action="#"
        accept=".pdf, .docx"
        v-model="fileList"
        :on-change="handleChange"
        :on-exceed="handleExceed"
        :auto-upload="false"
        :limit="1">
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
        <div class="el-upload__tip" slot="tip">
          友情提示：请参照本刊投稿模版，系统将自动抽取稿件信息，并对稿件格式做必要的检查，
          如果严重不符合本刊规范，投稿时将增加很多工作量。系统只接收PDF格式文档。
        </div>
      </el-upload>
  </div>
</template>

<script>


export default {
  name: 'FileUpload',
   props:{
    value:{
        type:Array,//在props接受父组件传递数据
        default:[]
    }
  },
  data () {
    return {
      fileList: this.value,
    }
  },
  methods: {
      nextCheck() {
        if(this.fileList.length > 0) {
          this.$emit('getChildData',this.fileList)
          return true
        }
        this.$message.warning('请先上传文件!');
        return false
      },
      handleChange(file, fileList) {
        this.fileList = fileList;
      },

      handleExceed(file, fileList) {
        this.$message.warning("最多只能上传一个文件");
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
