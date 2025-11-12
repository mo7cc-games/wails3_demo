<template>
  <div class="CounterOptions">
    <h2>Pinia 选项式 API Demo</h2>
    <div>计数：{{ count }} ；双倍计数：{{ doubleCount }} ； {{ greeting }}</div>
    <button @click="increment">+1</button>
    <button @click="reset">重置</button>
    <button @click="incrementAsync(1000)">1 秒后 +1</button>
  </div>
</template>

<script lang="ts">
// 选项式 Api 使用 Demo
import { useCounterOptionsStore } from '@src/stores/counterOptions';
import { mapState, mapActions } from 'pinia';

export default {
  // 将 Store 的 state 和 getters 映射到组件的计算属性
  computed: {
    ...mapState(useCounterOptionsStore, ['count', 'doubleCount', 'greeting']),
  },
  // 将 Store 的 actions 映射到组件的方法
  methods: {
    ...mapActions(useCounterOptionsStore, ['increment', 'reset', 'incrementAsync']),
  },
};
</script>
<style scoped lang="scss">
.CounterOptions {
  padding: 15px;
  border: 1px solid #000;
}
</style>
