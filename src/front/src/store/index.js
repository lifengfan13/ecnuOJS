import Vue from 'vue';
import Vuex from 'vuex';

import storage from '@/model/storage'

Vue.use(Vuex);
//新建store实例
const store = new Vuex.Store({
  state: {
    // 存储token
    isLogin:false,
    token:"",
    userId:"",
    userInfo: {},
  },

  getters: {
    getToken(state){
      return state.token || storage.get("token") || "";
    },
    getUserId(state){
      return state.userId || storage.get("userId") || "";
    }
  },

  mutations: {
    // 修改token，并将token存入localStorage
    setLogin(state,isLogin) {
      storage.set('isLogin', isLogin);
      state.isLogin = isLogin;
    },
    setUserId(state,userId) {
      storage.set('userId', userId);
      state.userId = userId;
    },
    delUserId(state) {
      state.userId = '';
      storage.remove("userId");
    },
    setUserInfo(state,userInfo) {
      state.userInfo = userInfo;
    },
    delUserInfo(state) {
      state.userInfo = {};
      storage.remove("UserInfo");
    },
    setToken(state,token) {
      state.token = token;
      storage.set('token', token);
      console.log('store、localstorage保存token成功！');
    },
    delToken(state) {
      state.token = "";
      storage.remove("token");
    }
  },

 actions: {
   // removeToken: (context) => {
     // context.commit('setToken')
   // }
 },
})
//导出store
export default store;
