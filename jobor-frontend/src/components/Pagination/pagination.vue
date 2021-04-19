<template>
  <div style="width: 100%">
    <!--region 分页-->
    <el-pagination
      :small="small"
      :page-size.sync="pageSize"
      :current-page.sync="currentPage"
      :pager-count="pagerCount"
      :page-sizes="pageSizes"
      :layout="layout"
      :total="total"
      style="margin: 5px 0"
      @size-change="handleSizeChange"
      @current-change="handleIndexChange"
    />
    <!--endregion-->
  </div>

</template>

<script>
export default {
  name: 'Pagination',
  props: {
    // eslint-disable-next-line vue/require-prop-types
    small: { default() { return false } },
    // eslint-disable-next-line vue/require-prop-types
    pageSizes: {
      require: true,
      default() {
        return [5, 10, 25, 50, 100, 200]
      } },
    // pageSize:{default: 10},
    // currentPage:{default() { return 1}},
    // eslint-disable-next-line vue/require-prop-types
    limit: { default: 10 },
    // eslint-disable-next-line vue/require-prop-types
    page: { default() { return 1 } },
    // eslint-disable-next-line vue/require-prop-types
    pagerCount: { default: 7 },
    // eslint-disable-next-line vue/require-prop-types
    total: { default: 0 },
    layout: {
      type: String,
      default: 'prev, pager, next, jumper,sizes,total'
    },
    background: {
      type: Boolean,
      default: true
    }
  },
  data() {
    return {
      // tableData: [],
      // currentPage: 0,
      // total: 0,
    }
  },
  computed: {
    currentPage: {
      get() {
        return this.page
      },
      set(val) {
        this.$emit('update:page', val)
      }
    },
    pageSize: {
      get() {
        return this.limit
      },
      set(val) {
        this.$emit('update:limit', val)
      }
    }
  },
  created: function() {
    // this.currentPage=1;
  },
  methods: {
    handleSizeChange: function(val) {
      // console.log("handleSizeChange:", val, this.page, this.currentPage)
      this.$emit('pagination', { page: this.currentPage, limit: val })
    },
    handleIndexChange: function(val) {
      // console.log("handleIndexChange:", val, this.pageSize)
      this.$emit('pagination', { page: val, limit: this.pageSize })
    }
  }
}

</script>

<style scoped>

</style>
