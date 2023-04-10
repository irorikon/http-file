<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 22:21:37
 * @FilePath: \http-file\frontend\src\views\Home.vue
-->
<template>
  <div>
    <div class="home">
      <div class="layout">
        <Header />
        <div class="content">
          <div class="tool">
            <Path />
          </div>
          <a-divider style="margin: 10px 0 5px 0" />
          <Files />
        </div>
        <Footer />
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "Home",
  components: {
    Header,
    Footer,
    Path,
    Files,
  },
};
</script>
<script setup>
import { onMounted } from "vue";
import Header from "~/components/Header.vue";
import Footer from "~/components/Footer.vue";
import Path from "~/components/Path.vue";
import Files from "~/components/Files.vue";
import { useSiteStore } from "~/pinia/modules/site";
import { useGlobalStore } from "~/pinia/modules/global";

const siteStore = useSiteStore();
onMounted(() => {
  setTimeout(() => {
    siteStore.getSiteInfo();
    return;
  });
});
</script>
<style lang="scss" scoped>
.home {
  /* background-color: var(--bgColor, white); */
  color: var(--textColor, black);
  width: 100%;
  display: flex;
  display: -webkit-flex; /* Safari */
  justify-content: center;
  padding: 0;
  margin: 0;
  /* filter: invert(1) hue-rotate(180deg); */
}
.layout {
  display: flex;
  display: -webkit-flex; /* Safari */
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: min(98vw, 1000px);
}
@media screen and (max-width: 1000px) {
  .layout {
    width: 98vw;
  }
}
.content {
  min-height: 80vh;
  width: 100%;
  display: flex;
  display: -webkit-flex;
  flex-direction: column;
  justify-content: flex-start;
}
</style>
