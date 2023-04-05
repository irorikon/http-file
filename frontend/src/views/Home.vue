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
          <a-spin v-if="type === 'loading'" size="large" />
          <Files v-if="type === 'folder'" />
          <Readme v-if="type === 'folder'" />
          <Preview v-if="type === 'file'" />
          <NotFound v-if="type === 'no'" />
        </div>
        <Footer />
      </div>
    </div>
    <a-modal
      :visible="showPassword"
      title="请输入密码"
      @ok="okPassword"
      @cancel="cancelPassword"
    >
      <a-input-password
        ref="inputRef"
        placeholder="此处为盘符密码"
        :value="password"
        @pressEnter="okPassword"
      />
    </a-modal>
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
    Preview,
    NotFound,
    Readme,
  },
};
</script>
<script setup>
import { computed, defineComponent, onMounted, ref, watch } from "vue";
import Header from "../components/Header.vue";
import Footer from "../components/Footer.vue";
import Path from "../components/Path.vue";
import Files from "../components/Files.vue";
import Readme from "../components/Readme.vue";
import Preview from "../components/Preview.vue";
import NotFound from "../components/NotFound.vue";
import { GlobalDataProps } from "../store";
import { useRoute, useRouter } from "vue-router";
import { useStore } from "vuex";
import useRefresh from "../hooks/useRefresh";

const route = useRoute();
const router = useRouter();
const store = useStore();
const type = computed(() => store.state.type);
const { refresh } = useRefresh();
watch(
  () => route.fullPath,
  () => {
    if (route.fullPath !== "/") {
      refresh();
    } else {
      if (store.state.info.roots) {
        router.push(store.state.info.roots[0]);
      }
    }
  }
);
onMounted(() => {
  store.dispatch("fetchInfo").then(() => {
    if (route.fullPath !== "/") {
      refresh();
    } else {
      if (store.state.info.roots) {
        router.push(store.state.info.roots[0]);
      }
    }
  });
});
const showPassword = computed(() => {
  return store.state.meta.code === 401;
});
const inputRef = ref(null);
watch(showPassword, () => {
  setTimeout(() => {
    if (showPassword.value && inputRef.value) {
      inputRef.value.focus();
    }
  }, 50);
});
const password = ref(store.state.password);
const okPassword = () => {
  store.commit("setPassword", password.value);
  refresh();
};
const cancelPassword = () => {
  route.go(-1);
};
</script>
<style lang="sass" scoped>
.home
  color: var(--textColor, black)
  width: 100%
  display: flex
  display: -webkit-flex
  justify-content: center
  padding: 0
  margin: 0
.layout
  display: flex
  display: -webkit-flex
  flex-direction: column
  justify-content: center
  align-items: center
  width: min(98vw, 980px)

@media screen and (max-width: 980px)
.layout
  width: 98vw

.content
  min-height: 80vh
  width: 100%
  display: flex
  display: -webkit-flex
  flex-direction: column
  justify-content: flex-start
</style>
