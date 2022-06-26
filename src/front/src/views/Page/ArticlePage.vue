<template>
  <div>
    <div class="pkp_structure_main" role="main">
      <div class="page page_article">
		  	<div class="obj_article_details">
          <h1 class="page_title">
		        {{title}}
	        </h1>
          <br>
	        <div class="row">
		        <div class="main_entry">
							<section class="item authors">
					      <h2 class="pkp_screen_reader">作者</h2>
					      <ul class="authors">
									<li v-for="(item,index) in authors" :key="index">
                    <span class="name">
                      {{item}}
                    </span>
									</li>
								</ul>
				      </section>
							<!-- <section class="item keywords">
				        <h2 class="label">
									关键词:
				        </h2>
                
				        <span class="value" v-for="(item,index) in keywords" :key="index">
									{{item}}
                  <span v-if="index!=keywords.length-1">,</span>									
                </span>
			        </section> -->
			        <section class="item abstract">
					      <h2 class="label">摘要</h2>
                <p class="p_abstract">{{abstract}}</p>
					    </section>	
		        </div>
            
            <div class="entry_details">
              <div class="item cover_image">
                <div class="sub_item">
                  <router-link :to="{
                    path:'/main/journal',
                    query:{
                      journal_identifier:upper_journal,
                    }
                  }">
                    <img src="../../../build/journal.png">
                  </router-link>
                </div>
              </div>
              <div class="item galleys">
                <h2 class="pkp_screen_reader">
                  下载文章
                </h2>
                <ul class="value galleys_links">
                  <li>
                    <!-- download_address -->
                    <a class="obj_galley_link pdf" :href="'http://47.103.81.142:8080/api/v1/submission/paper?id='+identifier">
                      PDF
                    </a>
                  </li>
                </ul>
              </div>
              <div class="item published">
                <section class="sub_item">
                  <h2 class="label">
                    出版时间
                  </h2>
                  <div class="value">
                    <span>{{upper_journal_date}}</span>
                  </div>
                </section>
              </div>
              <!-- <div class="item citation">
                <section class="sub_item citation_display">
                  <h2 class="label">
                    如何引用
                  </h2>
                  <div class="value">
                    <div id="citationOutput" role="region" aria-live="polite">
                      <div class="csl-bib-body">
                        <div class="csl-entry">{{cite_way}} </div>
                      </div>
                    </div>
                    
                  </div>
                </section>
              </div> -->

            </div>
	        </div>
        </div>


      </div><!-- .page -->
	  </div>
    <!-- <Article></Article> -->
  </div>
</template>

<script>
import {getArticle,getUserInfo,getIssue} from '@/request/front/api'
export default {
  name: 'ArticlePage',
  data(){
    return{
      identifier:0,
      title:"",
      authors:[
        "Charlie Taylor",
        "Kenya Juarez"
      ],
      keywords:[
        "provide", 
        "peaceful"
      ],
      abstract:"",
      upper_journal:1,
      upper_journal_date:"",
    }
  },
  methods:{
    getAuthor(authorid){
      if(authorid>0){
        getUserInfo({"id":authorid}).then(res => {
          if(res.status == 200) {
            this.authors.push(res.data.Username)
          }else {
            this.$message.error("获取作者失败")
          }
        }).catch(err =>{
          console.error(err)
        })
      }
    },
    getUpperjournal(journalid){
      if(journalid>0){
        getIssue({"id":journalid}).then(res => {
          if(res.status == 200) {
            this.upper_journal_date=res.data.PublishDate
          }else {
            this.$message.error("获取期刊失败")
          }
        }).catch(err =>{
          console.error(err)
        })
      }
    },
  },
  created(){
    this.identifier=this.$route.query.article_identifier
    // 需要根据identifier获取信息
    getArticle({"id":this.identifier}).then(res => {
      if(res.status == 200) {
        this.title=res.data.Name
        this.authors.splice(0,this.authors.length)
        this.getAuthor(res.data.FirstAuthor)
        this.getAuthor(res.data.SecondAuthor)
        this.getAuthor(res.data.ThirdAuthor)
        this.getAuthor(res.data.ForthAuthor)
        //this.keywords
        this.abstract=res.data.Abstract
        this.upper_journal=res.data.Issue
        this.getUpperjournal(res.data.Issue)
        //this.cite_way
      }else {
        this.$message.error("获取文章失败")
      }
    }).catch(err =>{
      console.error(err)
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.pkp_structure_main{padding:.714rem}
.obj_article_details>.page_title{margin:0}
.obj_article_details>.subtitle{margin:0;font-size:1rem;line-height:2.143rem;font-weight:400}
.obj_article_details .row{margin-top:2.143rem}
.obj_article_details .item{padding-top:1.43rem;padding-bottom:1.43rem}
.obj_article_details .item>*:first-child{margin-top:0}
.obj_article_details .item>*:last-child{margin-bottom:0}
.obj_article_details .sub_item{margin-bottom:1.43rem}
.obj_article_details .sub_item:last-child{margin-bottom:0}
.obj_article_details .main_entry .item .label{margin:0 0 1.43rem;font-family:"Noto Sans",-apple-system,BlinkMacSystemFont,"Segoe UI","Roboto","Oxygen-Sans","Ubuntu","Cantarell","Helvetica Neue",sans-serif;font-size:1.143rem;font-weight:700}
.obj_article_details .main_entry .item.doi .label,.obj_article_details .main_entry .item.keywords .label{display:inline;font-size:1rem}
.obj_article_details .main_entry .sub_item .label{font-size:1rem}
.obj_article_details .authors li{margin-bottom:.714rem}
.obj_article_details .authors .name{font-weight:bold;display:block}
.obj_article_details .authors .orcid{display:block;font-size:.75rem;line-height:1.43rem}
.obj_article_details .authors .orcid a{vertical-align:middle}
.obj_article_details .authors .orcid_icon{width:1.43rem;height:1.43rem}
.obj_article_details .authors .affiliation{font-size:.93rem;color:rgba(0,0,0,0.54)}
.obj_article_details .author_bios .sub_item .label{margin-bottom:0}
.obj_article_details .author_bios .sub_item .value>p:first-child{margin-top:0}
.obj_article_details .item.doi,.obj_article_details .item.keywords{padding-top:0}
.obj_article_details .galleys_links li{display:inline-block}
.obj_article_details .supplementary_galleys_links{margin-top:.714rem}
.obj_article_details .copyright{font-size:.93rem;line-height:1.43rem}
.obj_article_details .copyright a[rel="license"]+p{margin-top:0}
.obj_article_details .entry_details{margin-left:-1.43rem;margin-right:-1.43rem;border-top:1px solid #ddd}
.obj_article_details .entry_details .item{padding:1.43rem;border-bottom:1px solid #ddd;word-wrap:break-word}
.obj_article_details .entry_details .item:last-child{border-bottom:none}
.obj_article_details .entry_details .item .label{margin:0;font-family:"Noto Sans",-apple-system,BlinkMacSystemFont,"Segoe UI","Roboto","Oxygen-Sans","Ubuntu","Cantarell","Helvetica Neue",sans-serif;font-size:.93rem;font-weight:400;color:rgba(0,0,0,0.54)}
.obj_article_details .categories{margin:0;padding:0;list-style:none}
.obj_article_details .versions ul{margin:0;padding:0;list-style:none}
.obj_article_details .citation_display .value{font-size:.75rem}
.obj_article_details .citation_display .csl-left-margin{display:none}
.obj_article_details .citation_display [aria-hidden="true"]{display:none}
.obj_article_details .citation_display .citation_formats{margin-top:1em;border:1px solid rgba(0,0,0,0.4);border-radius:3px}
.obj_article_details .citation_display .citation_formats_button{position:relative;background:transparent;border:none;border-bottom-left-radius:0;border-bottom-right-radius:0;box-shadow:none;padding:0 1em;width:100%;font-family:"Noto Sans",-apple-system,BlinkMacSystemFont,"Segoe UI","Roboto","Oxygen-Sans","Ubuntu","Cantarell","Helvetica Neue",sans-serif;font-weight:400;color:rgba(0,0,0,0.54);text-align:left}
.obj_article_details .citation_display .citation_formats_button:after{display:inline-block;font:normal normal normal 14px/1 FontAwesome;font-size:inherit;text-rendering:auto;-webkit-font-smoothing:antialiased;-moz-osx-font-smoothing:grayscale;transform:translate(0, 0);content:"\f0d7";position:absolute;top:50%;right:1em;transform:translateY(-50%)}
.obj_article_details .citation_display .citation_formats_button[aria-expanded="true"]:after{content:"\f0d8"}
.obj_article_details .citation_display .citation_formats_button:focus{background:#ddd;outline:0}
.obj_article_details .citation_display .citation_formats_styles{margin:0;padding:0;list-style:none}
.obj_article_details .citation_display .citation_formats_styles a{display:block;padding:.5em 1em;border-bottom:1px solid #ddd;text-decoration:none}
.obj_article_details .citation_display .citation_formats_styles a:focus{background:#ddd;outline:0}
.obj_article_details .citation_display .citation_formats_styles li:last-child a{border-bottom:none}
.obj_article_details .citation_display .citation_formats_list .label{padding:1em 1em .25em 1em}
.obj_article_details .citation_display .citation_formats_styles+.label{border-top:1px solid #ddd}
.obj_article_details .citation_display span{margin-right:0.5em}
@media (min-width:480px){.obj_article_details .entry_details{margin-left:-2.143rem;margin-right:-2.143rem}}
@media (min-width:768px){.obj_article_details .row{margin-left:-1.43rem;margin-right:-1.43rem;border-top:1px solid #ddd;border-bottom:1px solid #ddd}.obj_article_details .main_entry{float:left;width:428px;border-right:1px solid #ddd}.obj_article_details .item{padding:1.43rem}.obj_article_details .item .label{margin:0 0 1.43rem;font-family:"Noto Sans",-apple-system,BlinkMacSystemFont,"Segoe UI","Roboto","Oxygen-Sans","Ubuntu","Cantarell","Helvetica Neue",sans-serif;font-size:1.143rem;font-weight:700}.obj_article_details .item.doi .label,.obj_article_details .item.keywords .label{display:inline;font-size:1rem}.obj_article_details .entry_details{float:left;width:300px;margin:0 0 0 -1px;border-top:none;border-left:1px solid #ddd}.obj_article_details .entry_details .item{margin-right:-1px;border-bottom:1px solid #ddd}.obj_article_details .entry_details .item:last-child{border-bottom:none}}@media (min-width:992px){.obj_article_details .row{margin-left:-2.143rem;margin-right:-2.143rem}.obj_article_details .main_entry{width:352px}.obj_article_details .item{padding:2.143rem}}
@media (min-width:1200px){.obj_article_details .main_entry{width:560px}}
.obj_galley_link {display: inline-block;padding: 0 1em;background: #fff;border: 1px solid #006798;border-radius: 3px;font-size: .93rem;line-height: calc(2.143rem - 2px);color: #006798;text-decoration: none;}
.p_abstract{
	word-wrap: break-word;
	word-break: break-word;
	}
</style>
