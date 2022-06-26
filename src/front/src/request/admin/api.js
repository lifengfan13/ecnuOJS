
import { get, post,del,put,request} from '../http'
import QS from 'qs';

//获取所有用户信息
/**
 * get方法，对应get请求
 * @param {String} idx [分页索引]
 */
export const getAllUsers = params => get('/users', params);
//获取用户信息
export const getUserInfo = params => get('/auth/userinfo', params);
// submission operation
export const submitArticle = params => post('/submission', params);
//修改投稿
export const updateSubmission = params => put('/submission', params);
//删除投稿
export const delArticle = params => del('/submission', params);

//根据投稿id 获取投稿
export const getSubmission = params => get('/submission', params);

//根据用户id获取投稿，以分页的形式
/**
 * get方法，对应get请求
 * @param {String} id [用户id]
 * @param {String} idx [分页索引]
 */
export const getPageAllSubmissionByUserId = params => get('/submission/user/paging', params);

//获取所有投稿 分页获取
export const getAllSubmission = params => get('/submission/admin/paging', params);
//下载投稿对应论文
export const downloadPaper = params => get('/submission/paper', params);

//接受论文投稿
export const acceptSubmisson = params => get('/approvesubmission', params);
//拒绝论文投稿
export const refusesubmission = params => get('/refusesubmission', params);
//发表投稿
export const issuesubmission = params => request({
    method: 'post',
    url: '/issuesubmission',
    params: params
});


//根据评审人id获取所有审稿
/**
 * get方法，对应get请求
 * @param {String} id [评审人id]
 * @param {String} idx [分页索引]
 */
export const getAllReviewsForUser = params => get('/reviewer-papers-review', params);

//根据投稿id 获取所有评审论文
export const getReviewsBySubmission = params => get('/contributor-papers-review', params);

//新建评审
export const newReview = params => request({
    method: 'post',
    url: '/reviewersbm',
    params: params
});
export const delReview = params => del('/reviewersbm', params);
//更新投稿
export const updateReviewcmt = params => post('/reviewcmt', params);

//获取评审人
export const getReviewers = () => get('/reviewer');

//获取所有issue
export const getAllIssues = () => get('/allissue');
//删除issue
export const delIssue = params => del('/issue',params)
//添加issue
export const addIssue = params => request({
    method: 'post',
    url: '/issue'
});

//发表issue
export const pubIssue = params => request({
    method: 'put',
    url: '/issue',
    params: params
});
//修改用户权限
export const changeUserMode = params => request({
    method: 'put',
    url: '/mode',
    params: params
});
//从期刊中删除文章
export const delArticlefromIssue = params => del('/issuesubmission',params)