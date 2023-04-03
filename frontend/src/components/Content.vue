<!--
 * @Author: iRorikon
 * @Date: 2023-04-03 21:20:47
 * @FilePath: \http-file\src\components\Content.vue
-->
<template>
  <div class="dcontent-1">
    <d-breadcrumb class="dbreadcrumb">
      <d-breadcrumb-item>
        <a href="/"><span>首页</span></a>
      </d-breadcrumb-item>
      <d-breadcrumb-item v-for="breadcrumb in breadcrumbs" :key="breadcrumb">
        <a :href="breadcrumb.href"
          ><span>{{ breadcrumb.title }}</span></a
        >
      </d-breadcrumb-item>
    </d-breadcrumb>
  </div>
</template>
<script>
export default {
  name: "Content",
};
</script>
<script setup>
import { ref } from "vue";
const breadcrumbs = [
  {
    title: "一级",
    href: "config",
  },
  {
    title: "二级",
    href: "config",
  },
  {
    title: "三级",
    href: "config",
  },
];
</script>
<style lang="sass" scoped>
.dcontent-1
  padding: 0 40px
  height: 300px
  background-color: #f3f6f8

  .dbreadcrumb
    margin: 8px 0

  .inner-content
    background-color: #fff
    height: 100%
</style>