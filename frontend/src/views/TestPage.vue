<template>
  <div class="TestPage">
    <SimpleBar class="SimpleBar">
      <h1>{{ title }} {{ TimeStr }}</h1>
      <div class="hint">若此处没时间显示，说明与后端通信有问题！</div>
      <div class="input-box">
        <input class="n-input" v-model="math.n1" type="text" autocomplete="off" />
        <input class="n-input" v-model="math.n2" type="text" autocomplete="off" />
        <button class="btn" @click="doAdd">Add</button>
        <div class="result">
          Result： {{ math.result }} 。<span class="hint">如果计算的值不正确则表示有问题</span>
        </div>
      </div>

      <div>
        计数：{{ count }} 双倍计数：{{ doubleCount }}
        <br />
        <button @click="OpenTestWindow">OpenTestWindow</button>
        <span class="hint">点击下方 +1 按钮所有窗口都会通知 Go 并同步计数</span>
        <br />
      </div>

      <CounterOptions />

      <div @contextmenu.prevent="showMenu($event)">
        <p>前端页面级右键菜单 (右键点我)</p>
        <ul
          v-if="menuInfo.visible"
          class="context-menu"
          :style="{ top: menuInfo.y + 'px', left: menuInfo.x + 'px' }"
        >
          <li @click="refresh">刷新页面</li>
          <li @click="quit">退出应用</li>
        </ul>
      </div>

      <!-- 在 go 中定义的 -->
      <div style="--custom-contextmenu: editor-menu; --custom-contextmenu-data: '墨七的数据'">
        自定义右键菜单 (右键点我)
      </div>

      <!-- Parent sets hide .env PRODUCTION=true 的时候生效 -->
      <div style="--default-contextmenu: hide">
        <!-- This inherits hide -->
        <p>没有右键菜单</p>
        <!-- This overrides to show -->
        <div style="--default-contextmenu: show">
          <p>显示默认的右键菜单</p>
          <!-- This inherits show -->
          <span>继承显示默认菜单</span>
          <!-- This resets to automatic behaviour -->
          <div style="--default-contextmenu: auto">
            <p>重置为自动行为模式</p>
          </div>
        </div>
      </div>

      <button @click="OpenBallWindow">打开悬浮球</button>
    </SimpleBar>
  </div>
</template>

<script lang="ts">
import CounterOptions from '@src/stores/CounterOptions.vue';
import { useCounterOptionsStore } from '@src/stores/counterOptions';
import { useWailsDataStore } from '@src/stores/WailsData';
import { mapState } from 'pinia';
import { WailsServe, OpenTestWindow, OpenBallWindow } from '@src/utils/wails';
import SimpleBar from 'simplebar-vue';

export default {
  name: 'TestPage',
  components: {
    CounterOptions,
    SimpleBar,
  },
  data() {
    return {
      title: '',
      math: {
        n1: '',
        n2: '',
        result: '',
      },
      menuInfo: {
        visible: false,
        x: 0,
        y: 0,
      },
    };
  },
  computed: {
    ...mapState(useCounterOptionsStore, ['count', 'doubleCount']),
    ...mapState(useWailsDataStore, ['TimeStr']),
  },
  mounted() {
    this.title = `This is ${this.$options.name} ! `;
    // 点击其他地方时关闭菜单
    window.addEventListener('click', this.hideMenu);
  },
  beforeUnmount() {
    window.removeEventListener('click', this.hideMenu);
  },
  methods: {
    OpenTestWindow,
    OpenBallWindow,
    doAdd() {
      const _this = this;
      const n1 = Number(_this.math.n1);
      const n2 = Number(_this.math.n2);
      WailsServe.Add(n1, n2).then((val: string) => {
        _this.math.result = val;
      });
    },
    showMenu(e: MouseEvent) {
      this.menuInfo.x = e.clientX; // 相对于整个文档左上角
      this.menuInfo.y = e.clientY;
      this.menuInfo.visible = true;
    },
    refresh() {
      console.info('刷新逻辑');
      this.menuInfo.visible = false;
    },
    quit() {
      console.info('退出逻辑');
      this.menuInfo.visible = false;
    },
    hideMenu() {
      this.menuInfo.visible = false;
    },
  },
};
</script>

<style scoped lang="scss">
.TestPage {
  border: 1px solid #000;
  padding: 2px;
  border-radius: 10px;
  box-sizing: border-box;
  height: 100%;
  background-color: rgba($color: #fff, $alpha: 0.8);

  .SimpleBar {
    padding: 15px;
    padding-bottom: 40px;
    box-sizing: border-box;
    height: 100%;
  }
}

.hint {
  color: #888;
  text-shadow: 1px 1px 2px #ccc;
  font-size: 12px;
}

.n-input {
  width: 80px;
  margin-right: 10px;
  padding: 4px 8px;
  font-size: 14px;
}

.context-menu {
  position: fixed;
  background: #fff;
  border: 1px solid #ccc;
  list-style: none;
  padding: 4px 0;
  margin: 0;
  z-index: 9999;
}
.context-menu li {
  padding: 6px 12px;
  cursor: pointer;
}
.context-menu li:hover {
  background: #eee;
}
</style>
