<template>
    <el-button type="text" :size="size" style="color: #00c752" @click="handleClick">{{title}}</el-button>
</template>

<script>
    export default {
      name: "green",
      props:{
        size: {default: "mini"},
        title: {default: "title"},
      },
      methods:{
        handleClick(evt) {
          this.$emit('click', evt);
        }
      },
    }
</script>

<style scoped>

</style>
