(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-79ef7dc7"],{"3f88":function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("div",{staticStyle:{margin:"5px 0"}},[a("el-form",{staticClass:"demo-form-inline",attrs:{inline:!0,model:t.searchForm,size:"small"},nativeOn:{submit:function(t){t.preventDefault()}}},[a("el-form-item",{attrs:{label:""}},[a("el-input",{attrs:{placeholder:"主机名"},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.getData(e)}},model:{value:t.searchForm.hostname,callback:function(e){t.$set(t.searchForm,"hostname",e)},expression:"searchForm.hostname"}})],1),t._v(" "),a("el-form-item",{attrs:{label:""}},[a("el-input",{attrs:{placeholder:"地址"},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.getData(e)}},model:{value:t.searchForm.addr,callback:function(e){t.$set(t.searchForm,"addr",e)},expression:"searchForm.addr"}})],1),t._v(" "),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:t.getData}},[t._v("查询")]),t._v(" "),a("el-button",{staticStyle:{float:"right"},on:{click:t.getData}},[t._v("刷新")])],1)],1)],1),t._v(" "),a("div",{staticStyle:{"margin-top":"10px"}},[a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.loading,expression:"loading"}],attrs:{border:"",data:t.data_list,size:"small"}},[a("el-table-column",{attrs:{label:"ID",prop:"id",width:"60"}}),t._v(" "),a("el-table-column",{attrs:{label:"主机名称",prop:"hostname",width:"150"}}),t._v(" "),a("el-table-column",{attrs:{label:"地址",prop:"addr",width:"160"}}),t._v(" "),a("el-table-column",{attrs:{label:"路由标识",prop:"routing_key",width:""}}),t._v(" "),a("el-table-column",{attrs:{label:"权重",prop:"weight",width:"100"}}),t._v(" "),a("el-table-column",{attrs:{label:"状态",prop:"",width:"100",align:""},scopedSlots:t._u([{key:"default",fn:function(t){return["running"===t.row.status?a("div",{staticStyle:{"vertical-align":"middle",display:"flex","flex-direction":"row","align-items":"center"}},[a("Badge",{attrs:{status:"success",text:"运行中"}})],1):"stop"===t.row.status?a("div",[a("Badge",{attrs:{status:"warning",text:"停止"}})],1):"offline"===t.row.status?a("div",[a("Badge",{attrs:{status:"error",text:"离线"}})],1):a("div",[a("Badge",{attrs:{status:"default",text:t.row.result}})],1)]}}])}),t._v(" "),a("el-table-column",{attrs:{label:"更新时间",prop:"mtime",width:"140"}}),t._v(" "),a("el-table-column",{attrs:{label:"操作",align:"center",width:"120"},scopedSlots:t._u([{key:"default",fn:function(e){return[-1!==["stop"].indexOf(e.row.status)?a("el-popconfirm",{attrs:{icon:"el-icon-info",title:"确认启用此节点吗？"},on:{onConfirm:function(a){return t.changeWorkerStatus(e.row,"running")}}},[a("green_button",{attrs:{slot:"reference",title:"启用"},slot:"reference"})],1):t._e(),t._v(" "),-1!==["running"].indexOf(e.row.status)?a("el-popconfirm",{attrs:{icon:"el-icon-info",title:"确认禁用此节点吗？"},on:{onConfirm:function(a){return t.changeWorkerStatus(e.row,"stop")}}},[a("edit_button",{attrs:{slot:"reference",title:"禁用"},slot:"reference"})],1):t._e(),t._v(" "),a("delete_button",{attrs:{title:"删除"},on:{click:function(a){return t.confirmDelRows(e.row,e.row.addr,e.row.id)}}})]}}])})],1),t._v(" "),a("div",{staticStyle:{"margin-top":"5px"}},[a("div",{staticStyle:{display:"inline"}}),t._v(" "),a("div",{staticClass:"block",staticStyle:{float:"right",display:"inline"}},[a("pagination",{attrs:{total:t.total,page:t.page,limit:t.limit},on:{"update:page":function(e){t.page=e},"update:limit":function(e){t.limit=e},pagination:function(e){return t.getData()}}})],1)])],1)])},r=[],i=(a("96cf"),a("3b8d")),l=a("26e4"),o=a("5d46"),s=a("1db2"),c=a("d329"),u=a("8fb7d"),d=a("aa11"),m=a("fe8c"),p={name:"jobor_task_worker",mixins:[m["a"]],data:function(){return{url:this.$store.state.urls.jobor_worker_url,searchForm:{addr:"",hostname:""}}},methods:{changeWorkerStatus:function(){var t=Object(i["a"])(regeneratorRuntime.mark((function t(e,a){var n,r,i;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.prev=0,n={status:a},r="".concat(this.url,"/").concat(e.id),t.next=5,this.$store.dispatch("common/Request",{url:r,method:"PUT",data:n});case 5:return i=t.sent,this.$message({message:i.data.message,type:"success"}),t.next=9,this.getData();case 9:t.next=14;break;case 11:t.prev=11,t.t0=t["catch"](0),this.$message({message:String(t.t0),type:"error"});case 14:return t.prev=14,t.finish(14);case 16:case"end":return t.stop()}}),t,this,[[0,11,14,16]])})));function e(e,a){return t.apply(this,arguments)}return e}()},mounted:function(){this.getData()},components:{pagination:l["a"],delete_button:o["a"],edit_button:s["a"],user_select:d["a"],green_button:c["a"],detail_button:u["a"]}},f=p,h=a("2877"),g=Object(h["a"])(f,n,r,!1,null,"8a80fa76",null);e["default"]=g.exports},"8fb7d":function(t,e,a){"use strict";var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-button",{attrs:{type:"text",size:t.size},on:{click:t.handleClick}},[t._v(t._s(t.title))])},r=[],i={name:"delete",props:{size:{default:"mini"},title:{default:"详细"}},methods:{handleClick:function(t){this.$emit("click",t)}}},l=i,o=a("2877"),s=Object(o["a"])(l,n,r,!1,null,"14939b4e",null);e["a"]=s.exports}}]);