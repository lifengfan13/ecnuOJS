
import { get, post,del,put,request} from '../http'
import QS from 'qs';

// for users
export const register = params => post('/auth/register', params);
export const login = params => post('/auth/login', params);
export const getUserInfo = params => get('/auth/userinfo', params);

//获取所有期刊
export const getAllissue = params => get('/allissue', params); 
//获取某一期刊下所有文章
export const getIssueArticle = params => get('/issuesubmission', params); 
//获取特定期刊
export const getIssue = params => get('/issue', params); 
//获取最新期刊
export const getLatestIssue = params => get('/latestissue', params); 
//获取特定文章
export const getArticle = params => get('/submission', params); 
//修改用户信息
export const changeUserInfo = params => put('/users', params);
//获取搜索文章
export const getSearchArticle = params => get('/searchsubmission', params); 