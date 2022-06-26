<template>
  <div class="submission-page">
    <div class="pkp_structure_main" role="main">
			<a id="pkp_content_main"></a>

      <div class="page page_submissions">
        <h1>
          提交
        </h1>
        <!--未登录-->
        <div class="cmp_notification" v-if="!isLogin">
          <router-link to="login">登录</router-link> 或 <router-link to="register">注册</router-link> 来进行提交.
        </div>
        <!--登录后-->
        <div class="cmp_notification" v-if="isLogin">
					<router-link to="/admin/submission">创建新的提交</router-link> 或 <router-link to="/admin/historySubmission">查看历史提交</router-link>.
				</div>

        <div class="submission_checklist">
          <h2> 
            提交所需
          </h2>
          作为提交过程的一部分，作者需要检查其提交的内容是否符合以下所有项目，并且提交的内容可能会退还给不遵守这些指南的作者。
          <ul>
            <li>
              <span class="fa fa-check" aria-hidden="true"></span>
              提交的材料之前没有发表过，也没有提交给其他杂志考虑（或者在给编辑的评论中提供了解释）。
            </li>
            <li>
              <span class="fa fa-check" aria-hidden="true"></span>
              提交文件为PDF文档文件格式。
            </li>
            <li>
              <span class="fa fa-check" aria-hidden="true"></span>
              在可用的情况下，已提供参考的 URL。
            </li>
            <li>
              <span class="fa fa-check" aria-hidden="true"></span>
              文本为单间距；使用12号字体；使用斜体，而不是下划线（URL地址除外）；所有插图、图表和表格都放在文本的适当位置，而不是放在末尾。
            </li>
            <li>
              <span class="fa fa-check" aria-hidden="true"></span>
              文本符合作者指南中概述的文体和书目要求。
            </li>
          </ul>
        </div>  
        <div class="section_policy">
          <h2>文章</h2>
          默认策略
        </div>  
        <div class="privacy_statement" id="privacyStatement">
          <h2>
            隐私声明
          </h2>
          <p>本期刊网站中输入的姓名和电子邮件地址将仅用于本期刊的指定目的，不会用于任何其他目的或提供给任何其他方。</p>
        </div>     
      </div><!-- .page -->
	  </div><!-- pkp_structure_main -->
  </div>
</template>

<script>
import { mapGetters, mapState } from 'vuex'
export default {

  name: 'SubmissionPage',
  data () {
    return{
      
    }
  },
  computed: {
    ...mapState ({
      isLogin: state => state.isLogin,
    }),
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
*, *:before, *:after {
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
}
.cmp_notification {
    display: block;
    width: 93%;
    padding: 1.43rem;
    margin-bottom: 2.857rem;
    background: #ddd;
    border-left: 0.357rem solid #006798;
    font-size: 1rem;
    line-height: 1.43rem;
}

.page_submissions .submission_checklist ul {
    border: 1px solid #ddd;
    border-bottom: none;
}

.page_submissions .submission_checklist ul {
    margin: 1.43rem 0 0;
    padding: 0;
    list-style: none;
    font-size: .93rem;
}

.page_submissions .submission_checklist li {
    padding: 1.43rem 1.43rem 1.43rem 2.857rem;
    border-bottom: 1px solid #ddd;
}
</style>
