<template>
  <div class="profile-review">
      <title-header title="审稿队列"></title-header>
      <div class="reviewList">
        <el-table 
          v-loading="listLoading"
          :data="reviewList" 
          highlight-current-row style="width: 100%">
          <el-table-column type="index"></el-table-column>

          <el-table-column label="标题"  align="center">
            <template slot-scope="scope">
              <span>
                <a :href="'http://47.103.81.142:8080/api/v1/submission/paper?id='+scope.row.SubmissionID">
                  {{ articleList[scope.$index] == null ? "" : articleList[scope.$index].Name}}.pdf
                </a>
              </span>
            </template>
          </el-table-column>

          <el-table-column prop="Status" label="评论" align="center" width="500px">
            <template slot-scope="scope">
              <span>
                <el-input type="textarea" :rows="2"
                  placeholder="请输入评论" v-model="scope.row.ReviewComment">
                </el-input>
              </span>
            </template>
          </el-table-column>

          <el-table-column fixed="right" label="操作" witdh="50px" align="center">
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="handleEdit(scope.row.ID,scope.$index)">提交</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
  </div>
</template>

<script>
import TitleHeader from '@/components/common/TitleHeader.vue'

import { mapGetters} from 'vuex'
import {getAllReviewsForUser,getAllReviews,getSubmission,updateReviewcmt} from '@/request/admin/api';

export default {
  name: 'ProfileReview',
  components: {
    TitleHeader
  },
  data () {
    return {
      listLoading:true,
      pageIdx:1,
      reviewList:[
        {
            "ID": 2,
            "CreatedAt": "2022-06-16T00:00:00+08:00",
            "UpdatedAt": "2022-06-16T00:00:00+08:00",
            "DeletedAt": null,
            "ReviewerId": 2,
            "SubmissionID": 13,
            "ReviewComment": "",
            "Status": 0
        },
        {
            "ID": 3,
            "CreatedAt": "2022-06-16T00:00:00+08:00",
            "UpdatedAt": "2022-06-16T00:00:00+08:00",
            "DeletedAt": null,
            "ReviewerId": 2,
            "SubmissionID": 13,
            "ReviewComment": "",
            "Status": 0
        },
      ],
      articleList:[]
    }
  },
  computed: {
    ...mapGetters([
      'getUserId'
    ])
  },
  methods: {
    getAllReviews() {
      let params = {
        id:this.getUserId
      }
      let reviewList = []
      getAllReviewsForUser(params).then(res => {
        if(res.status == 200) {
          reviewList = res.data
          return this.getAllArticles(res.data)
        }else {
          this.listLoading = false
          console.error(res.error.msg)
        }
      }).then(() => {
        this.reviewList = reviewList
        // 这里成功获取了作者信息， loading 结束
        this.listLoading = false
      }).catch(err =>{
        this.listLoading = false
        console.error(err)
      })
    },
    getAllArticles(reviewList) {
      const _this = this
      let submissions = new Set(reviewList.map(v => v.SubmissionID))
      let promises = [];
      function promiseFunc(id){
        return new Promise((resolve, reject) => {
          getSubmission({id:id}).then(res => {
            _this.articleList.push(res.data);
            resolve()
          })
          .catch(err => {
            resolve()
          })
        })
      };
      submissions.forEach(id => {
        if(id){
          promises.push(promiseFunc(id))
        }
      })
      return Promise.all(promises)
    },
    handleEdit(id,idx) {
      let params = {
        id:id,
        comment:this.reviewList[idx].ReviewComment
      };
      updateReviewcmt(params).then(res => {
        if(res.status == 200) {
          this.$message.success("提交成功"); 
        }else {
          this.$message.error("提交失败"); 
          console.error(res.error.msg)
        }
      }).catch(err => {
        this.$message.error("提交失败"); 
        console.error(err)
      })
    }
  },
  created() {
    this.getAllReviews();
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
a {
  text-decoration: none;
  color: #409EFF;
}
</style>
