<template>
  <div class="issue-detail">
      <title-header :title="'第'+id+'期次信息'"></title-header>
      <div class="content">
        <el-table
            :data="articleList"
            style="width: 100%">
            <el-table-column type="index"></el-table-column>
            <el-table-column label="文章题目" width="400px">
              <template slot-scope="scope">
                  <span><a :href="'http://47.103.81.142:8080/api/v1/submission/paper?id='+scope.row.ID">{{ scope.row.Name }}.pdf</a></span>
              </template>
            </el-table-column>

            <el-table-column label="操作"  align="center">
              <template slot-scope="scope">
                <el-button type="success" size="small" @click="$router.push('submissondetail?id='+ scope.row.ID)">
                    查看
                </el-button>
                <el-button type="danger" size="small" @click="delArticle(scope.row.ID,scope.$index)">删除</el-button>
              </template>
            </el-table-column>
        </el-table>
      </div>
  </div>
</template>

<script>
import TitleHeader from '@/components/common/TitleHeader.vue'
import {getIssueArticle} from '@/request/front/api'
import {delArticlefromIssue} from '@/request/admin/api'
export default {
  name: 'IssueDetail',
   components: {
    TitleHeader
  },
  data () {
    return {
      id: this.$route.query.id,
      articleList:[],
    }
  },
  created(){
    getIssueArticle({'iid':this.id}).then(res => {
        if(res.status == 200) {
          this.articleList = res.data
        }else {
          console.error(res.error.msg)
        }
      }).catch(err =>{
        console.error(err)
      })
  },
  methods:{
    delArticle(aid,idx) {
      let params = {
        'sid':aid,
      }
      delArticlefromIssue(params).then(res => {
        if(res.status == 200) {
          this.$message.success('删除成功');
          this.articleList.splice(idx,1);
        }else {
          this.$message.error('删除失败');
          console.error(res.error.msg)
        }
      }).catch(err =>{
        this.$message.error('删除失败');
        console.error(err)
      })
    }
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
