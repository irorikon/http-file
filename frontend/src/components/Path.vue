<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 22:53:07
 * @FilePath: \http-file\frontend\src\components\Path.vue
-->
<template>
  <div class="path">
    <a-breadcrumb-item href="">
      <home style="margin-right: 10px" @click="defaultPath" />
    </a-breadcrumb-item>
    <a-breadcrumb>
      <a-breadcrumb-item
        v-for="(value, index) in globalStore.paths"
        :key="index"
      >
        <a :selectPath="index" @click="changePath(index)">{{ value }}</a>
        <!-- <template #overlay v-if="isNotEmpty(value)">
          <a-menu>
            <a-menu-item v-for="(i, j) in value" :key="j">
              <a @click="changePath">{{ i }}</a>
            </a-menu-item>
          </a-menu>
        </template> -->
      </a-breadcrumb-item>
    </a-breadcrumb>
  </div>
</template>

<script>
export default {
  name: "Path",
};
</script>
<script setup>
import { onMounted, ref } from "vue";
import { isNotEmpty } from "~/utils/common";
import { useGlobalStore } from "~/pinia/modules/global";
import { useFileStore } from "~/pinia/modules/file";

const globalStore = useGlobalStore();
const fileStore = useFileStore();

const selectPath = ref("");

// 查询文件目录
const searchFileList = (path) => {
  fileStore.getFiles(path, globalStore.storage_type);
  globalStore.setPaths("/");
  globalStore.setPaths(path.split("\\").join("/"));
};

// 根据用户点击切换目录列表
const changePath = (e) => {
  searchFileList(globalStore.paths.slice(0, e + 1).join("/"));
};

// 查询默认文件目录
const defaultPath = () => {
  selectPath.value = "/";
  searchFileList(selectPath.value);
};
onMounted(() => {
  if (!isNotEmpty(globalStore.storage_type)) {
    globalStore.setStorageType("local");
  }
  defaultPath();
});
</script>
<style lang="scss" scoped>
.path {
  /* box-shadow: 0 2px 4px rgba(0, 0, 0, .12), 0 0 6px rgba(0, 0, 0, .04); */
  margin: 2px 2px;
  padding: 2px;
  display: flex;
  display: -webkit-flex;
  font-size: 20px;
  align-items: center;
}
.path-item {
  color: var(--textColor, black);
}
</style>