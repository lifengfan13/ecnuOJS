import Vue from 'vue'

// 1.完整引入element ui的组件
import ElementUI from 'element-ui'
// 2.导入element ui组件的样式
import 'element-ui/lib/theme-chalk/index.css'

// 3.把ElementUI注册为vue组件，注册之后即可在每个组件中直接使用element ui的组件
Vue.use(ElementUI)
