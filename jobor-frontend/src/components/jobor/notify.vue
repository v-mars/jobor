<template>
  <ul style="list-style-type:none;">
    <li>
      <el-checkbox v-model="Notify.webhook['enable']">Webhook</el-checkbox>
      <el-button size="mini" type="text" class="add-btn"
                 @click="Notify.webhook.urls.push('')">添加</el-button>
      <div v-for="(item,index) in Notify.webhook.urls" :key="index" style="margin: 3px 0">
        <el-input v-model="Notify.webhook.urls[index]" aria-valuemax="512"
                  placeholder="请输入api地址" style="max-width: 360px">
          <template slot="prepend">api地址</template>
        </el-input>
        <el-button :disabled="false" size="mini" @click="Notify.webhook.urls.splice(index, 1)"
                   icon="el-icon-delete" type="danger" circle class="add-btn"></el-button>
      </div>

    </li>
    <li>
      <el-checkbox v-model="Notify.email.enable">邮件</el-checkbox>
      <el-button size="mini" type="text" class="add-btn"
                 @click="Notify.email.receivers.push('')">添加</el-button>
      <div v-for="(item,index) in Notify.email.receivers" :key="index" style="margin: 3px 0">
        <el-input v-model="Notify.email.receivers[index]" aria-valuemax="256"
                  placeholder="请输入邮箱地址" style="max-width: 360px">
          <template slot="prepend">邮箱地址</template>
        </el-input>
        <el-button :disabled="false" size="mini" @click="Notify.webhook.receivers.splice(index, 1)"
                   icon="el-icon-delete" type="danger" circle class="add-btn"></el-button>
      </div>

    </li>
    <li>
      <el-checkbox v-model="Notify.lark.enable">飞书</el-checkbox>
      <el-button size="mini" type="text" class="add-btn"
                 @click="Notify.lark.webhooks.push({})">添加</el-button>
      <div v-for="(item,index) in Notify.lark.webhooks" :key="index" style="margin: 3px 0">
        <el-input v-model="item.webhook_url" aria-valuemax="256"
                  placeholder="请输入飞书webhook url" style="max-width: 360px">
          <template slot="prepend">webhook url</template>
        </el-input>
        <el-input v-model="item.secret" aria-valuemax="256" type="password"
                  placeholder="请输入飞书webhook Secret" style="max-width: 240px;margin-left: 3px;">
          <template slot="prepend">Secret</template>
        </el-input>
        <el-button :disabled="false" size="mini" @click="Notify.lark.webhooks.splice(index, 1)"
                   icon="el-icon-delete" type="danger" circle class="add-btn"></el-button>
      </div>
    </li>
    <li>
      <el-checkbox v-model="Notify.dingding.enable">钉钉</el-checkbox>
      <el-button size="mini" type="text" class="add-btn"
                 @click="Notify.dingding.webhooks.push({})">添加</el-button>
      <div v-for="(item,index) in Notify.dingding.webhooks" :key="index" style="margin: 3px 0">
        <el-input v-model="item.webhook_url" aria-valuemax="256"
                  placeholder="请输入钉钉webhook url" style="max-width: 360px">
          <template slot="prepend">webhook url</template>
        </el-input>
        <el-input v-model="item.secret" aria-valuemax="256" type="password"
                  placeholder="请输入钉钉webhook Secret" style="max-width: 240px;margin-left: 3px;">
          <template slot="prepend">Secret</template>
        </el-input>
        <el-button :disabled="false" size="mini" @click="Notify.dingding.webhooks.splice(index, 1)"
                   icon="el-icon-delete" type="danger" circle class="add-btn"></el-button>
      </div>
    </li>
  </ul>
</template>

<script>
  import delete_button from '@/components/crud/delete_button'
  import edit_button from '@/components/crud/edit_button'
  import green_button from '@/components/crud/green_button'
    export default {
      name: "notify",
      props: {
        notify: {
          email: {enable:false,receivers:[]},
          webhook:{enable:false,urls:[]},
          lark:{enable:false,webhooks:[]},
          dingding:{enable:false,webhooks:[]}
        },
      },
      watch: {
        "notify": function () {
          this.Notify = this.notify
        },
        "Notify": function () {
          this.$emit("update:notify", this.Notify)
        },
      },
      data(){return{
        Notify: this.notify,
      }},

      components: {
        delete_button,
        edit_button,
        green_button,

      },
    }
</script>

<style lang="scss" scoped>
  .add-btn{
    margin-left: 5px;
  }
</style>
