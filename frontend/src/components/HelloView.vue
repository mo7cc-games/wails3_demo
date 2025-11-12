<template>
  <div class="HelloView">
    <h1>{{ message }}</h1>
  </div>
</template>

<script lang="ts">
export default {
  name: 'HelloView',
  data() {
    return {
      message: '',
    };
  },
  mounted() {
    this.message = `Hello from ${this.$options.name} ! `;
  },
};
</script>

<style scoped lang="scss">
.HelloView {
  padding: 15px;
  border: 1px solid #000;
}
</style>
