<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 22:39:47
 * @FilePath: \http-file\frontend\src\components\Footer.vue
-->
<template>
  <div>
    <div class="footer">
      Powered By
      <a target="_blank" href="https://github.com/Xhofe/alist">AList</a>
      <a-divider type="vertical" />
      <span v-if="siteStore.siteInfo.footer_text">
        ©2023
        <a target="_blank" :href="siteStore.siteInfo.footer_url">
          {{ siteStore.siteInfo.footer_text }}</a
        >
      </span>
    </div>
  </div>
</template>
<script>
export default {
  name: "Footer",
};
</script>
<script setup>
import { useSiteStore } from "~/pinia/modules/site";
const siteStore = useSiteStore();
</script>
<style lang="scss" scoped>
.footer {
  width: 100%;
  text-align: center;
  height: 40px;
}
#footer-line {
  margin: 10px 0;
}
.rebuilding {
  margin: 0;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  position: fixed;
  display: block;
  background-color: rgba(0, 0, 0, 0.7);
  z-index: 2000;
  text-align: center;
  padding-top: 40vh;
}
</style>