<template>
  <div class="article-detail">
      <title-header title="投稿处理"></title-header>
      <div class="content">

        <!-- 文件信息 -->
        <div class="articleInfo">
          <h3>投稿的文章</h3>
          <el-table
            :data="tableData"
            style="width: 100%">

            <el-table-column label="标题" width="500" align="center">
              <template slot-scope="scope">
                <span><a :href="'http://47.103.81.142:8080/api/v1/submission/paper?id='+id">{{ scope.row.Name }}.pdf</a></span>
              </template>
            </el-table-column>

            <el-table-column label="投稿日期" align="center">
              <template slot-scope="scope">
                <span>
                   {{ scope.row.CreatedAt.split('T')[0] }}
                </span>
              </template>
            </el-table-column>
            
            <el-table-column label="作者" align="center">
              <template slot-scope="scope">
                <span>{{ authorInfo.Username }}</span>
              </template>
            </el-table-column>

            <el-table-column label="状态" align="center">
              <template slot-scope="scope">
                <span>{{ scope.row.Issue == 0 ? statusList[scope.row.Status]:"已发表" }}</span>
              </template>
            </el-table-column>

            <el-table-column label="操作">
              <template slot-scope="scope">
                <el-button type="text" size="small" :disabled="submissionInfo.Status > 0" @click="dialogVisible = true">上传文件</el-button>
                <el-dialog
                  title="上传文件"
                  :visible.sync="dialogVisible"
                  width="30%">
                  <div class="upload">
                    <el-upload drag=""
                      ref="upload" 
                      action="#"
                      accept=".pdf, .docx"
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
                  <span slot="footer" class="dialog-footer">
                    <el-button @click="dialogVisible = false">取 消</el-button>
                    <el-button type="primary" @click="uploadPaper()">确 定</el-button>
                  </span>
                </el-dialog>
              </template>
            </el-table-column>
          </el-table>
        </div>

        <div class="abstract">
          <h3>摘要</h3>
          <p>{{submissionInfo.Abstract}}</p>
        </div>
      </div>
  </div>
</template>

<script>
import TitleHeader from '@/components/common/TitleHeader.vue'

import { getSubmission,getUserInfo,updateSubmission} from '@/request/admin/api';

export default {
  name: 'ArticleDetail',
  components: {
    TitleHeader
  },
  data () {
    return {
        id: this.$route.query.id,
        submissionInfo:{},
        statusList: ["审核中","审稿通过","发表"],
        tableData: [],
        authorInfo:{},
        dialogVisible: false,
        file:null,
    }
  },
  getters: {
   
  },
  methods: {
    getSubmission() {
        let params = {
          id: this.id
        }
        getSubmission(params).then(res => {
          if(res.status == 200) {
            this.submissionInfo = res.data
            this.tableData = [this.submissionInfo]

            this.getUserInfo(this.submissionInfo.FirstAuthor)
          }else {
            console.error(res.error.msg)
          }
        }).catch(err =>{
          console.error(err)
        })
    },
    getUserInfo(id) {
      getUserInfo({id:id}).then(res => {
        if(res.status == 200) {
          this.authorInfo = res.data
        }else {
          console.error(res.error.msg)
        }
      }).catch(err =>{
        console.error(err)
      })
    },
    handleChange(file, fileList) {
      console.log(file)
      this.file = file;
    },
    handleExceed(file, fileList) {
      this.$message.warning("最多只能上传一个文件");
    },
    uploadPaper() {
       if(this.file != null && this.file != undefined) {
          let params = {
            id: this.id,
            file:this.file.raw
          }
          console.log(params)
          updateSubmission(params).then(res => {
            if(res.status == 201) {
              this.$message.success('上传成功!');
              this.dialogVisible = false
            }else {
              console.error(res.error.msg)
              this.$message.eroor('上传失败!');
            }
          }).catch(err =>{
            console.error(err)
            this.$message.error('上传失败!');
          })
        }else{
          this.$message.warning('请先上传文件!');
        }
    }
  },
  created() {
    this.getSubmission();
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
a {
  text-decoration: none;
  color: #409EFF;
}

.review-add {
  margin: 10px 0;
  height:50px
}

.review-add button{
  float: right;
}
</style>
