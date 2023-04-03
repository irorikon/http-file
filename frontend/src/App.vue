<!--
 * @Author: iRorikon
 * @Date: 2023-04-03 21:04:38
 * @FilePath: \http-file\src\App.vue
-->
<script setup>
import Layout from "./components/Layout.vue";
</script>

<template>
  <Layout msg="Vite + Vue" />
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
