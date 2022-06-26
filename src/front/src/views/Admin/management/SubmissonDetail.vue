<template>
  <div class="submisson-detail">
      <title-header title="投稿处理"></title-header>
      <div class="content">

        <!-- 文件信息 -->
        <div class="articleInfo">
          <h4>投稿的文章</h4>
          <el-table
            :data="tableData"
            style="width: 100%">

            <el-table-column label="标题" width="400" align="center">
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
                <span class="">{{ scope.row.Issue == 0 ? statusList[scope.row.Status]:"已发表" }}</span>
              </template>
            </el-table-column>

            <el-table-column label="操作" width="250px" align="center">
              <template slot-scope="scope">
                <el-button type="danger" size="small" :disabled="isAccept" @click="rejectSubmisson()">拒收</el-button>
                <el-button type="success" size="small" :disabled="isAccept" @click="acceptSubmisson()">接收</el-button>
                <el-button type="success" size="small" :disabled="isPublish" @click="pubSubmission">发表</el-button>
              </template>
            </el-table-column>
          </el-table>
          <el-dialog
              title="安排发表"
              :visible.sync="issueDialogVisible"
              width="30%" >
              <div>
                 <h2>期次</h2>
                 <el-select v-model="issueId" placeholder="请选择" size="medium">
                  <el-option
                    v-for="(item, index) in issueList"
                    :key="index"
                    :label='"华东师范大学软件期刊第"+item.ID+"期"'
                    :value="item.ID">
                  </el-option>
                </el-select>
              </div>
              <span slot="footer" class="dialog-footer">
                <el-button @click="issueDialogVisible = false">取 消</el-button>
                <el-button type="primary" @click="assginIssue">确 定</el-button>
              </span>
            </el-dialog>
        </div>

        <div class="reviewInfo">
          <h4>审稿情况</h4>
          <el-table
            :data="reviewList"
            style="width: 100%">

            <el-table-column type="index"></el-table-column>

            <el-table-column label="评审人" width="100">
              <template slot-scope="scope">
                <span>{{ '评审人'+(scope.$index+1)}}</a></span>
              </template>
            </el-table-column>

            <el-table-column label="更新时间" width="100" >
              <template slot-scope="scope">
                <span>
                   {{ scope.row.UpdatedAt.split('T')[0] }}
                </span>
              </template>
            </el-table-column>

            <el-table-column label="状态" align="center"  width="100">
              <template slot-scope="scope">
                <span class="">{{ reviewStatusList[scope.row.Status] }}</span>
              </template>
            </el-table-column>
            
            <el-table-column label="评论"  align="center">
              <template slot-scope="scope">
                <span>{{ scope.row.ReviewComment }}</span>
              </template>
            </el-table-column>


            <el-table-column label="操作" width="200" align="center">
              <template slot-scope="scope">
                <el-button type="danger" size="small" @click="deleteReview(scope.row.ID,scope.$index)">删除</el-button>
              </template>
            </el-table-column>
          </el-table>

          <div class="review-add">
            <el-button type="success" size="small" :disabled="isPublish" @click="addReviewer">添加评审人</el-button>
          </div>
          <el-dialog
              title="分配评审人"
              :visible.sync="reviewDialogVisible"
              width="30%">
              <div class="reviewer-content">
                <el-table
                  :data="reviewerList"
                  style="width: 100%">

                  <el-table-column type="index"></el-table-column>

                  <el-table-column label="评审人">
                    <template slot-scope="scope">
                      <span>{{ scope.row.Username}}</a></span>
                    </template>
                  </el-table-column>

                  <el-table-column label="操作" width="100">
                    <template slot-scope="scope">
                      <el-button type="danger" size="small" @click="assignReviewer(scope.row.ID)">分配</el-button>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="confirmAssgin()">确 定</el-button>
              </span>
            </el-dialog>
        </div>

      </div>
  </div>
</template>

<script>
import TitleHeader from '@/components/common/TitleHeader.vue'

import { 
  getSubmission,getUserInfo,updateArticle,acceptSubmisson,refusesubmission,
  getReviewsBySubmission,delReview,getReviewers,newReview,getAllIssues,issuesubmission
} from '@/request/admin/api';

export default {
  name: 'SubmissonDetail',
  components: {
    TitleHeader
  },
  data () {
    return {
        id: this.$route.query.id,
        submissionInfo:{},
        statusList: ["审核中","拒收","审稿通过"],
        reviewStatusList:['审稿中','评审中','审稿通过'],
        tableData: [],
        authorInfo:{},
        issueDialogVisible:false,
        issueList:[],
        issueId:"  ",
        reviewList:[],
        reviewDialogVisible:false,
        reviewerList:[],
    }
  },
  computed: {
   isAccept() {
    return this.submissionInfo.Status > 0;
   },
   isPublish() {
      return this.submissionInfo.Issue !=0;
   }
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
    rejectSubmisson() {
      let params = {id:this.id}
      refusesubmission(params).then(res => {
        if(res.status == 200) {
          this.$message.success('拒收成功!');
          this.tableData[0].Status = 1;
        }else {
          this.$message.error('拒收失败!');
          console.error(res.error.msg)
        }
      }).catch(err =>{
        this.$message.error('拒收失败!');
        console.error(err)
      })
    },
    acceptSubmisson () {
      let params = {
        id:this.id
      }
      acceptSubmisson(params).then(res => {
        if(res.status == 200) {
          this.$message.success('接收成功!');
          this.tableData[0].Status = 2;
        }else {
          this.$message.error('接收失败!');
          console.error(res.error.msg)
        }
      }).catch(err =>{
        this.$message.error('接收失败!');
        console.error(err)
      })
    },
    pubSubmission() {
      this.issueDialogVisible = true;
      this.getAllIssues();
    },
    getAllIssues() {
      getAllIssues().then(res => {
        if(res.status == 200) {
          this.issueList = res.data
          for(var i=0;i<this.issueList.length;i++){
            if(this.issueList[i].PublishDate!=""){
              this.issueList.splice(i,1)
              i--
            }
          }
        }else {
          console.error(res.error.msg)
        }
      }).catch(err =>{
        console.error(err)
      })
    },
    assginIssue() {
      let params = {
        iid:this.issueId,
        sid:this.id
      }
      issuesubmission(params).then(res => {
        if(res.status == 200) {
          this.$message.success('发表成功!');
          this.submissionInfo.Issue=this.issueId,
          this.issueDialogVisible = false
        }else {
          this.$message.error('发表失败!');
          console.error(res.error.msg)
        }
      }).catch(err =>{
        this.$message.error('发表失败!');
        console.error(err)
      })
    },
    getAllReviews() { 
      let params = {
        id: this.id
      }
      getReviewsBySubmission(params).then(res => {
        if(res.status == 200) {
          this.reviewList = res.data;
        }else {
          console.error(res.error.msg)
        }
      }).catch(err =>{
        console.error(err)
      })
    },
    deleteReview(id,index) { //删除评审
      delReview({id:id}).then(res => {
        if(res.status == 200) {
          this.$message.success('删除成功!');
          this.reviewList.splice(index,1);
        }else {
          this.$message.error('删除失败!');
          console.error(res.error.msg)
        }
      }).catch(err =>{
        this.$message.error('删除失败!');
        console.error(err)
      })
    },
    getReviewers() { //获取所有评审人
      getReviewers().then(res => {
        if(res.status == 200) {
          this.reviewerList = res.data;
        }else {
          console.error(res.error.msg)
        }
      }).catch(err =>{
        console.error(err)
      })
    },
    addReviewer() {
      this.reviewDialogVisible = true;
      this.getReviewers();
    },
    assignReviewer(rid) {
      let params = {
        rid:rid,
        sid:Number(this.id)
      }
      newReview(params).then(res => {
        if(res.status == 200) {
          this.$message.success('分配成功');
        }else {
          this.$message.error('分配失败!');
          console.error(res.error.msg)
        }
      }).catch(err =>{
        this.$message.error('分配失败!');
        console.error(err)
      })
    },
    confirmAssgin() {
      this.getAllReviews();
      this.reviewDialogVisible = false;
    }
  },
  created() {
    this.getSubmission();
    this.getAllReviews();
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
  margin-right: 50px;
  height:50px
}

.review-add button{
  float: right;
}
</style>
