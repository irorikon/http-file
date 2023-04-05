<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 22:39:47
 * @FilePath: \http-file\frontend\src\components\Footer.vue
-->
<template>
  <div>
    <a-divider id="footer-line" />
    <div class="footer">
      Powered By
      <a target="_blank" href="https://github.com/Xhofe/alist">AList</a>
      <a-divider type="vertical" />
      <a-button type="link" size="small" @click="preRebuild">重构目录</a-button>
      <a-divider v-if="info.footer_text" type="vertical" />
      <a v-if="info.footer_text" target="_blank" :href="info.footer_url">{{
        info.footer_text
      }}</a>
    </div>
    <a-modal
      :visible="showPassword"
      title="重建目录(深度为负数则不限制)"
      @ok="rebuild"
      @cancel="showPassword = false"
    >
      <a-input
        ref="inputRef"
        placeholder="重建目录深度"
        :value="depth"
        type="number"
        @pressEnter="rebuild"
      ></a-input>
      <div style="margin: 24px 0" />
      <a-input-password
        placeholder="重建目录密码"
        :value="password"
        @pressEnter="rebuild"
      />
    </a-modal>
    <div class="rebuilding" v-if="rebuilding">
      <a-spin size="large" tip="重建目录中..." />
    </div>
  </div>
</template>
<script>
export default {
  name: "Footer",
};
</script>
<script setup>
import { GlobalDataProps } from "../store";
import { rebuildPost } from "../utils/api";
import { message } from "ant-design-vue";
import { computed, defineComponent, ref } from "vue";
import { useStore } from "vuex";
import useRefresh from "../hooks/useRefresh";
import { useRoute } from "vue-router";

const store = useStore();
const info = computed(() => store.state.info);
const showPassword = ref(false);
const password = ref(localStorage.getItem("rebuild-password") || "");
const { refresh } = useRefresh();
const rebuilding = ref(false);
const route = useRoute();
const depth = ref(3);
const rebuild = () => {
  localStorage.setItem("rebuild-password", password.value);
  showPassword.value = false;
  rebuilding.value = true;
  rebuildPost(
    decodeURI(route.path.substring(1)),
    password.value,
    depth.value
  ).then((res) => {
    rebuilding.value = false;
    if (res.data.code === 200) {
      message.success(res.data.message);
      refresh();
    } else {
      message.error(res.data.message);
    }
  });
};
const inputRef = ref();
const preRebuild = () => {
  showPassword.value = true;
  setTimeout(() => {
    inputRef.value.focus();
  }, 50);
};
</script>
<style lang="sass" scoped>
.footer
  width: 100%
  text-align: center
  height: 40px

#footer-line
  margin: 10px 0

.rebuilding
  margin: 0
  top: 0
  right: 0
  bottom: 0
  left: 0
  position: fixed
  display: block
  background-color: rgba(0, 0, 0, 0.7)
  z-index: 2000
  text-align: center
  padding-top: 40vh
</style>