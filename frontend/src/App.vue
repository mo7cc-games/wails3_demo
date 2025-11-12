<template>
  <TitleBar />
  <div class="App">
    <div class="left">
      <div class="info">
        <div class="name">{{ WindowName }} 窗口完整地址：</div>
        <div class="url">
          {{ fullUrl }}
        </div>
      </div>
      <nav class="nav">
        <RouterLink class="nav-item" to="/">Home</RouterLink>
        <RouterLink class="nav-item" to="/about">About</RouterLink>
        <RouterLink class="nav-item" to="/test">Test</RouterLink>
      </nav>
    </div>
    <div class="right">
      <img alt="AppLogo" class="AppLogo" src="@src/assets/appicon.png" draggable="false" />
    </div>
  </div>
  <div class="router-view">
    <RouterView />
  </div>
</template>

<script lang="ts">
import { RouterLink, RouterView } from 'vue-router';
import TitleBar from '@src/components/TitleBar.vue';
import { useWailsDataStore } from '@src/stores/WailsData';
import { mapState } from 'pinia';
export default {
  name: 'App',
  components: {
    TitleBar,
    RouterLink,
    RouterView,
  },
  data() {
    return {
      fullUrl: window.location.href,
    };
  },
  computed: {
    ...mapState(useWailsDataStore, ['WindowName']),
  },
  watch: {
    $route() {
      this.fullUrl = window.location.href;
    },
  },
  mounted() {},
  methods: {
    getFullUrl() {},
  },
};
</script>

<style lang="scss" scoped>
@use '@src/assets/global.scss';

.router-view {
  background-color: rgba($color: #fff, $alpha: 0.8);
}

.info {
  width: 370px;
  text-align: center;
  .name {
    font-weight: bold;
  }
  .url {
    color: #888;
    word-break: break-all;
  }
}

.App {
  background-color: rgba($color: #ffff, $alpha: 0.7);
  width: 500px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  .AppLogo {
    width: 125px;
    @extend .wails-drag;
  }
  .nav {
    display: flex;
    align-items: center;
    justify-content: center;
    .nav-item {
      margin: 0 10px;
      background-color: transparent;
      padding: 6px 10px;
      border-radius: 4px;
      cursor: pointer;
      @extend .no-select;
    }
    .router-link-exact-active {
      // 精确匹配（只有当前路由和链接路径完全一致时才会加上）
      background-color: aquamarine;
      padding: 6px 10px;
      border-radius: 4px;
    }
    .router-link-active {
      // 模糊匹配（当前路由包含该链接的路径时就会加上）
      @extend .router-link-exact-active;
    }
  }
}
</style>
