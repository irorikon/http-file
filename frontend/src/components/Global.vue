<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 20:04:34
 * @FilePath: \http-file\frontend\src\components\Global.vue
-->
<template>
  <div id="global">
    <MyIcon
      class="icon"
      :type="isDark ? 'icon-Sun' : 'icon-moon'"
      @click="change"
    />
  </div>
</template>
<script>
export default {
  name: "Global",
};
</script>
<script setup>
import { ref } from "vue";
import { createFromIconfontCN } from "@ant-design/icons-vue";
const MyIcon = createFromIconfontCN({
  scriptUrl: "//at.alicdn.com/t/font_2594335_ba51zkwh1a4.js",
});
const darkMedia = window.matchMedia("(prefers-color-scheme: dark)");
const isDark = ref(darkMedia.matches);
const change = () => {
  isDark.value = !isDark.value;
  if (isDark.value) {
    document.querySelector("body")?.style.setProperty("--bgColor", "#27282b");
    document.querySelector("body")?.style.setProperty("--textColor", "#d9d9d9");
    document.querySelector("body")?.style.setProperty("--bg2Color", "#4e4e4e");
    document
      .querySelector("body")
      ?.style.setProperty("--lineColor", "transparent");
  } else {
    document.querySelector("body")?.style.removeProperty("--bgColor");
    document.querySelector("body")?.style.removeProperty("--textColor");
    document.querySelector("body")?.style.removeProperty("--bg2Color");
    document.querySelector("body")?.style.removeProperty("--lineColor");
  }
};
darkMedia.onchange = () => {
  isDark.value = darkMedia.matches;
  change();
};
change();
</script>
<style lang="scss" scoped>
#global {
  position: fixed;
  top: 50%;
  transform: translate(0, -50%);
  right: 5px;
  z-index: 999;
}
.icon {
  font-size: 25px;
}
/* @media screen and (max-width: 1000px) {
    #global {
      display: none;
    }
  } */
</style>