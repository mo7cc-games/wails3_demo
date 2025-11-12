import { defineStore } from 'pinia';
import { WailsService } from '@server/app';

export const useCounterOptionsStore = defineStore('counterOptions', {
  // 状态：存储数据
  state: () => ({
    count: 0,
    name: '墨七',
  }),

  // Getters：类似计算属性，用于派生状态
  getters: {
    doubleCount(state) {
      return state.count * 2;
    },
    greeting(state) {
      return `Hello, ${state.name}! 计数是 ${state.count}`;
    },
  },
  // Actions ：类似 methods，用于修改状态（支持同步和异步）
  actions: {
    increment() {
      // 专门用于事件通信 会 触发 go 端的 Events.On('Action',()=>{})
      WailsService.Action('counterOptions.add');
    },
    add() {
      this.count++; // 使用 `this` 访问和修改状态
    },
    reset() {
      this.count = 0;
    },
    async incrementAsync(delay: number) {
      await new Promise((resolve) => {
        setTimeout(resolve, delay);
      });
      this.increment(); // 在 action 中调用其他 action
    },
  },
});
