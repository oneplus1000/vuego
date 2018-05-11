 

import Vue from 'vue';
import  SuiVue from 'semantic-ui-vue';
import App from './App.vue';
import router from './router';

const s = require('semantic-ui-css/semantic.min.css');
 

Vue.config.productionTip = false;
Vue.use(SuiVue);

new Vue({
  router,
  render: (h) => h(App),
}).$mount('#app');

