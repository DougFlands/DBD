<template>
  <div class="page">
    <el-button type="primary" @click="jump('dbd')" id="btn-dbd">夺宝岛</el-button>
  </div>
</template>

<script>
export default {
  name: 'home',
  data() {
    return {
      
    }
  },
  methods: {
    jump(name) {
      this.$router.push({
        name
      })
    }
  },
};
</script>

<style lang="scss" scoped>
.page {
  display: flex;
  flex-direction: column;
}
</style>

