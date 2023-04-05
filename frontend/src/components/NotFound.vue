<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 22:52:30
 * @FilePath: \http-file\frontend\src\components\NotFound.vue
-->
<template>
  <div>
    <div class="not-found">
      <a-result status="404" title="404" sub-title="抱歉,找不到信息...">
        <template #extra>
          <a-button type="primary">
            <router-link to="/"> 返回首页 </router-link>
          </a-button>
        </template>
      </a-result>
    </div>
  </div>
</template>

<script>
export default {
  name: "NotFound",
};
</script>