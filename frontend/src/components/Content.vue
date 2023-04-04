<!--
 * @Author: iRorikon
 * @Date: 2023-04-03 21:20:47
 * @FilePath: \http-file\frontend\src\components\Content.vue
-->
<template>
  <div>
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
    <div class="table">
      <d-table
        ref="tableRef"
        table-layout="auto"
        header-bg="true"
        size="md"
        :data="stripedTableData"
        :row-key="(item) => item.id"
        :show-loading="showLoading"
        @sort-change="handleSortChange"
        @row-click="onRowClick"
        @check-change="checkChange"
        @check-all-change="checkAllChange"
      >
        <d-column type="checkable" width="40" reserve-check />
        <d-column
          field="isDir"
          header="#"
          sortable
          :sort-method="sortDirMethod"
        >
          <template #default="scope">
            <i class="icon icon-folder" v-show="scope.row.isDir"></i>
            <i class="icon icon-file" v-show="!scope.row.isDir"></i>
          </template>
        </d-column>
        <d-column
          field="filename"
          header="名称"
          sortable
          :sort-method="sortNameMethod"
          min-width="150"
        ></d-column>
        <d-column field="size" header="大小"></d-column>
        <d-column field="date" header="时间"></d-column>
        <d-column field="operate" header="操作">
          <template #default="scope">
            <d-button @click="handleClick(scope.row)">删除</d-button>
          </template>
        </d-column>
      </d-table>
    </div>
  </div>
</template>
<script>
export default {
  name: "Content",
};
</script>
<script setup>
import { ref } from "vue";
const showLoading = ref(false);
const breadcrumbs = [
  {
    id: 1,
    title: "一级",
    href: "config",
  },
  {
    id: 2,
    title: "二级",
    href: "config",
  },
  {
    id: 3,
    title: "三级",
    href: "config",
  },
];
const handleSortChange = (sort) => {
  console.log(sort);
};

const sortNameMethod = (a, b) => {
  return a.firstName < b.firstName;
};

const sortDirMethod = (a, b) => {
  return a.isDir > b.isDir;
};

const stripedTableData = ref([]);
const getData = () => {
  showLoading.value = true;
  setTimeout(() => {
    showLoading.value = false;
    stripedTableData.value = [
      {
        id: 1,
        isDir: true,
        filename: "文件1",
        size: "10M",
        date: "2020-01-01",
      },
      {
        id: 2,
        isDir: true,
        filename: "文件2",
        size: "10M",
        date: "2020-01-01",
      },
      {
        id: 3,
        isDir: false,
        filename: "文件3",
        size: "10M",
        date: "2020-01-01",
      },
      {
        id: 4,
        isDir: false,
        filename: "文件4",
        size: "10M",
        date: "2020-01-01",
      },
    ];
  }, 1000);
};
getData();

var selectRow = [];
const tableRef = ref(null);
const checkChange = (checked, row, selection) => {
  if (checked) {
    selectRow.push(row);
  } else {
    var new_arr = [];
    for (let i = 0; i < selectRow.length; i++) {
      if (selectRow[i].id != row.id) {
        new_arr.push(selectRow[i]);
      }
    }
    selectRow = new_arr;
  }
};

const checkAllChange = (checked, selection) => {
  if (checked) {
    selectRow = selection;
  } else {
    selectRow = [];
  }
};
const onRowClick = (params) => {
  console.log("row-click", params);
};
</script>
<style lang="sass" scoped>
.dcontent-1
  padding: 0 40px
  height: auto
  display: flex
  background-color: #f3f6f8

  .dbreadcrumb
    margin: 8px 0

  .inner-content
    background-color: #fff
    height: 100%

.table-btn-groups
  display: flex
  flex-wrap: wrap
.table-btn
  display: flex
  justify-content: flex-start
  align-items: center
  margin-right: 1rem
.table
  width: 100%
  height: 100%
  display: flex
  justify-content: center
  text-align: center
  align-items: center
</style>