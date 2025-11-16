<template>
  <div class="FloatingBall" :class="BallWindow.ZoomDir">
    <div
      class="Ball"
      @click="OnClickFunc"
      style="--custom-contextmenu: ball-menu"
      :style="{ width: `${BallWindow.OrgW}px`, height: `${BallWindow.OrgH}px` }"
    >
      <img alt="AppLogo" class="img" src="/img/appicon.png" draggable="false" />
    </div>
    <div class="menu" v-if="IsMenu">这里是一些文字~</div>
  </div>
</template>

<script lang="ts">
import { useWailsDataStore } from '@src/stores/WailsData';
import { WailsServe } from '@src/utils/wails';
import { mapState } from 'pinia';

export default {
  name: 'FloatingBall',
  devtools: false,
  components: {},
  data() {
    return {
      IsMenu: false,
      ballSize: '',
    };
  },
  computed: {
    ...mapState(useWailsDataStore, ['WindowName', 'BallWindow', 'ScreenInfo']),
  },
  watch: {},
  mounted() {},
  methods: {
    OnClickFunc() {
      this.SwitchWindowSize();
    },
    SwitchWindowSize() {
      const _this = this;
      if (_this.WindowName !== 'Ball') {
        return;
      }
      if (_this.BallWindow.ZoomDir === 'Reset') {
        WailsServe.BallWindowZoomIn().then(() => {
          _this.IsMenu = true; // 展开后再打开菜单
        });
      } else {
        _this.IsMenu = false; // 先关闭菜单再收起
        WailsServe.BallWindowReset();
      }
    },
  },
};
</script>

<style scoped lang="scss">
@use '@src/assets/global.scss';

$width: 100px;

.FloatingBall {
  width: 100%;
  height: 100%;
  position: relative;
  box-sizing: border-box;
}

.menu {
  position: absolute;
  width: 90%;
  height: 90%;
  left: 0;
  top: 0;
  border-radius: 20px;
  background-color: rgba($color: #fff, $alpha: 0.7);
  z-index: 1;
}

.Ball {
  position: absolute;
  padding: 5px;
  z-index: 2;
  left: 0;
  top: 0;
  border-radius: 50%;
  overflow: hidden;
  box-sizing: border-box;
  @include global.wails-drag;
  .img {
    width: 100%;
    height: 100%;
    display: block;
    border-radius: 50%;
    overflow: hidden;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  }
}

.RightDown {
  .Ball,
  .menu {
    right: auto;
    bottom: auto;
    left: auto;
    top: auto;
  }
  .Ball {
    top: 0;
    left: 0;
  }
  .menu {
    right: 0;
    bottom: 0;
  }
}

.RightUp {
  .Ball,
  .menu {
    right: auto;
    bottom: auto;
    left: auto;
    top: auto;
  }
  .Ball {
    left: 0;
    bottom: 0;
  }
  .menu {
    right: 0;
    top: 0;
  }
}

.LeftDown {
  .Ball,
  .menu {
    right: auto;
    bottom: auto;
    left: auto;
    top: auto;
  }
  .Ball {
    right: 0;
    top: 0;
  }
  .menu {
    left: 0;
    bottom: 0;
  }
}
.LeftUp {
  .Ball,
  .menu {
    right: auto;
    bottom: auto;
    left: auto;
    top: auto;
  }
  .Ball {
    right: 0;
    bottom: 0;
  }
  .menu {
    left: 0;
    top: 0;
  }
}
</style>

<style>
#__vue-devtools-container__ {
  display: none !important;
}
</style>
