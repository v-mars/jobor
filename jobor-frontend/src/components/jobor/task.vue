<template>
  <div>
    <div style="margin: 5px 0">
      <!-- 搜索 -->
      <el-form :inline="true" :model="searchForm" size="small" class="demo-form-inline" @submit.native.prevent>
        <el-form-item label="">
          <el-input v-model="searchForm.name" placeholder="名称" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="getData">查询</el-button>
          <el-button type="primary" @click="newRow">新增</el-button>
          <el-button style="float: right" @click="getData">刷新</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div style="margin-top: 10px">
      <el-table border :data="data_list" size="small" v-loading="loading" key="jobor_task">
<!--        <el-table-column type="selection" width="45" align="center"></el-table-column>-->
        <el-table-column label="ID" prop="id" width="60"></el-table-column>
        <el-table-column label="名称" prop="name" width="">
          <template slot-scope="scope">
            <span>{{scope.row.name}}</span>
            <el-tooltip class="item" effect="dark"
                        placement="top-start" style="margin-left: 3px">
              <div v-html="'<pre style=\'margin: 1px\'>描述:\n'+scope.row.description+'</pre>'" slot="content"></div>
              <i class="el-icon-info" v-if="scope.row.description"></i>
            </el-tooltip>
          </template>
        </el-table-column>
        <el-table-column label="类型" prop="lang" width="80"></el-table-column>
        <el-table-column label="表达式" prop="expr" width=""></el-table-column>
<!--        <el-table-column label="任务脚本" prop="data" show-overflow-tooltip></el-table-column>-->
        <el-table-column label="routingKey" prop="routing_key"></el-table-column>
        <el-table-column label="状态" prop="" width="70" align="">
          <template slot-scope="scope">
            <el-popconfirm icon-color="red" icon="el-icon-info"
                           v-if="scope.row.status==='running'" :title="'确认停止任务吗？'"
                           @onConfirm="changeStatus(scope.row)">
              <el-switch v-model="scope.row.status==='running'" slot="reference"></el-switch>
            </el-popconfirm>
            <el-popconfirm icon="el-icon-info" v-else :title="'确认开始运行任务吗？'" @onConfirm="changeStatus(scope.row)">
              <el-switch v-model="scope.row.status==='running'" slot="reference"></el-switch>
            </el-popconfirm>
          </template>
        </el-table-column>
<!--        <el-table-column label="创建时间" prop="ctime" width="140"></el-table-column>-->
        <el-table-column label="上次执行时间" prop="prev" width="160"></el-table-column>
        <el-table-column label="操作" align="center" width="200">
          <template slot-scope="scope">
            <el-popconfirm icon="el-icon-info" :title="'确认开始手动执行任务吗？'" @onConfirm="runTask(scope.row)">
              <green_button title="运行" slot="reference"></green_button>
            </el-popconfirm>

<!--            <delete_button @click="changeStatus(scope.row,'stop')" v-if="scope.row.status==='running'" title="停止"></delete_button>-->
            <edit_button @click="showEdit(scope.row)" style="margin-left:10px"></edit_button>
            <delete_button @click="confirmDelRows(scope.row, scope.row.name, true)"></delete_button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 5px">
        <div style="display: inline"></div>
        <div class="block" style="float: right;display: inline">
          <pagination :total="total" :page.sync="page" :limit.sync="limit" @pagination="getData()"></pagination>
        </div>
      </div>
    </div>

    <el-dialog :title="title" :visible.sync="newOrUpdate" :close-on-click-modal="false" width="80%">
      <el-form label-width="120px" :model="rowData" size="small" ref="rowData" :rules="taskRules">
        <el-form-item label="名称" prop="name" :rules="[{required:true,message:'请输入名称', trigger: 'blur'}]">
          <el-input v-model="rowData.name" aria-valuemax="60"></el-input>
        </el-form-item>
        <el-form-item label="描述">
          <el-input type="textarea" :rows="2" placeholder="请输入描述内容" v-model="rowData.description"></el-input>
        </el-form-item>

        <el-form-item label="表达式" prop="expr">
          <el-input v-model="rowData.expr"></el-input>
          <span style="font-size: smaller;color: #c3c4c9">注：秒、分、时、天、月、周</span>
        </el-form-item>

        <el-form-item label="类型" prop="lang" :rules="[{required:true,message:'请输入代号', trigger: 'blur'}]">
          <el-select v-model="rowData.lang" @change="changeLang">
            <el-option value="shell" label="Shell"></el-option>
            <el-option value="api" label="Api"></el-option>
            <el-option value="python3" label="Python3"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="执行内容" prop="data" :rules="[{required:false,message:'请输入执行内容', trigger: 'blur'}]">
          <div v-if="rowData.lang === 'api'">
            <ul>
              <!--method url-->
              <div>
                <el-table size="mini" :data="[{}]" style="width: 100%" empty-text="please press add new">
                  <el-table-column label="Method" min-width="30px;">
                    <template slot-scope="scope">
                      <el-select :disabled="false" size="mini" v-model="rowData.data.api.method" placeholder="请选择">
                        <el-option v-for="item in methodOption" :key="item.label" :label="item.label" :value="item.value"
                        ></el-option>
                      </el-select>
                    </template>
                  </el-table-column>
                  <el-table-column label="URL" min-width="100px;">
                    <template slot-scope="scope">
                      <el-input :disabled="false" size="mini" v-model="rowData.data.api.url"></el-input>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
              <!--headers-->
              <div>
                <el-table size="mini" :data="rowData.data.api.header" style="width: 100%" empty-text=" ">
                  <el-table-column prop="name" label="Header头列表" min-width="100px;">
                    <template slot-scope="scope">
                      <el-input :disabled="false" size="mini" v-model="rowData.data.api.header[scope.$index].key"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column prop="age" min-width="100px;">
                    <template slot-scope="scope">
                      <el-input :disabled="false" size="mini" v-model="rowData.data.api.header[scope.$index].value"></el-input>
                    </template>
                  </el-table-column>
                  <el-table-column width="70px;">
                    <template slot-scope="scope">
                      <el-button :disabled="false" size="mini" @click="rowData.data.api.header.splice(scope.$index, 1)"
                        icon="el-icon-delete" type="danger" circle></el-button>
                    </template>
                  </el-table-column>
                </el-table>
                <div style="margin-left:11px">
                  <el-button :disabled="false" type="text" size="mini" @click="rowData.data.api.header.push({})">添加</el-button>
                </div>
              </div>
              <div v-if="['GET','HEAD'].includes(rowData.data.api.method) === false">
                <div>
                  <el-table size="mini" :data="[{}]" style="width: 100%" empty-text=" ">
                    <el-table-column prop="name" label="Content Type" min-width="100px;">
                      <template slot-scope="scope">
                        <el-select :disabled="false" size="mini" filterable v-model="rowData.data.api.content_type"
                                   style="width:100%;">
                          <el-option v-for="item in contentTypeOption" :key="item.label" :label="item.label" :value="item.value"
                          ></el-option>
                        </el-select>
                      </template>
                    </el-table-column>
                  </el-table>
                </div>
                <div v-if="rowData.data.api.content_type === 'application/x-www-form-urlencoded'">
                  <el-table size="mini" :data="rowData.data.api.forms" style="width: 100%" empty-text=" ">
                    <el-table-column prop="name" label="Form List" min-width="100px;">
                      <template slot-scope="scope">
                        <el-input :disabled="false" size="mini" v-model="rowData.data.api.forms[scope.$index].key"></el-input>
                      </template>
                    </el-table-column>
                    <el-table-column prop="age" min-width="100px;">
                      <template slot-scope="scope">
                        <el-input :disabled="false" size="mini" v-model="rowData.data.api.forms[scope.$index].value"></el-input>
                      </template>
                    </el-table-column>
                    <el-table-column width="70px;">
                      <template slot-scope="scope">
                        <el-button :disabled="false" size="mini"
                          @click="rowData.data.api.forms.splice(scope.$index, 1)"
                          icon="el-icon-delete" type="danger" circle></el-button>
                      </template>
                    </el-table-column>
                  </el-table>
                  <div style="margin-left:11px">
                    <el-button :disabled="false" type="text" size="mini" @click="rowData.data.api.forms.push({})">添加</el-button>
                  </div>
                </div>
                <div v-if="rowData.data.api.content_type === 'application/json'">
                  <el-table size="mini" :data="[{}]" style="width: 100%" empty-text=" ">
                    <el-table-column prop="name" label="Raw Requesy Body" min-width="100px;">
                      <!-- <div style="margin-top:5px;"> -->
                      <el-card :body-style="{ padding: '0px' }">
                        <editor
                          v-model="rowData.data.api.payload"
                          theme="solarized_dark"
                          lang="json"
                          height="300"
                          @init="rowReqBodyInitEditor"
                          :options="{
                            // readOnly: is_preview,
                            wrap: 'free',
                            indentedSoftWrap: false}"
                        ></editor>
                      </el-card>
                    </el-table-column>
                  </el-table>
                </div>
              </div>
            </ul>
          </div>
          <div style="margin-top:5px;" v-else>
            <el-card :body-style="{ padding: '0px' }">
              <editor
                v-model="rowData.data.data"
                theme="solarized_dark"
                :lang="lang[rowData.lang]"
                height="300"
                @init="codeInitEditor"
                :options="{
                      // readOnly: is_preview,
                      wrap: 'free',
                      indentedSoftWrap: false,
                      enableBasicAutocompletion: true,
                      enableLiveAutocompletion: true,
                      enableSnippets: false
                 }"
              ></editor>
            </el-card>
<!--            <span style="text-align: right;color: #909399;font-size: 12px;">双击编辑框全屏</span>-->
            <br />
          </div>

        </el-form-item>
        <el-form-item label="告警策略" prop="alarm_policy" :rules="[{required:true,message:'请输入', trigger: 'blur'}]">
          <el-select v-model="rowData.alarm_policy">
            <el-option :value="value" :label="alarmMapZH[value]" v-for="(value,key,index) in alarmMap" :key="index"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="超时时间" prop="timeout" :rules="[{required:true,message:'请输入', trigger: 'blur'}]">
          <el-input-number controls-position="right" :min="-1" :max="86400" v-model="rowData.timeout" style="width: auto"></el-input-number>
        </el-form-item>
<!--        <el-form-item label="失败重试" prop="retry" :rules="[{required:true,message:'请输入', trigger: 'blur'}]">-->
<!--          <el-input-number controls-position="right" :min="0" :max="20" v-model="rowData.retry" style="width: auto"></el-input-number>-->
<!--        </el-form-item>-->
        <el-form-item label="期望返回码" prop="expect_code" :rules="[{required:true,message:'请输入期望返回码', trigger: 'blur'}]">
          <el-input-number controls-position="right" v-model="rowData.expect_code" style="width: auto"></el-input-number>
        </el-form-item>
        <el-form-item label="期望返回内容">
          <el-input type="textarea" :rows="2" placeholder="请输入期望返回内容" v-model="rowData.expect_context"></el-input>
        </el-form-item>
        <el-form-item label="告警通知" prop="notify">
          <notify :notify.sync="rowData.notify"></notify>
        </el-form-item>

      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="newOrUpdate = false">取 消</el-button>
    <el-button :plain="true" type="warning" @click="createOrUpdate('PUT', false, rowData.id)" v-if="rowData.id" :loading="action_loading">更新</el-button>
    <el-button :plain="true" type="warning" @click="createOrUpdate('POST', true)" v-if="!rowData.id" :loading="action_loading">保存并继续</el-button>
    <el-button :plain="true" type="primary" @click="createOrUpdate('POST', false)" v-if="!rowData.id" :loading="action_loading">保 存</el-button>
  </span>
    </el-dialog>
  </div>

</template>

<script>
  import pagination from '@/components/Pagination/pagination'
  import delete_button from '@/components/crud/delete_button'
  import edit_button from '@/components/crud/edit_button'
  import green_button from '@/components/crud/green_button'
  import user_select from '@/components/sys/user_select'
  import notify from '@/components/jobor/notify'
  import common_mixin from '@/components/crud/common_mixin'
  import {validateUrl,EmailReCheck} from '@/utils/common'
  import { isValidCron } from 'cron-validator'
    export default {
      name: "jobor_task",
      mixins: [common_mixin],
      data(){
        const validateCronExpr = (rule, value, callback) => {
          if (!isValidCron(value, { seconds: true })) {
            callback(new Error('cron表达式不正确, 例：* * * * * *'))
          } else {
            callback()
          }
        }

        const validateNotify = (rule, value, callback) => {
          if(value.webhook && value.webhook.urls){
            for(let i=0;i<value.webhook.urls.length;i++){
              if(!validateUrl(value.webhook.urls[i])){
                callback(new Error('webhook url格式不正确, 例：http://api.xx.com/ping'))
              }
            }
          }
          if(value.email && value.email.receivers){
            for(let i=0;i<value.email.receivers.length;i++){
              if(!EmailReCheck(value.email.receivers[i])){
                callback(new Error('邮箱格式不正确, 例：zhangsan@email.com'))
              }
            }
          }
          if(value.lark && value.lark.webhooks){
            for(let i=0;i<value.lark.webhooks.length;i++){
              if(!validateUrl(value.lark.webhooks[i].webhook_url)){
                callback(new Error('飞书 webhook url格式不正确, 例：http://lark.xx.com/ping'))
              }
            }
          }
          if(value.lark && value.dingding.webhooks){
            for(let i=0;i<value.dingding.webhooks.length;i++){
              if(!validateUrl(value.dingding.webhooks[i].webhook_url)){
                callback(new Error('钉钉 webhook url格式不正确, 例：http://dingding.xx.com/ping'))
              }
            }
          }
          callback()
        }
        return{
          taskRules: {
            "expr": [{ required: true, trigger: 'blur', validator: validateCronExpr }],
            "notify": [{ required: false, trigger: 'blur', validator: validateNotify }]
          },

          notify: {
            email: {enable:false,receivers:[]},
            webhook:{enable:false,urls:[]},
            lark:{enable:false,webhooks:[]},
            dingding:{enable:false,webhooks:[]}
          },
          rowData: {
              name: "", lang: "shell",id: "",expr: "* * * * * *", timeout: -1,retry:0,
              expect_code: 0,expect_context: "",alarm_policy: 2, description: "",
              data:{"data": "",
                api:{
                  url: "",
                  method: "GET",
                  content_type: "",
                  payload: "",
                  header: [],
                  forms: [],
                }},
              notify: {
                email: {enable:false,receivers:[]},
                webhook:{enable:false,urls:[]},
                lark:{enable:false,webhooks:[]},
                dingding:{enable:false,webhooks:[]}
              },
          },
          url: this.$store.state.urls.jobor_task_url,

          alarmMap: {"never":0,"always":1,"failed":2,"success":3},
            // alarmMap: {0:"从不",1:"总是",2:"失败",3:"成功"},
          alarmMapZH: {0:"从不",1:"总是",2:"失败",3:"成功"},
          lang: {
            "shell": "sh",
            "python3": "python",
            "golang": "golang",
            "api": "sh",
            // 4: "python",
            // 5: "nodejs",
            // 6: "windowsbat"
          },
          apiRow: {
            url: "",
            method: "GET",
            content_type: "",
            payload: "",
            header: {},
            forms: [{}],
          },
          headerList:[{}],
          method: "",
          methodOption: [
            {
              value: "GET",
              label: "GET"
            },
            {
              value: "HEAD",
              label: "HEAD"
            },
            {
              value: "POST",
              label: "POST"
            },
            {
              value: "PUT",
              label: "PUT"
            },
            {
              value: "PATCH",
              label: "PATCH"
            },
            {
              value: "DELETE",
              label: "DELETE"
            }
          ],
          content_type: "",
          contentTypeOption: [
            {
              value: "application/json",
              label: "application/json"
            },
            {
              value: "application/x-www-form-urlencoded",
              label: "application/x-www-form-urlencoded"
            }
          ],
          langExample: {
            "shell": `#!/usr/bin/env sh
  function main() {
      echo "run shell"
  }

  main
          `,
            "python3": `#!/usr/bin/env python3
  def main():
      print("run python3")

  if __name__ == '__main__':
      main()`,
            "python": `#!/usr/bin/env python
  def main():
      print("run python")

  if __name__ == '__main__':
      main()`,
            "golang": `package main

  import "fmt"

  func main() {
    fmt.Println("run golang")
  }`,
            "nodejs": `#!/usr/bin/env node
  console.log("run nodejs")`,
            "windowsbat": `tasklist`,
            "api": "",
          },
        }},
      methods: {

        showEdit: function(row) {
          this.newOrUpdate = true
          this.title = '编辑'
          this.rowData.id = row.id
          this.rowData.name = row.name
          this.rowData.data = row.data
          this.rowData.notify = row.notify
          this.rowData.lang = row.lang
          this.rowData.timeout = Number(row.timeout)
          this.rowData.expr = row.expr
          this.rowData.expect_code = row.expect_code
          this.rowData.expect_context = row.expect_context
          this.rowData.alarm_policy = row.alarm_policy
          this.rowData.description = row.description
          if(!this.rowData.notify.email.receivers){this.rowData.notify.email.receivers=[]}
          if(!this.rowData.notify.webhook.urls){this.rowData.notify.webhook.urls=[]}
          if(!this.rowData.notify.lark.webhooks){this.rowData.notify.lark.webhooks=[]}
          if(!this.rowData.notify.dingding.webhooks){this.rowData.notify.dingding.webhooks=[]}
        },

        newRow: function() {
          this.newOrUpdate = true
          this.title = '添加'
          this.rowData.id = ""
          this.rowData.name = ""
          this.rowData.lang = "shell"
          this.rowData.timeout = -1
          this.rowData.expr = "* * * * * *"
          this.rowData.retry = 0
          this.rowData.expect_code = 0
          this.rowData.expect_context = ""
          this.rowData.alarm_policy = 2
          this.rowData.description = ""
          this.rowData.data = {data:this.langExample[this.rowData.lang],api:{
              url: "",
              method: "GET",
              content_type: "",
              payload: "",
              header: [],
              forms: [],
            }}
          this.rowData.notify = JSON.parse(JSON.stringify(this.notify))
        },

        changeStatus: function (row) {
          let status='stop'
          if(row.status!=="running"){status="running"}
          let api_url = this.url+"/"+row.id+"/"+status
          this.$store.dispatch("common/CreateUpdate",{url: api_url,"method": "PUT", "data": {}}).then((response) => {
            this.$message.success(response.data.message)
            this.getData()
            // this.action_loading = false
          }).catch((response) => {
            this.getData()
            // this.action_loading = false
          })
        },

        changeLang: function(){
          this.rowData.data.data = this.langExample[this.rowData.lang] || ""
        },

        codeInitEditor: function(editor) {
          editor.setAutoScrollEditorIntoView(true);
          editor.setShowPrintMargin(false);
          // editor.on("dblclick", function() {
          //   editor.container.webkitRequestFullscreen();
          // });
          require("brace/ext/language_tools");
          // require("brace/mode/text");
          // require("brace/mode/json");
          require("brace/mode/sh");
          require("brace/mode/python");
          require("brace/mode/golang");
          require('brace/snippets/sh');
          // require('brace/snippets/python');
          require('brace/snippets/golang');
          require("brace/theme/solarized_dark");
        },
        realLogInitEditor: function(editor) {
          editor.setAutoScrollEditorIntoView(true);
          editor.setShowPrintMargin(false);
          editor.on("change", function() {
            editor.renderer.scrollToLine(Number.POSITIVE_INFINITY);
          });
          // require("brace/ext/language_tools");
          // require("brace/mode/sh");
          // require("brace/mode/python");
          // require("brace/mode/javascript");
          require("brace/mode/text");
          // require("brace/mode/json");
          // require("brace/mode/golang");
          require("brace/theme/solarized_dark");
        },
        rowReqBodyInitEditor: function(editor) {
          editor.setAutoScrollEditorIntoView(true);
          editor.setShowPrintMargin(false);
          // require("brace/ext/language_tools");
          // require("brace/mode/text");
          require("brace/mode/json");
          require("brace/theme/solarized_dark");
        },

        runTask: async function (row) {
          try {
            let apiUrl = `${this.url}/${row.id}/run`
            const response = await this.$store.dispatch("common/Request",
              {url: apiUrl,method:"POST", data: {}});
            this.$message({message:response.data.message, type: "success"})
            await this.getData()
          } catch (e) {
            this.$message({message:String(e), type: "error"})
          } finally {
          }
        },

      },
      mounted () {
        this.getData()
      },
      components: {
        pagination: pagination,notify,
        delete_button,
        edit_button,
        user_select,
        green_button,
        editor: require("vue2-ace-editor"),
      },
    }
</script>

<style lang="scss" scoped>
  .add-btn{
    margin-left: 5px;
  }
</style>
