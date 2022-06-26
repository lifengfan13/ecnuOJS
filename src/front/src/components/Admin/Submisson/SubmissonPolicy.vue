<template>
  <div class="submisson-policy">
      <p>
        所有作者必须同意以下声明才能向期刊提交文章。如果您不同意这些条款，您将无法继续提交。
      </p>
      <el-divider></el-divider>
      <div class="requirement">
        <h4>投稿要求</h4>
        <p>
          在继续之前，您必须阅读并确认您已完成以下要求。
        </p>
        <el-checkbox-group class="requireList" v-model="checkList">
          <el-checkbox v-for="(item,index) in requireList" :key="index" :label="index">
            {{item.description}}
          </el-checkbox>
        </el-checkbox-group>
      </div>
      <div class="comment">
        <h4>给编辑的留言</h4>
        <el-input type="textarea" :rows="4" placeholder="请输入内容" :value="comment"></el-input>
      </div>
      <div class="privacy_statement">
        <el-checkbox v-model="privacy_check" >
            是的，我同意根据隐私声明收集和存储我的数据。
        </el-checkbox>
      </div>
  </div>
</template>

<script>


export default {
  name: 'SubmissonPolicy',
  data () {
    return {
      requireList: [
        {index: 0,description:"提交的内容以前没有发表过，也没有提交给其他期刊供考虑（或者在给编辑的评论中提供了解释）。"},
        {index: 1,description:"提交文件为 OpenOffice、Microsoft Word 或 RTF 文档文件格式。"},
        {index: 2,description:"在可用的情况下，已提供参考的 URL。"},
        {index: 3,description:"文本是单行距的；使用 12 号字体；使用斜体，而不是下划线；并且所有插图、图表和表格都放在正文中的适当位置，而不是放在最后。"},
        {index: 4,description:"文本符合作者指南中概述的文体和书目要求。"}
      ],
      comment:"",
      checkList:[ ],
      privacy_check: false
    }
  },
  methods: {
     nextCheck() {
       if (!(this.checkList.length == 5 && this.privacy_check)) {
         this.$message.warning("所有选项必选全选")
         return false;
       }  
       return true;
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

.requirement,.comment {
  margin: 10px 0;
}

.requirement .requireList label{
  display: block;
}

</style>
