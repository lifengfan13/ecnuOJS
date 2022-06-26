<template>
  <div class="user-board">
      <title-header title="用户列表"></title-header>
      <div class="content">
        <el-table :data="userList" highlight-current-row style="width: 100%">
          <el-table-column type="index"></el-table-column>

          <el-table-column label="用户名" align="center">
            <template slot-scope="scope">
              <i class="el-icon-user"></i>
              <span>{{ scope.row.Username }}</span>
            </template>
          </el-table-column>

          <el-table-column label="邮箱" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.Email }}</span>
            </template>
          </el-table-column>

          <el-table-column label="组织" align="center">
            <template slot-scope="scope">
              <span>{{ scope.row.Organization}}</span>
            </template>
          </el-table-column>

          <el-table-column label="角色" align="center">
            <template slot-scope="scope">
              <span>{{ authList[scope.row.Role - 1]}}</span>
            </template>
          </el-table-column>

          <el-table-column fixed="right" label="操作" width="100">
            <template slot-scope="scope">
              <el-button type="text" size="small" @click="showDialog(scope.row.ID,scope.$index)">编辑</el-button>
            </template>
          </el-table-column>
        </el-table>
         <el-dialog
            title="分配评审人"
            :visible.sync="dialogVisible"
            width="30%">
            <div class="dialog-main">
              <div class="sole-assign">
                <h6>角色分配</h6>
                <el-select v-model="ruleMode" placeholder="请选择">
                  <el-option
                    v-for="item in roleList"
                    :key="item.value"
                    :label="item.name"
                    :value="item.mode">
                  </el-option>
                </el-select>
              </div>
            </div>
          
            <span slot="footer" class="dialog-footer">
              <el-button @click="dialogVisible = false">取 消</el-button>
              <el-button type="primary"  @click="handleEdit()">确 定</el-button>
            </span>
          </el-dialog>

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
  </div>
</template>

<script>
import TitleHeader from '@/components/common/TitleHeader.vue'
import { getAllUsers,changeUserMode } from '@/request/admin/api';

export default {
  name: 'UserBoard',
  components: {
    TitleHeader
  },
  data () {
    return {
      showChange:false,
      cuser:{},
      pageIdx:1,
      userList:[],
      authList:['作者','审稿人','管理员'],
      dialogVisible:false,
      current: {
        id:0,
        idx:0
      },
      roleList: [
        {name:"作者",mode:1},
        {name:"审稿人",mode:2},
        {name:"管理员",mode:3}
      ],
      ruleMode:undefined,
    }
  },
  methods: {
    assgindata(id,mode){
      changeUserMode({id:id,mode:mode}).then(res => {
        if(res.status == 200) {
          this.showChange=false
          this.$message.success('修改成功');
          this.getAllUsers(this.pageIdx);
        }else {
          console.error(res.error.msg)
        }
      }).catch(err => {
        console.error(err)
      })
    },
    getAllUsers(pageIdx) {
      let params = {
        idx:pageIdx,
      }
      getAllUsers(params).then(res => {
        if(res.status == 200) {
          this.userList = res.data;
        }else {
          console.error(res.error.msg)
        }
      }).catch(err => {
        console.error(err)
      })
    },
    changePage(pageIdx) {
     this.getAllUsers(pageIdx);
    },
    showDialog(id,idx) {
      this.dialogVisible = true;
      this.current.id = id;
      this.current.idx = idx;
    },
    handleEdit() {
      if(this.ruleMode == undefined) {
        this.$message.warning('请先选择角色再提交');
        return;
      }
      let params = {
        id:this.current.id,
        mode:this.ruleMode,
      }
      changeUserMode(params).then(res => {
        if(res.status == 200) {
          this.dialogVisible = false;
          this.$message.success('修改成功!');
          this.userList[this.current.idx].Role = this.ruleMode;
          this.ruleMode = undefined;
        }else {
          this.$message.error('修改失败!');
          console.error(res.error.msg)
        }
      }).catch(err => {
        this.$message.error('修改失败!');
        console.error(err)
      })
    }
  },
  created() {
    this.getAllUsers(this.pageIdx);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>


.pagination-box {
  margin-top: 20px;
  height: 50px;
}

.pagination-page{
  float: right;
}
</style>
