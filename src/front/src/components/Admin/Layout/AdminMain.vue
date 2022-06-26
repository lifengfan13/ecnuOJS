<template>
  <div class="admin-main">
    <div class="main">
      <!-- <keep-alive exclude="SubmissonDetail"> -->
        <router-view></router-view>
      <!-- </keep-alive> -->
    </div>
  </div>
</template>

<script>


export default {
  name: 'AdminMain'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.admin-main {
  /* background-color: #eaedee; */
  margin: 0 300px;
}

.main {
  padding: 20px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04);
  font-size: 16px;
}

</style>
