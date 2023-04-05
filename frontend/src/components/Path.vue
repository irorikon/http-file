<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 22:53:07
 * @FilePath: \http-file\frontend\src\components\Path.vue
-->
<template>
  <div class="path">
    <router-link :to="'/' + routes[0].path">
      <home style="margin-right: 10px" />
    </router-link>
    <a-breadcrumb :routes="routes">
      <template #itemRender="{ route, routes, paths }">
        <span
          v-if="routes.indexOf(route) === 0"
          :to="`/${paths.join('/')}`"
          class="path-item"
        >
          <!-- <home style="margin-right: 1px;"/> -->
          {{ route.breadcrumbName }}
        </span>
        <router-link
          v-else-if="routes[0].children.indexOf(route) !== -1"
          :to="`/${paths[1]}`"
          class="path-item"
        >
          {{ route.breadcrumbName }}
        </router-link>
        <span
          v-else-if="!q && routes.indexOf(route) === routes.length - 1"
          class="path-item"
        >
          {{ route.breadcrumbName }}
        </span>
        <router-link v-else :to="`/${paths.join('/')}`" class="path-item">
          {{ route.breadcrumbName }}
        </router-link>
      </template>
    </a-breadcrumb>
  </div>
</template>

<script>
export default {
  name: "Path",
};
</script>
<script setup>
import store from "../store";
import { computed, defineComponent } from "vue";
import { useRoute } from "vue-router";

const route = useRoute();
const q = computed(() => route.query.q);
const routes = computed(() => {
  const paths = decodeURI(route.fullPath.substring(1)).split("?")[0].split("/");
  if (!paths) {
    return [{ path: "/", breadcrumbName: "home" }];
  }
  const res = paths.map((item) => {
    return {
      path: item,
      breadcrumbName: item,
    };
  });
  res[0].children =
    store.state.info.roots?.map((item) => {
      return {
        path: item,
        breadcrumbName: item,
      };
    }) || [];
  return res;
});
</script>
<style lang="sass" scoped>
.path
  margin: 2px 2px
  padding: 2px
  display: flex
  display: -webkit-flex
  font-size: 20px
  align-items: center

  &-item
    color: var(--textColor, black)
</style>