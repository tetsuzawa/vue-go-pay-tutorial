<template>
  <div class="hello">
    <ul>
      <li v-for="item in items" :key="item.id" @click="pageto(item.id)">
        <item-card :item="item"></item-card>
      </li>
    </ul>
  </div>
</template>

<script>
  import axios from 'axios'
  import ItemCard from '@/components/ItemCard'

  export default {
    name: 'Home',
    components: {
      ItemCard
    },
    data() {
      return {
        items: []
      }
    },
    methods: {
      // ページ移動
      pageto: function(id) {
        this.$router.push(`/items/${id}`)
      }
      //  TODO ページ遷移
    },
    // 商品リストをすべてとってくる
    created() {
      axios.get('http://localhost:8888/api/v1/items').then(res => {
        this.items = res.data
      })
    }
  }
</script>


<!--
<template>
  <div class="container">
    <div>
      <logo />
      <h1 class="title">
        frontend-spa
      </h1>
      <h2 class="subtitle">
        My transcendent Nuxt.js project
      </h2>
      <div class="links">
        <a
          href="https://nuxtjs.org/"
          target="_blank"
          class="button--green"
        >
          Documentation
        </a>
        <a
          href="https://github.com/nuxt/nuxt.js"
          target="_blank"
          class="button--grey"
        >
          GitHub
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import Logo from '~/components/Logo.vue'

export default {
  components: {
    Logo
  }
}
</script>
-->

<style>
  .container {
    margin: 0 auto;
    min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
  }

  .hello {
    font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
    display: block;
    font-weight: 300;
    font-size: 60px;
    color: #35495e;
    letter-spacing: 1px;
  }


  .title {
    font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
    display: block;
    font-weight: 300;
    font-size: 100px;
    color: #35495e;
    letter-spacing: 1px;
  }

  .subtitle {
    font-weight: 300;
    font-size: 42px;
    color: #526488;
    word-spacing: 5px;
    padding-bottom: 15px;
  }

  .links {
    padding-top: 15px;
  }
</style>
