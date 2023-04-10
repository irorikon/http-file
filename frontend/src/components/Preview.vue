<!--
 * @Author: iRorikon
 * @Date: 2023-04-04 22:57:55
 * @FilePath: \http-file\frontend\src\components\Preview.vue
-->
<template>
  <div class="preview">
    <!-- 无法预览，显示下载与直链 -->
    <a-result :title="file.name" v-if="show.other">
      <template #icon>
        <component :is="file.icon" class="file-icon" />
      </template>
      <template #extra>
        <a target="_blank" :href="downloadLink">
          <a-button type="primary" :download="file.name">下载</a-button>
        </a>
      </template>
    </a-result>
    <!-- 显示加载的部分 -->
    <a-spin :spinning="previewSpinning" v-if="show.spinning">
      <!-- 文档预览 -->
      <div class="doc-preview" id="doc-preview" v-if="show.doc"></div>
      <!-- iframe -->
      <iframe
        :src="downloadLink"
        id="iframe-preview"
        ref="iframe-preview"
        v-if="show.iframe"
        :allowfullscreen="allowfullscreen"
        webkitallowfullscreen="true"
        mozallowfullscreen="true"
        class="iframe-preview"
        frameborder="no"
        @load="previewSpinning = false"
      ></iframe>
      <!-- 图片预览 -->
      <div class="img-preview" v-if="show.image">
        <img @load="previewSpinning = false" :src="downloadLink" />
      </div>
      <!-- 文本预览 -->
      <div class="text-preview" v-if="show.text">
        <v-md-preview :text="text"></v-md-preview>
      </div>
    </a-spin>
    <!-- 视频预览 -->
    <div v-if="show.video" ref="artRef"></div>
    <!-- 音频预览 -->
    <div class="audio-preview" v-if="show.audio" id="audio-preview"></div>
    <!-- HTML预览 -->
    <iframe
      class="html-preview"
      v-if="show.html"
      id="html-preview"
      v-bind:srcdoc="text"
      style="width: 100%; box-shadow: #00000031 0px 1px 10px 5px"
      frameborder="0"
    ></iframe>
  </div>
</template>
<script>
export default {
  name: "Preview",
};
</script>
<script setup>
import {
  getIcon,
  getCategory,
  getDownloadLink,
  isNotEmpty,
} from "~/utils/common";
import request from "~/utils/request";
import { encodePath } from "~/utils/base64";
import { onBeforeUnmount, onMounted, ref } from "vue";
import Artplayer from "artplayer";
import APlayer from "aplayer";
import "aplayer/dist/APlayer.min.css";
import Hls from "artplayer-plugin-hls-quality";
import flvjs from "flv.js";
import dashjs from "artplayer-plugin-dash-quality";

import { useSiteStore } from "~/pinia/modules/site";
import { useFileStore } from "~/pinia/modules/file";

const siteStore = useSiteStore();
const fileStore = useFileStore();

const setIcon = () => {
  const res = fileStore.file;
  res.icon = getIcon(fileStore.file);
  return res;
};
const file = ref({});
const artRef = ref(null);

const downloadLink = ref("");

const previewSpinning = ref(true);
const show = ref({
  image: false,
  video: false,
  audio: false,
  doc: false,
  other: false,
  text: false,
  iframe: false,
  spinning: false,
  html: false,
});

const text = ref("");
let ap, art;
// 预览
const preview = async (file) => {
  const category = getCategory(file);
  downloadLink.value = getDownloadLink(file);
  if (file.ext === "html") {
    show.value.html = true;
    previewSpinning.value = true;
    show.value.spinning = true;
    request({
      url: "file/d/" + encodePath(file),
      method: "get",
    }).then((res) => {
      text.value = res;
      setTimeout(function () {
        const htmlDom = window.frames["html-preview"];
        if (!htmlDom) return;
        htmlDom.setAttribute(
          "height",
          htmlDom.contentWindow.document.body.scrollHeight
        );
      }, 200);
      previewSpinning.value = false;
    });
    return;
    // } else if (isNotEmpty(file.ext)) {
    //   if (siteStore.siteInfo.preview.text.includes(file.ext.toLowerCase())) {
    //     show.value.text = true;
    //     previewSpinning.value = true;
    //     show.value.spinning = true;
    //     request({
    //       url: "file/d/" + encodePath(file),
    //       method: "get",
    //     }).then((res) => {
    //       if (file.ext.toLowerCase() === "md") {
    //         text.value = res;
    //       } else {
    //         text.value = "```" + file.ext.toLowerCase() + "\n" + res + "\n```";
    //       }
    //       previewSpinning.value = false;
    //     });
    //     console.log(show.value);
    //     return;
    //   }
  } else if (category === "image") {
    show.value.image = true;
    show.value.spinning = true;
    previewSpinning.value = true;
    return;
  } else if (category === "audio") {
    show.value.audio = true;
    const audioOptions = {
      container: document.getElementById("audio-preview"),
      autoplay: false,
      audio: [
        {
          // 音频名称
          name: file.name,
          // 音频 URL
          url: downloadLink.value,
          // 音频照片
          cover: siteStore.siteInfo.music_img,
        },
        ...siteStore.siteInfo.audios,
      ],
      listMaxHeight: "65vh",
    };
    ap = new APlayer(audioOptions);
    return;
  } else if (
    category === "video" ||
    file.ext === "m3u8" ||
    file.ext === "flv" ||
    file.ext === "mpd"
  ) {
    let pluginOptions, customTypes, videoType;
    if (file.ext === "m3u8" || file.ext === "flv" || file.ext === "mpd") {
      pluginOptions = [
        artplayerPluginHlsQuality({
          // Show quality in control
          control: true,
          // Show quality in setting
          setting: true,
          // Get the resolution text from level
          getResolution: (level) => level.height + "P",
          // I18n
          title: "视频质量",
          auto: "自动",
        }),
      ];
    }
    if (file.ext === "m3u8") {
      customTypes = {
        m3u8: playM3u8,
      };
      videoType = "m3u8";
    } else if (file.ext === "flv") {
      customTypes = {
        flv: playFlv,
      };
      videoType = "flv";
    } else if (file.ext === "mpd") {
      customTypes = {
        mpd: playMpd,
      };
      videoType = "mpd";
    } else {
      customTypes = {};
      videoType = file.type;
    }
    show.value.video = true;
    art = new Artplayer({
      container: artRef.value,
      url: downloadLink.value,
      title: file.name,
      autoplay: !!siteStore.siteInfo.autoplay,
      autoSize: true,
      autoMini: true,
      playbackRate: true,
      aspectRatio: true,
      setting: true,
      hotkey: true,
      pip: true,
      fullscreen: true,
      fullscreenWeb: true,
      playsInline: true,
      lock: true,
      fastForward: true,
      autoOrientation: true,
      type: videoType,
      i18n: {},
      // pluginOptions: pluginOptions,
      // customTypes: customTypes,
    });
    return;
  } else {
    show.value.other = true;
  }
};
function playM3u8(video, url, art) {
  if (Hls.isSupported()) {
    const hls = new Hls();
    hls.loadSource(url);
    hls.attachMedia(video);

    // optional
    art.hls = hls;
    art.once("url", () => hls.destroy());
    art.once("destroy", () => hls.destroy());
  } else if (video.canPlayType("application/vnd.apple.mpegurl")) {
    video.src = url;
  } else {
    art.notice.show = "Unsupported playback format: m3u8";
  }
}

function playFlv(video, url, art) {
  if (flvjs.isSupported()) {
    const flv = flvjs.createPlayer({ type: "flv", url });
    flv.attachMediaElement(video);
    flv.load();

    // optional
    art.flv = flv;
    art.once("url", () => flv.destroy());
    art.once("destroy", () => flv.destroy());
  } else {
    art.notice.show = "Unsupported playback format: flv";
  }
}

function playMpd(video, url, art) {
  if (dashjs.supportsMediaSource()) {
    const dash = dashjs.MediaPlayer().create();
    dash.initialize(video, url, art.option.autoplay);

    // optional
    art.dash = dash;
    art.once("url", () => dash.destroy());
    art.once("destroy", () => dash.destroy());
  } else {
    art.notice.show = "Unsupported playback format: mpd";
  }
}

onMounted(() => {
  if (isNotEmpty(fileStore.file)) {
    preview(fileStore.file);
    file.value = setIcon(fileStore.file);
  } else {
    setTimeout(() => {
      preview(fileStore.file);
      file.value = setIcon(fileStore.file);
    }, 500);
  }
});

onBeforeUnmount(() => {
  if (ap) {
    ap.destroy();
  }
  if (art) {
    art.destroy();
  }
});
</script>
<style lang="scss" scoped>
.video-preview {
  width: 100%;
  margin-top: 5px;
}
@media screen and (min-width: 600px) {
  .video {
    height: 80vh;
  }
}
.iframe-preview {
  width: 100%;
  height: 80vh;
  box-sizing: inherit;
}
.doc-preview {
  width: 100%;
  height: 80vh;
}
.img-preview {
  width: 100%;
  height: 80vh;
}
.img-preview img {
  max-width: 100%;
  max-height: 100%;
  display: block;
  margin: auto;
}
</style>