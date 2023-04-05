<template>
  <div class="readme" v-show="readme !== undefined">
    <a-spin :spinning="readmeSpinning">
      <a-card title="Readme.md" style="width: 100%" size="small">
        <v-md-preview :text="readmeValue"></v-md-preview>
      </a-card>
    </a-spin>
  </div>
</template>
<script>
export default {
  name: "Readme",
};
</script>
<script setup>
import { FileProps, GlobalDataProps } from "../store";
import { useStore } from "vuex";
import { computed, defineComponent, ref, watch } from "vue";
import { getPost, getText } from "../utils/api";

const store = useStore();
const type = computed(() => store.state.type);
const readmeSpinning = ref(true);
const readmeValue = ref("");
const refreshReadme = (readmeFile) => {
  if (readmeFile !== undefined) {
    getPost(readmeFile.dir + readmeFile.name, store.state.password).then(
      (resp) => {
        const res = resp.data;
        if (res.code === 200) {
          getText(res.data.url).then((resp) => {
            readmeValue.value = resp.data;
            readmeSpinning.value = false;
          });
        }
      }
    );
  }
};
const readme = computed(() => {
  const file = store.state.data;
  if (file.type) {
    return undefined;
  }
  const files = store.state.data;
  const readmeFile = files.find(
    (item) => item.name.toLowerCase() === "readme.md"
  );
  refreshReadme(readmeFile);
  return readmeFile;
});
</script>