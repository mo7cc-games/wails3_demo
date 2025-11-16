<template>
  <TitleBar v-if="IsShow_TitleBar()" />
  <div class="page">
    <RouterView />
  </div>
</template>

<script lang="ts">
import { RouterView } from 'vue-router';
import { useWailsDataStore } from '@src/stores/WailsData';
import { mapState } from 'pinia';

import TitleBar from '@src/components/TitleBar.vue';

export default {
  name: 'App',
  components: {
    TitleBar,
    RouterView,
  },
  data() {
    return {};
  },
  computed: {
    ...mapState(useWailsDataStore, ['WindowName']),
  },

  mounted() {},
  methods: {
    IsShow_TitleBar() {
      if (!this.WindowName) {
        return false;
      }
      if (this.WindowName === 'Ball') {
        return false;
      }
      return true;
    },
  },
};
</script>

<style lang="scss" scoped>
.page {
  height: 100%;
  width: 100%;
}
</style>
