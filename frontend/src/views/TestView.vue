<template>
  <div class="TestView">
    <h1>{{ message }} {{ TimeStr }}</h1>
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
  </div>
</template>

<script lang="ts">
import CounterOptions from '@src/stores/CounterOptions.vue';
import { useCounterOptionsStore } from '@src/stores/counterOptions';
import { useWailsDataStore } from '@src/stores/WailsData';
import { mapState } from 'pinia';
import { WailsService } from '@server/app';
import { throttle } from 'radash';

// 节流函数，防止短时间内多次点击打开多个窗口
const OpenTestWindowFunc = throttle({ interval: 1000 }, () => {
  WailsService.OpenTestWindow();
});

export default {
  name: 'TestView',
  components: {
    CounterOptions,
  },
  data() {
    return {
      message: '',
      math: {
        n1: '',
        n2: '',
        result: '',
      },
    };
  },
  computed: {
    ...mapState(useCounterOptionsStore, ['count', 'doubleCount']),
    ...mapState(useWailsDataStore, ['TimeStr']),
  },
  mounted() {
    this.message = `This is ${this.$options.name} ! `;
  },
  methods: {
    doAdd() {
      const _this = this;
      const n1 = Number(_this.math.n1);
      const n2 = Number(_this.math.n2);
      WailsService.Add(n1, n2).then((val: string) => {
        _this.math.result = val;
      });
    },
    OpenTestWindow() {
      OpenTestWindowFunc();
    },
    ActionFunc() {},
  },
};
</script>

<style scoped lang="scss">
.TestView {
  padding: 15px;
  border: 4px solid #000;
  padding-bottom: 40px;
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
</style>
