<template>
  <div class="TitleBar">
    <div v-if="!isHide">
      <div class="SysInfo">
        <pre
          >{{ JSON.stringify(ScreenInfo, null, 2) }}
WebGPU: {{ IsWebGPU }}</pre
        >
      </div>

      <button class="btn" @click="WindowControl('mini')">最小化</button>

      <button class="btn" v-if="!ScreenInfo.IsMaximise" @click="WindowControl('max')">
        最大化
      </button>
      <button class="btn" v-if="ScreenInfo.IsMaximise" @click="WindowControl('resize')">
        还原
      </button>

      <button class="btn" v-if="!ScreenInfo.IsFullscreen" @click="WindowControl('fullscreen')">
        全屏
      </button>
      <button class="btn" v-if="ScreenInfo.IsFullscreen" @click="WindowControl('unfullscreen')">
        退出全屏
      </button>
      <button class="btn" v-if="!IsFrameless" @click="SetFrameless">无边框</button>
      <button class="btn" v-if="IsFrameless" @click="OutFrameless">退出无边框</button>
      <button class="btn" @click="WindowControl('close')">关闭窗口</button>
    </div>
    <button class="btn" @click="isHide = !isHide">{{ isHide ? '显示菜单' : '隐藏菜单' }}</button>
  </div>
</template>

<script lang="ts">
import wailsio from '@src/utils/wailsio';
import { useWailsDataStore } from '@src/stores/WailsData';
import { mapState, mapActions } from 'pinia';

interface data {
  isHide: boolean;
}

export default {
  name: 'TitleBar',
  data(): data {
    return {
      isHide: true,
    };
  },
  computed: {
    ...mapState(useWailsDataStore, ['IsFrameless', 'WindowName', 'IsWebGPU', 'ScreenInfo']),
  },
  async mounted() {},
  beforeUnmount() {},
  methods: {
    ...mapActions(useWailsDataStore, ['SetFrameless', 'OutFrameless']),
    async WindowControl(type: string) {
      if (type === 'mini') {
        await wailsio.Window.Minimise();
      }
      if (type === 'max') {
        await wailsio.Window.Maximise();
      }
      if (type === 'resize') {
        await wailsio.Window.Restore();
      }
      if (type === 'toggle') {
        await wailsio.Window.ToggleMaximise();
      }
      if (type === 'close') {
        await wailsio.Window.Close();
      }
      if (type === 'fullscreen') {
        await wailsio.Window.Fullscreen();
      }
      if (type === 'unfullscreen') {
        await wailsio.Window.UnFullscreen();
      }
    },
  },
};
</script>

<style scoped lang="scss">
@use '@src/assets/global.scss';
.TitleBar {
  padding: 15px;
  border: 1px solid #000;
  background-color: rgba($color: #f46583, $alpha: 0.6);
  color: #fff;
  position: fixed;
  z-index: 999;
  left: 0;
  top: 0;
  @extend .wails-drag;
  cursor: move;
}

.SysInfo {
  background-color: #000;
}

.btn {
  @extend .wails-no-drag;
  cursor: pointer;
}
</style>
