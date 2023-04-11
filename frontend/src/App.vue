<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 19:39:23
 * @FilePath: \http-file\frontend\src\App.vue
-->
<template>
  <a-config-provider :locale="locale">
    <global />
    <router-view />
  </a-config-provider>
</template>

<script>
import zhCN from "ant-design-vue/lib/locale-provider/zh_CN";
import Global from "./components/Global.vue";
export default {
  name: "App",
  components: { Global },
  data() {
    return {
      locale: zhCN,
    };
  },
};
</script>