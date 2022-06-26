<template>
  <div class="current-page">
      最新一期
      跳转至journalPage,获取最新一期的期号
  </div>
</template>

<script>
import {getLatestIssue} from '@/request/front/api'
export default {
  name: 'CurrentPage',
  created(){
    getLatestIssue().then(res => {
      if(res.status == 200) {
        this.$router.push({path:'/main/journal',query: {journal_identifier:res.data.ID}})
      }else {
        this.$message.error("获取期刊失败")
      }
    }).catch(err =>{
      console.error(err)
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
