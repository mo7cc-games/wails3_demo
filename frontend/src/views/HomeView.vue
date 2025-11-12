<template>
  <div class="HomeView">
    <h1>{{ message }}</h1>
  </div>
</template>

<script lang="ts">
export default {
  name: 'HomeView',
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
.HomeView {
  padding: 15px;
  border: 1px solid #000;
}
</style>
