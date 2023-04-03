<!--
 * @Author: iRorikon
 * @Date: 2023-04-03 21:19:18
 * @FilePath: \http-file\src\components\Header.vue
-->
<template>
  <div class="dheader-1">
    <img
      class="logo-devui"
      src="https://devui.design/components/assets/logo.svg"
      alt=""
    />
    <span class="text">DevUI</span>
    <div class="right-search-box">
      <d-search icon-position="left" no-border style="width: 100%"></d-search>
    </div>
    <d-button @click="handleClick" class="right-buton">上传文件</d-button>
    <d-modal v-model="visible" style="width: auto; height: auto">
      <div class="modal">
        <d-upload
          class="upload-box"
          v-model="uploadedFiles"
          :upload-options="uploadOptions"
          droppable
          :on-success="onSuccess"
          :on-error="onError"
          :on-progress="onProgress"
          @file-drop="fileDrop"
          @file-select="fileSelect"
        >
          <div class="upload-trigger">
            <div><d-icon name="upload" size="24px"></d-icon></div>
            <div style="margin-top: 20px">
              将文件拖到此处，或
              <span class="link">点击上传</span>
            </div>
          </div>
          <template v-slot:uploaded-files="slotProps">
            <table
              class="table uploaded-files"
              v-if="slotProps.uploadedFiles.length > 0"
            >
              <tbody>
                <tr
                  v-for="(uploadedFile, index) in slotProps.uploadedFiles"
                  :key="index"
                  class="row"
                >
                  <td width="75%">
                    <span>{{ uploadedFile.name }}</span>
                  </td>
                  <td width="25%">
                    <span>上传成功</span>
                  </td>
                </tr>
              </tbody>
            </table>
          </template>
        </d-upload>
      </div>
    </d-modal>
  </div>
</template>
<script>
export default {
  name: "Header",
};
</script>
<script setup>
import { ref, defineComponent, reactive, watch } from "vue";
import { NotificationService } from "#/vue-devui/notification";
const visible = ref(false);
const handleClick = () => {
  visible.value = true;
};
const hidden = () => {
  visible.value = false;
};

const uploadedFiles = [];
const uploadOptions = {
  uri: "https://www.mocky.io/v2/5cc8019d300000980a055e76",
};
const onSuccess = (res) => {
  notifyMsg.title = "上传成功";
  notifyMsg.content = "文件上传成功: " + res.file.name;
  notifyMsg.type = "success";
  showService(notifyMsg);
};
const onError = (err) => {
  console.log("err", err);
  notifyMsg.title = "上传失败";
  notifyMsg.content = "文件上传失败: " + err.file.name;
  notifyMsg.type = "error";
  showService(notifyMsg);
};
const onProgress = (selectFile, uploadedFiles) => {
  console.log("selectFile: ", selectFile);
  console.log("uploadedFiles: ", uploadedFiles);
};
const fileDrop = (file) => {
  for (let i = 0; i < file.length; i++) {
    console.log("filesize", file[i].size);
    const isLt2M = file[i].size / 1024 / 1024 < 2;
    if (!isLt2M) {
      notifyMsg.title = "上传失败";
      notifyMsg.content = "文件大小不能超过 2MB!";
      notifyMsg.type = "warning";
      showService(notifyMsg);
    }
  }
};
const fileSelect = (file) => {
  for (let i = 0; i < file.length; i++) {
    console.log("filesize", file[i].size);
    const isLt2M = file[i].size / 1024 / 1024 < 2;
    if (!isLt2M) {
      notifyMsg.title = "上传失败";
      notifyMsg.content = "文件大小不能超过 2MB!";
      notifyMsg.type = "warning";
      showService(notifyMsg);
    }
  }
};
watch(uploadedFiles, (newValue, oldValue) => {
  console.log("uploadedFiles", {
    newValue,
    oldValue,
  });
});
const notifyMsg = {
  title: "",
  content: "",
  type: "",
};
const showService = (msg) => {
  NotificationService.open(msg);
};
</script>
<style lang="sass" scoped>
.dheader-1
  text-align: left
  position: relative
  height: 60px
  background-color: #333854
  color: #fffS
.dheader-1
  position: relative
  height: 60px
  display: flex
  align-items: center
  background-color: #333854
  .logo-devui
    width: 48px
    height: 48px
    margin: 0 10px
  .text
    color: #fff
    font-size: 24px
  .right-search-box
    margin-left: auto
    text-align: right
    margin-right: 10px
    cursor: pointer
  .right-buton
    color: #fff
    font-size: 20px
    text-align: right
    margin-right: 10px
    cursor: pointer
    background-color: #333854

.upload-box .upload-trigger
  background-color: #fff
  border: 1px dashed #d9d9d9
  border-radius: 6px
  box-sizing: border-box
  width: 600px
  height: 300px
  text-align: center
  cursor: pointer
  position: relative
  overflow: hidden
  display: flex
  flex-direction: column
  justify-content: center
  align-items: center

  & .link
    color: #5e7ce0
.modal
  display: flex
  justify-content: center
  align-items: center
</style>