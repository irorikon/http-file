<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 22:47:00
 * @FilePath: \http-file\frontend\src\components\Header.vue
-->
<template>
  <div>
    <div class="header">
      <router-link to="/">
        <img
          v-if="siteStore.siteInfo.logo"
          :src="siteStore.siteInfo.logo"
          alt="HTTP-File"
          class="logo"
          id="logo"
        />
      </router-link>
      <span class="title">HTTP-File</span>
      <a-space class="right">
        <a-space>
          <a-input-search
            :value="keyword.value"
            placeholder="搜索文件(夹)"
            :loading="globalStore.loading"
            enter-button
            @search="onSearch"
          />
        </a-space>
      </a-space>
    </div>
    <br />
  </div>
</template>
<script>
export default {
  name: "Header",
};
</script>
<script setup>
import { useSiteStore } from "~/pinia/modules/site";
import { useGlobalStore } from "~/pinia/modules/global";
import { useFileStore } from "~/pinia/modules/file";
import { searchFile } from "~/api/file";
import { ref } from "vue";

const siteStore = useSiteStore();
const globalStore = useGlobalStore();
const fileStore = useFileStore();
const keyword = ref("");
const onSearch = async (searchValue) => {
  // TODO: 根据 Search Value 重新获取 File 列表
  globalStore.setLoading(true);
  const res = await searchFile({
    storage_type: globalStore.storage_type,
    path: globalStore.path,
    keyword: searchValue,
  });
  if (res.code === 0) {
    fileStore.setFilelist(res.data);
    globalStore.setLoading(false);
    // 清空缓存
    fileStore.setFile({});
  }
};
</script>
<style lang="scss" scoped>
.header {
  padding-top: 20px;
  height: 56px;
  width: min(98vw, 800px);
  display: flex;
  display: -webkit-flex; /* Safari */
  justify-content: space-between;
  align-items: center;
  width: 100%;
  text-align: center;
}
.header .logo {
  height: 48px;
  width: auto;
  margin-right: 5px;
}
.header .title {
  font-size: 20px;
  font-weight: bold;
  margin-right: 10px;
  color: #40a9ff;
  flex: 1;
  /* flex-grow: 1; */
}
.header .right {
  align-items: center;
  justify-content: flex-end;
  margin-left: 10px;
}
.header-content {
  margin: 10px 0 5px 0;
}
</style>