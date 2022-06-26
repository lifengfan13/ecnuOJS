<template>
  <div class="title-header">
      <h2>{{title}}</h2>
      <el-divider></el-divider>
  </div>
</template>

<script>

export default {
  name: 'TitleHeader',
  props:{
    title:{
        type:String,//在props接受父组件传递数据
        default:'标题'
    }
  },
  
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
