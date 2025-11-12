<template>
  <div class="AboutView">
    <h1>{{ message }}</h1>
  </div>
</template>

<script lang="ts">
export default {
  name: 'AboutView',
  components: {},
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
.AboutView {
  padding: 15px;
  border: 1px solid #000;
}
</style>
