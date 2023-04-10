<template>
  <div class="flex flex-wrap h-screen text-center justify-around items-center">
    <div class="font-blod desc">
      <div class="text-7xl err-code">Error 403</div>
      <div ref="content" class="text-3xl content"></div>
      <button
        active="scale-90 transform"
        class="rounded-lg transition btn"
        @click="back"
      >
        👉 返回首页
      </button>
    </div>

    <img src="../../assets/icons/401.svg" class="cover" alt="page not found" />
  </div>
</template>

<script>
export default {
  name: "403",
};
</script>
<script setup>
import { ref, onMounted } from "vue";
import { init } from "ityped";
import router from "~/router/index";

const content = ref(null);
onMounted(() => {
  init(content.value, {
    showCursor: false,
    disableBackTyping: true,
    strings: ["禁止访问此页面!"],
  });
});
const back = () => {
  router.push("/");
};
</script>
<style lang="scss" scoped>
@tailwind base;
@tailwind components;
@tailwind utilities;
.err-code {
  margin-bottom: 20px;
}
.content {
  height: 40px;
}
.cover {
  height: auto;
  width: 600px;
  margin: 0 100px;
  max-width: 100%;
  max-height: 100%;
}
.desc {
  flex: 1;
  width: 300px;
}
.btn {
  font-size: 20px;
  margin-top: 30px;
  padding: 12px 20px;
  --tw-shadow-color: 8, 145, 178;
  --tw-shadow: 0 4px 6px -1px rgba(var(--tw-shadow-color), 0.1),
    0 2px 4px -1px rgba(var(--tw-shadow-color), 0.06);
  box-shadow: 0 0 #0000, 0 0 #0000, 0 0 #0000, var(--tw-shadow);
}
</style>