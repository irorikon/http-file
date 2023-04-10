<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 21:56:02
 * @FilePath: \http-file\frontend\src\components\Files.vue
-->
<template>
  <div class="files">
    <div class="files-list" v-contextmenu>
      <a-table
        :columns="columns"
        :data-source="files"
        :pagination="false"
        :customRow="customRow"
        :loading="globalStore.loading"
        :scroll="{ x: 'max-content' }"
        rowKey="name"
      >
        <template #bodyCell="{ text, record, column }">
          <template v-if="column.dataIndex === 'name'">
            <component :is="record.icon" class="file-icon" />
            {{ text }}
            <span v-if="!record.is_dir" class="action">
              <!-- <copy id="action-1" @click="copyFileLink(record)" /> -->
              <copy id="action-1" />
            </span>
          </template>
        </template>
      </a-table>
      <a-modal
        :visible="showDownloadPanle"
        ok-text="下载文件"
        @ok="download(selectFile)"
        @cancel="showDownloadPanle = false"
      >
        <a-result :title="selectFile.name">
          <template #icon>
            <smile-twoTone />
            <component :is="selectFile.icon" />
          </template>
        </a-result>
      </a-modal>
    </div>
    <br />
  </div>
</template>
<script>
export default {
  name: "Files",
};
</script>
<script setup>
import { computed, ref } from "vue";
import { formatDate } from "~/utils/date";
import { getIcon, getCategory, getFileSize, getFileLink } from "~/utils/common";
import { message } from "ant-design-vue";
import { useSiteStore } from "~/pinia/modules/site";
import { useGlobalStore } from "~/pinia/modules/global";
import { useFileStore } from "~/pinia/modules/file";

const siteStore = useSiteStore();
const globalStore = useGlobalStore();
const fileStore = useFileStore();

const sort = siteStore.siteInfo.sort.split("-") || ["", ""];
const columns = [
  {
    align: "left",
    dataIndex: "name",
    title: "文件名称",
    customRender: (text, record) => {},
    sorter: (a, b) => {
      return a.name < b.name ? 1 : -1;
    },
    defaultSortOrder: sort[0] === "name" ? sort[1] : undefined,
  },
  {
    align: "right",
    dataIndex: "sizeStr",
    title: "大小",
    width: 100,
    sorter: (a, b) => {
      return a.size - b.size;
    },
    defaultSortOrder: sort[0] === "size" ? sort[1] : undefined,
  },
  {
    align: "right",
    dataIndex: "time",
    title: "修改时间",
    width: 170,
    sorter: (a, b) => {
      return a.time < b.time ? 1 : -1;
    },
    defaultSortOrder: sort[0] === "time" ? sort[1] : undefined,
  },
];
const files = computed(() => {
  const data = fileStore.filelist;
  return data.map((item) => {
    item.time = formatDate(item.modified);
    if (item.is_dir) {
      item.sizeStr = "-";
    } else {
      item.sizeStr = getFileSize(item.size);
    }
    item.icon = getIcon(item);
    item.category = getCategory(item);
    return item;
  });
});
const selectFile = ref(null);
const showDownloadPanle = ref(false);
const customRow = (record) => {
  return {
    onClick: (e) => {
      if (e.target && e.target.tagName === "svg") {
        return;
      }
      if (record.is_dir) {
        fileStore.getFiles(
          record.path.split("\\").join("/"),
          globalStore.storage_type
        );
        globalStore.setPaths(record.path.split("\\").join("/"));
      } else {
        selectFile.value = record;
        showDownloadPanle.value = true;
      }
    },
  };
};

const download = (item) => {
  if (!item.is_dir) {
    window.open(getFileLink(item), "_blank");
  }
  showDownloadPanle.value = false;
};
</script>
<style lang="scss" scoped>
.file-icon {
  margin-right: 10px;
  font-size: 20px;
  color: #1890ff;
}
.action {
  color: gray;
}
#images {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  /* grid-gap: 24px; */
  list-style: none;
  justify-items: center;
}
.image {
  width: 100px;
  height: 90px;
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-shrink: 0;
  flex-grow: 0;
}
.image img {
  display: block;
  max-width: 100%;
  max-height: 100%;
  border-radius: 5px;
  position: relative;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
}
</style>