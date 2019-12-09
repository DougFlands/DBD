<template>
  <div class="home">
    <div style="margin-bottom: 20px">
      <el-input v-model="input" placeholder="请输入正则" style="width: 300px;margin-right: 20px"></el-input>
      <el-button @click="get(1)" type="primary">查找</el-button>
    </div>

    <el-table :data="list" style="width: 100%" border>
      <el-table-column prop="brandName" width="220" label="品牌"></el-table-column>
      <el-table-column prop="productName" label="名称"></el-table-column>
      <el-table-column prop="quality" label="新旧" width="120"></el-table-column>
    </el-table>
    <el-pagination layout="prev, pager, next" page-size="50" :total="total" @current-change="get"></el-pagination>
  </div>
</template>

<script>
export default {
  name: 'home',
  data() {
    return {
      list: [],
      pageNum: 1,
      total: 1,
      loading: false,
      input: '',
    }
  },
  methods: {
    get(page) {
      this.pageNum = page
      this.loading = true
      this.list = []
      this.$api.common.getList(this.pageNum).then(res => {
        this.loading = false
        this.list = res.auctionInfos.filter(val => {
          let re = new RegExp(this.input)
          return re.test(val.productName)
        })
        this.total = res.totalNumber
      }).catch(err => {
        this.loading = false
        console.log(err);
      })
    }
  },
  created() {
    this.get(1)
  }
};
</script>
