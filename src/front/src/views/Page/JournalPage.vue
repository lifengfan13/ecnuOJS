<template>
  	<div>
    <!-- 显示期刊和期刊中的文章 -->
    	<section class="current_issue">
			<a id="homepageIssue"></a>
			<div class="current_issue_title">
				{{journal_name}}
			</div>
				<div class="obj_issue_toc">
		      		<div class="heading">
						<img  class="journal_img_size" src="../../../build/journal.png">
						<div class="published">
				      		<span class="label">
					      		Published:
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
		</section>
  	</div>
</template>

<script>
import { getIssueArticle,getIssue,getUserInfo } from '@/request/front/api';
export default {
  name: 'JournalPage',
  data(){
		return {
			journal_identifier:0,
			journal_name:"",
			journal_publish_date:"",
			articleList:[
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
		this.journal_identifier=this.$route.query.journal_identifier
		getIssue({"id":this.journal_identifier}).then(res => {
            if(res.status == 200) {
				this.journal_name="华东师范大学软件期刊第"+this.journal_identifier+"期"
				this.journal_publish_date=res.data.PublishDate
            }else {
                this.$message.error("获取期刊失败")
            }
        }).catch(err =>{
            console.error(err)
        })
		//根据journalid获取全部简略article信息
		getIssueArticle({"iid":this.journal_identifier}).then(res => {
            if(res.status == 200) {
				this.articleList.splice(0,this.articleList.length)
				var i=0
				res.data.forEach(x => {
					this.articleList.push({
						"identifier":x.ID,
						"title":x.Name,
						"author":"",
					})
					this.getAuthor(i,x.FirstAuthor)
					i++
				});
				
				// this.journal_name="华东师范大学软件期刊第"+this.journal_identifier+"期"
				// this.journal_publish_date=res.data.PublishDate
            }else {
                this.$message.error("获取期刊失败")
            }
        }).catch(err =>{
			console.log("!!!!!!")
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