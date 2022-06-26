<template>
  <div class="sidebar-page">
      <div class="block_make_submission">
        <h2 >
          创建新的投稿
        </h2>
        <div class="content">
          <router-link to="/main/submission">创建新的投稿</router-link>
        </div>
      </div>
      <div class="block_information">
        <h2 class="title">提示</h2>
        <div class="content">
          <ul>
              <li>
                <router-link to="">关于读者</router-link>
              </li>
              <li>
                <router-link to="">关于作者</router-link>
              </li>
              <li>
                <router-link to="">关于期刊</router-link>
              </li>
          </ul>
        </div>
      </div>
  </div>
</template>

<script>

export default {
  name: 'SidebarPage'
}
</script>

<style scoped>
.block_make_submission,.block_information {
  padding: 30px;
}
.block_make_submission a {
    display: inline-block;
    padding: 0 1em;
    background: #eee;
    border: 1px solid rgba(0,0,0,0.4);
    border-top-color: #bbb;
    border-radius: 3px;
    box-shadow: inset 0 -1em 1em rgb(0 0 0 / 10%);
    font-size: .93rem;
    line-height: calc(2.143rem - 2px);
    font-weight: 700;
    color: #006798;
    text-decoration: none;
}

.block_information .title {
    display: block;
    margin-bottom: 0.714rem;
    margin-top: 0;
    font-family: "Noto Sans",-apple-system,BlinkMacSystemFont,"Segoe UI","Roboto","Oxygen-Sans","Ubuntu","Cantarell","Helvetica Neue",sans-serif;
    font-size: 1.143rem;
    font-weight: 700;
    line-height: 1.43rem;
    color: rgba(0,0,0,0.54);
}

.block_information .content ul {
    margin: 0;
    padding: 0;
    list-style: none;
}

.block_information .content ul li {
    line-height: 25px;
    padding: 3px 0;
}
.block_information .content ul li a {
    font-size: .93rem;
    color: #006798;
    background-color: transparent;
    text-decoration: none;
}

</style>
