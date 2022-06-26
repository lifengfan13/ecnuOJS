<template>
  <div class="home-page">
        <section class="homepage_about">
			<a id="homepageAbout"></a>
			<h2>关于本期刊</h2>
            <h3 class="post-title">华东师范大学软件期刊</h3>
            <div id="content">
            <p>是一本刊登计算机软件各领域原创性研究成果的期刊，所刊登的论文均经过严格的同行专家评议。《华东师范大学软件期刊》主要面向全球华人计算机软件学者，致力于创办与世界计算机科学和软件技术发展同步的以中文为主的"中文国际软件学术期刊"，为全球华人同行提供学术交流平台，本刊不接受任何语种翻译稿。</p>
            <p><strong>历史</strong></p>
            <p>本期刊创立于2022年6月，由华东师范大学学生***、***与***主办</p>
            <p><strong>主要涉及方向</strong></p>
            <ul>
            <li>理论计算机科学</li>
            <li>算法设计与分析</li>
            <li>系统软件与软件工程</li>
            <li>模式识别与人工智能</li>
            <li>数据库技术</li>
            <li>计算机网络</li>
            <li>信息安全</li>
            <li>计算机图形学与计算机辅助设计</li>
            <li>多媒体技术</li>
            </ul>
            </div>
		</section>
        <section class="current_issue">
			<a id="homepageIssue"></a>
			<h2>
				最新一期
			</h2>
			<div class="current_issue_title">
				{{journal_name}}
			</div>
				<div class="obj_issue_toc">
		            <div class="heading">
						<router-link :to="{
							path:'/main/journal',
							query:{
								journal_identifier:journal_identifier,
							}
						}">
							<img class="journal_img_size" src="../../../build/journal.png">
						</router-link>
						<div class="published">
				            <span class="label">
					            出版于:
				            </span>
				            <span class="value">
					            {{journal_publish_date}}
				            </span>
			            </div>
			        </div>
		            <div class="sections">
			            <div class="section">
							<h3>
					            内含文章
				            </h3>
						    <ul class="cmp_article_list articles">
                                <li v-for="(article,index) in articleList" :key="index">
                                    <div class="obj_article_summary">
                                        <h4 class="title">
											<router-link :to="{
												path:'/main/article',
												query:{
													article_identifier:article.identifier,
												}
											}">
												{{article.title}}
											</router-link>
	                                    </h4>
   	                                    <div class="meta">
				                            <div class="authors">
			                                    {{article.author}}
		                                    </div>
                                        </div>
                                    </div>
                                </li>
							</ul>
				    	</div>
		        	</div><!-- .sections -->
            	</div>
			<router-link :to="{
					path:'/main/journal',
					query:{
						journal_identifier:journal_identifier,
					}
				}">
					展示完整期次
			</router-link>
		</section>
  </div>
</template>

<script>
import {getLatestIssue,getIssueArticle,getUserInfo}from '@/request/front/api';
export default {
    name: "HomePage",
    data() {
        return {
			journal_identifier:2,
			journal_name:"Vol. 2 No. 2 (2020): 华东师大软件期刊",
			journal_publish_date:"2021-11-29",
			articleList:[
			{
				identifier:123,
				title:"hello,hxc",
				author:"lff"
			},{
				identifier:123,
				title:"hello,hxc",
				author:"lff"
			},{
				identifier:123,
				title:"hello,hxc",
				author:"lff"
			}
			]
		}
    },
	methods:{
		getAuthor(i,authorid){
			if(authorid>0){
				getUserInfo({"id":authorid}).then(res => {
				if(res.status == 200) {
					this.articleList[i].author=res.data.Username
				}else {
					this.$message.error("获取作者失败")
				}
				}).catch(err =>{
				console.error(err)
				})
			}
		}
	},
    created(){
        getLatestIssue().then(res => {
            if(res.status == 200) {
                this.journal_identifier=res.data.ID
                this.journal_name="华东师范大学软件期刊第"+res.data.ID+"期"
                this.journal_publish_date=res.data.PublishDate
				getIssueArticle({"iid":this.journal_identifier}).then(res => {
					if(res.status == 200) {
						res.data.splice(3)
						this.articleList.splice(0,this.articleList.length)
						var i=0
						res.data.forEach(x => {
							this.articleList.push({
								"identifier":x.ID,
								"title":x.Name,
								"author":""
							})
							this.getAuthor(i,x.FirstAuthor)
							i++
						});
					}else {
						this.$message.error("获取期刊失败")
					}
				}).catch(err =>{
					console.error(err)
				})
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
    .current_issue .current_issue_title{margin:1.43rem 0;font-weight:700}
    .current_issue .read_more{display:inline-block;position:relative;padding-right:2.143rem;font-size:.93rem;font-weight:700;line-height:2.143rem;color:#006798;text-decoration:none;margin-bottom:1.43rem}
    .current_issue .read_more:hover,.current_issue .read_more:focus{color:#008acb}@media (min-width:768px){.pkp_page_index .current_issue .section:last-child{margin-bottom:0}}
    .obj_issue_toc .sections:not(:first-child){margin-top:4.286rem}
    .obj_article_summary .pages {position: absolute;top: 0;right: 0;line-height: 2.143rem;}
    .obj_article_summary .authors {padding-right: 5em;}
    .obj_article_summary .meta {position: relative;padding-top: 0.357rem;font-size: .93rem;line-height: 1.43rem;}
	.journal_img_size{height:30rem ;}
</style>
