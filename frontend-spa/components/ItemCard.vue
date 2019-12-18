<template>
  <div class="itemcard">
    <h1>{{ item.name }}</h1>
    <h2>{{ item.description }}</h2>
    <h2>{{ item.amount }}円</h2>
  </div>
</template>

<script>
  export default {
    name: 'ItemCard',
    props: [
      'item'
    ]
  }
</script>

<style scoped>
  /* css は vue init で最初に作られていた Helloworld.vue の css も追加 */

  .itemcard {
    border: solid 1px gray;
  }
</style>
