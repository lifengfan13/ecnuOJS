<template>
  <div class="pkp_structure_main">
    <div class="archives-page">
      <div class="pkp_structure_main" role="main">
        <a id="pkp_content_main"></a>
        <div class="page page_issue_archive">
          <h1>
            Archives
          </h1>
          <ul class="issues_archive">
            <li v-for="(j,index) in journalList" :key="index">
              <div class="obj_issue_summary">
                <router-link class="cover" :to="{
                  path:'/main/journal',
                  query:{
                    journal_identifier:j.identifier,
                  }
                }">
                  <img src="../../../build/journal.png">
                </router-link>
                <h2>
                  <router-link class="title" :to="{
                    path:'/main/journal',
                    query:{
                      journal_identifier:j.identifier,
                    }
                  }">
                    {{j.title}}
                  </router-link>
                </h2>
                <div class="description">
                </div>
              </div><!-- .obj_issue_summary -->
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getAllissue } from '@/request/front/api';
export default {
  name: 'ArchivesPage',
  data(){
    return {
      journalList:[
        {
          identifier:1,
          title:"华东师范大学软件期刊第一期（2022）",
        },
        {
          identifier:2,
          title:"华东师范大学软件期刊第二期（2022）",
        },
        {
          identifier:3,
          title:"华东师范大学软件期刊第三期（2022）",
        }
      ]
    }
  },
  methods:{
    setjournals(data){
      if(data!=[]){
        this.journalList.splice(0,this.journalList.length)
        data.forEach(x => {
          if(x.PublishDate!="")
            this.journalList.push({"identifier":x.ID,"title":"华东师范大学软件期刊第"+x.ID+"期"})
        })
      }
    }
  },
  created(){
      getAllissue().then(res => {
              if(res.status == 200) {
                this.setjournals(res.data) 
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
.pkp_structure_main{padding:.714rem}
.pkp_structure_main h1{font-family:"Noto Sans",-apple-system,BlinkMacSystemFont,"Segoe UI","Roboto","Oxygen-Sans","Ubuntu","Cantarell","Helvetica Neue",sans-serif;font-size:1.714rem;line-height:2.143rem;font-weight:700}
.pkp_structure_main h1,.pkp_structure_main h2,.pkp_structure_main h3,.pkp_structure_main h4{margin:2.857rem 0 1.43rem}
.obj_issue_summary .cover {
    float: left;
    width: 25%;
    height: auto;
    margin-right: 1.43rem;
}
.obj_issue_summary .cover img {
    display: block;
    width: auto;
    max-height: 250px;
}
img {
    max-width: 100%;
    width: auto;
    height: auto;
}
.obj_issue_summary .cover {
    display: block;
    margin-bottom: 1.43rem;
}
.obj_issue_summary h2 {
    margin: 0;
    font-size: 1rem;
    line-height: 1.43rem;
    font-weight: 400;
}
h2 {
    display: block;
    margin-block-start: 0.83em;
    margin-block-end: 0.83em;
    margin-inline-start: 0px;
    margin-inline-end: 0px;
}
.obj_issue_summary .title {
    font-family: "Noto Sans",-apple-system,BlinkMacSystemFont,"Segoe UI","Roboto","Oxygen-Sans","Ubuntu","Cantarell","Helvetica Neue",sans-serif;
    font-weight: 700;
    text-decoration: none;
}
.obj_issue_summary h2 {
    margin: 0;
    font-size: 1rem;
    line-height: 1.43rem;
    font-weight: 400;
}
.obj_issue_summary:before, .obj_issue_summary:after {
    content: " ";
    display: table;
}
li {
    text-align: -webkit-match-parent;
}
.obj_issue_summary:before, .obj_issue_summary:after {
    content: " ";
    display: table;
}
.page_issue_archive .issues_archive{margin-left:-0.714rem;margin-right:-0.714rem;border-top:1px solid #ddd}
.page_issue_archive .issues_archive>li{padding:2.143rem .714rem;border-bottom:1px solid #ddd}
.page_issue_archive .issues_archive>li {padding-left: 2.143rem;padding-right: 2.143rem;    height: 20em;}
.obj_issue_summary .cover {float: left;width: 25%;height: auto;margin-right: 1.43rem;}
</style>
