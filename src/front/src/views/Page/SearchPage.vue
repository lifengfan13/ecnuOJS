<template>
  <div class="search-page">
    <div class="pkp_structure_main" role="main">
      <a id="pkp_content_main"></a>
      <div class="page page_search">
        <h1>
          搜索
        </h1>
	      <el-form class="cmp_form" @keyup.native.enter="searcharticle()">
				  <div class="search_input">
			      <label class="pkp_screen_reader" for="query">
				      搜索文章
			      </label>
            <el-input type="text" v-model="searchdata" class="query" placeholder="文章标题"></el-input>
            <input type="text" style="display:none" />
			    </div>

		      <br>

		      <div class="submit">
			      <el-button type="primary" @click="searcharticle()" @keyup.enter="searcharticle()">搜索</el-button>
		      </div>
	      </el-form>

	      <h2 class="pkp_screen_reader">Search Results</h2>
        <div>
          <template v-if="articleList.length!=0">
          <ul class="search_results">
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
          </template>
          <template v-else>
            <span role="status">
              <div class="cmp_notification notice">
                No Results
              </div>
            </span>	
          </template>
        </div>
    		
      </div><!-- .page -->
	  </div><!-- pkp_structure_main -->
  </div>
</template>

<script>
import {getUserInfo,getSearchArticle} from '@/request/front/api'
export default {
  name: 'SearchPage',
  data(){
    return {
      searchdata:"",
      articleList:[]
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
		},
    searcharticle(){
      this.articleList.splice(0,this.articleList.length)
      if(this.searchdata==""){
        this.$message.error("请输入文章标题")
      }else{
        getSearchArticle({'name':this.searchdata}).then(res => {
              if(res.status == 200) {
                if(res.data!=null){
                  
                  var i=0
                  res.data.forEach(x => {
                    if(x.Issue!=0){
                      this.articleList.push({
                        "identifier":x.ID,
                        "title":x.Name,
                        "author":"",
                      })
                      this.getAuthor(i,x.FirstAuthor)
                      i++
                    }
                  });
                }
                this.$message.success("搜索成功")
              }else {
                this.$message.error("搜索失败")
              }
            }).catch(err =>{
              console.error(err)
            })
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
*, *:before, *:after {box-sizing: border-box;}
.cmp_form fieldset {margin: 0 0 1.43rem;padding: 0;border: none;}
fieldset {display: block;margin-inline-start: 2px;margin-inline-end: 2px;padding-block-start: 0.35em;padding-inline-start: 0.75em;padding-inline-end: 0.75em;padding-block-end: 0.625em;min-inline-size: min-content;border-width: 2px;border-style: groove;border-color: threedface;border-image: initial;}
.cmp_form label {display: block;cursor: point}
legend {box-sizing: border-box;color: inherit;display: table;max-width: 100%;padding: 0;white-space: norm}
legend {display: block;padding-inline-start: 2px;padding-inline-end: 2px;border-width: initial;border-style: none;border-color: initial;border-image: initial;}
.page_search .date_range legend {padding: 0;font-size: 1rem;}
.page_search .search_advanced legend {padding: 0.714rem 1.43rem;margin: 0;font-weight: 400;}
.cmp_form input[type="text"], .cmp_form input[type="email"], .cmp_form input[type="password"], .cmp_form input[type="url"], .cmp_form input[type="tel"], .cmp_form select, .cmp_form textarea {padding: 0 0.5em;width: 100%;height: calc(2.143rem - 2px);background: #fff;border: 1px solid rgba(0,0,0,0.4);border-radius: 3px;font-size: .93rem;line-height: calc(2.143rem - 2px);}
.page_search .search_input .query{width:100%;max-width:100%;height:calc(2.857rem - 2px);font-size:1.285rem;line-height:calc(2.857rem - 2px)}
.page_search .search_advanced{border:1px solid #ddd;padding:0 1.43rem 1.43rem}
.page_search .search_advanced legend{padding:.714rem 1.43rem;margin:0;font-weight:400}
.page_search .search_advanced label{font-size:1rem;font-style:normal}
.page_search .date_range legend{padding:0;font-size:1rem}
.page_search .date_range label{clip:rect(1px, 1px, 1px, 1px);position:absolute !important;left:-2000px}
.page_search .date_range label:focus{background-color:#fff;border-radius:3px;box-shadow:0 0 2px 2px rgba(0,0,0,0.6);-webkit-box-shadow:0 0 2px 2px rgba(0,0,0,0.6);clip:auto !important;color:#000;display:block;font-size:1rem;height:auto;line-height:normal;padding:1rem;position:absolute;left:0.5rem;top:0.5rem;text-decoration:none;width:auto;z-index:100000}
.page_search .date_range select+label+select{margin-left:0.25em}
.page_search .date_range [name*="Year"]{width:6em}
.page_search .date_range [name*="Day"]{width:4em}
.page_search .date_range [name*="Month"]{width:10em}
.page_search .submit{text-align:right}
.page_search .submit button{display: inline-block;padding: 0 1em;background: #eee;border: 1px solid rgba(0,0,0,0.4);border-top-color: #bbb;border-radius: 3px;box-shadow: inset 0 -1em 1em rgb(0 0 0 / 10%);font-size: .93rem;line-height: calc(2.143rem - 2px);font-weight: 700;color: #006798;text-decoration: none;}
.page_search .submit button:hover,.page_search .submit button:focus{box-shadow:inset 0 1em 1em rgba(0,0,0,0.2)}
.page_search .submit button{position:relative;padding-right:3.2145rem;border-right:none;padding-right:1em;padding-left:1em;border-right:1px solid rgba(0,0,0,0.4);border-left:none}
.page_search .search_results{margin:2.857rem 0;padding:0;list-style:none}
.page_search .search_results .obj_article_summary{margin:1.43rem 0}
.page_search .cmp_pagination{margin-top:1.43rem;font-size:.93rem;line-height:1.43rem;color:rgba(0,0,0,0.54);text-align:right}
.page_search .cmp_pagination a{padding-left:0.5em;padding-right:0.5em}
.page_search .author{float:left;width:50%}

/*! normalize.css v7.0.0 | MIT License | github.com/necolas/normalize.css */
</style>
