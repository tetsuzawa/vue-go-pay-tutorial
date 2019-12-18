<template>
  <div class="hello">
    <h1>{{ item.name }}</h1>
    <h2>{{ item.description }}</h2>
    <h2>{{ item.amount }}円</h2>

    <!--api-key="<< PAY.JPの管理画面にある公開テストKey >>"-->
    <!--    api-key={{ check_pub_key }}-->
    <!--      v-bind:api-key="check_pub_key"-->
    <payjp-checkout
      api-key="pk_test_38e208de8e4dfcebc9bf863c"
      text="カードを情報を入力して購入"
      submit-text="購入確定"
      name-placeholder="田中 太郎"
      v-on:created="onTokenCreated"
      v-on:failed="onTokenFailed">
    </payjp-checkout>

    <p>{{ message }}</p>
<!--    <p>id: {{ id }}</p>-->
    <router-link to="/">HOMEへ</router-link>
  </div>
</template>

<script>
  import axios from 'axios'

  export default {
    name: 'ItemCard',
    data() {
      return {
        item: {},
        message: '',
        check_pub_key: process.env.CHECK_PUB_KEY,
      }
    },
    created() {
      // urlで指定された動的パラメーターから商品情報をとってくる。
      axios.get(`http://localhost:8888/api/v1/items/${this.$route.params.id}`).then(res => {
        this.item = res.data
      })
    },
    beforeDestroy() {
      window.PayjpCheckout = null
    },
    methods: {
      // カードのToken化に成功したら呼ばれる。そのTokenでそのまま商品購入にうつる。
      onTokenCreated: function(res) {
        console.log(res.id)
        const data = { Token: res.id }
        axios.post(`http://localhost:8888/api/v1/charge/items/${this.$route.params.id}`, data).then(res => {
          this.message = '商品の購入が完了しました！'
        })
      },
      // Token化に失敗したら呼ばれる。
      onTokenFailed: function(status, err) {
        console.log(status)
        console.log(err)
      }
    }
  }
</script>
