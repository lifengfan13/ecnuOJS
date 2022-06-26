<template>
  <div class="admin-issue">
      <title-header title="期次管理"></title-header>
      <div class="content">
        <el-table
            :data="issueList"
            style="width: 100%">
            <el-table-column type="index"></el-table-column>

            <el-table-column label="期次" width="400px">
              <template slot-scope="scope">
                <router-link :to='"issuedetail?id="+ scope.row.ID'>
                  <span>{{ "华东师范大学软件期刊第"+scope.row.ID+"期" +' ('+(scope.row.CreatedAt.split('-')[0]) +')' }}</span>
                </router-link>  
              </template>
            </el-table-column>

            <el-table-column label="创建时间" >
              <template slot-scope="scope">
                <span>
                   {{ scope.row.CreatedAt.split('T')[0] }}
                </span>
              </template>
            </el-table-column>

            <el-table-column label="状态" align="center" >
              <template slot-scope="scope">
                <span class="">{{ scope.row.Status == 1 ? "已发表":""}}</span>
              </template>
            </el-table-column>

            <el-table-column label="操作"  align="center" width="300px">
              <template slot-scope="scope">
                <el-button type="primary" size="small" @click="$router.push('issuedetail?id='+ scope.row.ID)">
                    查看
                </el-button>
                <el-button type="danger" size="small" @click="delIssue(scope.row.ID,scope.$index)">删除</el-button>
                <el-button type="success" size="small" @click="pubIssue(scope.row.ID,scope.$index)">发布</el-button>
              </template>
            </el-table-column>
        </el-table>

        <div class="issue-add">
          <el-button type="success" size="small" @click="addIssue">添加新一期</el-button>
        </div>
      </div>
  </div>
</template>

<script>
import TitleHeader from '@/components/common/TitleHeader.vue'
import { 
  getAllIssues,delIssue,pubIssue,addIssue
} from '@/request/admin/api';

export default {
  name: 'AdminIssue',
  components: {
    TitleHeader
  },
  data () {
    return {
      issueList:[],
    }
  },
  methods: {
    getAllIssues() {
      getAllIssues().then(res => {
        if(res.status == 200) {
          this.issueList = res.data
        }else {
          console.error(res.error.msg)
        }
      }).catch(err =>{
        console.error(err)
      })
    },
    delIssue(id,idx) {
      let params = {
        id:id
      }
      delIssue(params).then(res => {
        if(res.status == 200) {
          this.$message.success('删除成功');
          this.issueList.splice(idx,1);
        }else {
          this.$message.error('删除成功');
          console.error(res.error.msg)
        }
      }).catch(err =>{
        this.$message.error('删除成功');
        console.error(err)
      })
    },
    pubIssue(id,idx) {
      let params = {
        id:id
      }
      pubIssue(params).then(res => {
        if(res.status == 200) {
          this.$message.success('发表成功');
          this.issueList[idx].Status = 1;
        }else {
          this.$message.error('发表失败');
          console.error(res.error.msg)
        }
      }).catch(err =>{
        this.$message.error('发表失败');
        console.error(err)
      })
    },
    addIssue() {
      addIssue().then(res => {
        if(res.status == 200) {
          this.$message.success('添加成功');
          this.getAllIssues();
        }else {
          this.$message.error('添加失败');
          console.error(res.error.msg)
        }
      }).catch(err =>{
        this.$message.error('添加失败');
        console.error(err)
      })
    }
  },
  created() {
    this.getAllIssues();
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
a {
  text-decoration: none;
  color: #409EFF;
}

.issue-add {
  margin: 10px 0;
  margin-right: 50px;
  height:50px
}

.issue-add button{
  float: right;
}
</style>
