<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 10:53:34
 * @FilePath: \http-file\frontend\src\components\Footer.vue
-->
<template>
  <div>
    <div class="dfooter-1">
      <p>
        Copyright © 2020-PRESENT
        <a href="https://www.github.com/irorikon" class="link">iRorikon</a>
      </p>
    </div>
  </div>
</template>
<script>
export default {
  name: "Footer",
};
</script>
<script setup>
import { ref } from "vue";
</script>
<style lang="sass" scoped>
.dfooter-1
  background: #333854
  color: #fff
  text-align: center
  line-height: 40px
  min-height: 40px
  & .link
    color: #5e7ce0

.progress-container-circle
  height: 130px
  width: 130px
  font-size: 20px
  display: inline-block
  margin-right: 10px

.icon-position
  position: absolute
  top: 50%
  left: 50%
  width: 100%
  margin: 0
  padding: 0
  line-height: 1
  white-space: normal
  text-align: center
  transform: translate(-50%, -50%)
  color: #50d4ab
  font-size: 24px
</style>