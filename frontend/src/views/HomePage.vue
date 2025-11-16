<template>
  <div class="HomePage">
    <h1>{{ title }}</h1>
  </div>
</template>

<script lang="ts">
export default {
  name: 'HomePage',
  data() {
    return {
      title: '',
    };
  },
  mounted() {
    this.title = `Hello from ${this.$options.name} ! `;
  },
};
</script>

<style scoped lang="scss">
.HomePage {
  padding: 15px;
  border: 1px solid #000;
  background-color: rgba($color: #fff, $alpha: 0.7);
}
</style>
