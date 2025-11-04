<script setup lang="ts">
import HelloWorld from './components/HelloWorld.vue';
</script>

<template>
  <div class="container">
    <div>
      <a data-wml-openURL="https://wails.io">
        <img src="/wails.png" class="logo" alt="Wails logo" />
      </a>
      <a data-wml-openURL="https://vuejs.org/">
        <img src="/vue.svg" class="logo vue" alt="Vue logo" />
      </a>
    </div>
    <HelloWorld msg="Wails + Vue" />
  </div>
</template>

<style scoped></style>
