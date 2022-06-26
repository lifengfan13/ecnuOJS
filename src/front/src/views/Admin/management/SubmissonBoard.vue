<template>
  <div class="submisson-board">
      <title-header title="投稿管理"></title-header>
      <div class="content">
        <el-table :data="tableData" highlight-current-row style="width: 100%">
          <el-table-column type="index"></el-table-column>

          <el-table-column label="投稿日期" align="center">
            <template slot-scope="scope">
              <i class="el-icon-time"></i>
              <span>{{ scope.row.CreatedAt.split('T')[0] }}</span>
            </template>
          </el-table-column>

          <el-table-column label="标题" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.Name }}</span>
            </template>
          </el-table-column>

          <!-- <el-table-column label="作者" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.FirstAuthor }}</span>
            </template>
          </el-table-column> -->

          <el-table-column prop="Status" label="状态" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.Issue == 0 ? statusList[scope.row.Status]:"已发表"}}</span>
            </template>
          </el-table-column>

          <el-table-column fixed="right" label="操作" width="100">
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="handleEdit(scope.row.ID)">查看</el-button>
            </template>
          </el-table-column>
        </el-table>
    </div>

    <div class="pagination-box">
      <div class="pagination-page">
        <el-pagination
          background
          :page-size="10"
          :pager-count="11"
          :current-page="pageIdx"
          layout="prev, pager, next"
          :total="1000"
          @current-change="changePage">
        </el-pagination>
      </div>
    </div>
  </div>
</template>

<script>
import TitleHeader from '@/components/common/TitleHeader.vue'

import { getAllSubmission } from '@/request/admin/api';

export default {
  name: 'SubmissonBoard',
  components: {
    TitleHeader
  },
  data() {
    return {
      pageIdx:1,
      tableData: [{
            "ID": 1,
            "CreatedAt": "2022-06-12T00:00:00+08:00",
            "UpdatedAt": "2022-06-12T00:00:00+08:00",
            "DeletedAt": null,
            "Name": "用P4对数据平面进行编程",
            "FirstAuthor": 3,
            "CorrespondingAuthor": 7,
            "SecondAuthor": -1,
            "ThirdAuthor": 0,
            "ForthAuthor": 0,
            "Topic": 1,
            "Status": 0,
            "FilePath": "/tmp/paper/fcd0a7080816a21ec9e50005f6f38f47-helloword.jepg",
            "IsDel": 0,
            "Abstract": "123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789123456789",
            "Organization": "华东师范大学",
            "Address": "华东师范大学中北校区"
        }
      ],
      statusList: ["审核中","拒收","审稿通过"]
    }
  },
  methods: {
    handleEdit(id) {
      this.$router.push({ path: 'submissondetail', query: { id: id }})
    },
    getAllSubmission(pageIdx) {
      getAllSubmission({
        idx:pageIdx
      }).then(res => {
        if(res.status == 200) {
          this.tableData = res.data
        }else {
          console.error(res.error.msg)
        }
      }).catch(err =>{
        console.error(err)
      })
    },
    changePage(pageIdx) {
     this.getAllSubmission(pageIdx);
    }
  },
  created() {
   this.getAllSubmission(this.pageIdx);
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.state {
  color: #d00a0a;
}

.pagination-box {
  margin-top: 20px;
  height: 50px;
}

.pagination-page{
  float: right;
}
</style>
