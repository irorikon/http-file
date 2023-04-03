<!--
 * @Author: iRorikon
 * @Date: 2023-04-03 21:04:38
 * @FilePath: \http-file\src\components\Layout.vue
-->
<template>
  <d-layout>
    <d-header>
      <Header />
    </d-header>
    <d-content>
      <Content />
    </d-content>
    <d-footer>
      <Footer />
    </d-footer>
  </d-layout>
</template>
<script setup>
import { ref } from "vue";
import Content from "./Content.vue";
import Footer from "./Footer.vue";
import Header from "./Header.vue";

defineProps({
  msg: String,
});

const count = ref(0);
</script>
<script>
export default {
  name: "Layout",
  components: { Header, Content, Footer },
};
</script>
<style lang="sass" scoped>
.read-the-docs
  color: #888
</style>