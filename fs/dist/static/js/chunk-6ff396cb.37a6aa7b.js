(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-6ff396cb"],{"53d0":function(t,e,s){"use strict";s("7866")},"77a6":function(t,e,s){"use strict";s.r(e);var r=function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("div",[s("el-row",{staticClass:"system_state",attrs:{gutter:15}},[s("el-col",{attrs:{span:12}},[t.state.os?s("el-card",{staticClass:"card_item"},[s("div",{attrs:{slot:"header"},slot:"header"},[t._v("Runtime")]),t._v(" "),s("div",[s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("os:")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.os.goos)}})],1),t._v(" "),s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("cpu nums:")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.os.numCpu)}})],1),t._v(" "),s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("compiler:")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.os.compiler)}})],1),t._v(" "),s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("go version:")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.os.goVersion)}})],1),t._v(" "),s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("goroutine nums:")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.os.numGoroutine)}})],1)],1)]):t._e()],1),t._v(" "),s("el-col",{attrs:{span:12}},[t.state.disk?s("el-card",{staticClass:"card_item"},[s("div",{attrs:{slot:"header"},slot:"header"},[t._v("Disk")]),t._v(" "),s("div",[s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("total (MB)")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.disk.totalMb)}})],1),t._v(" "),s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("used (MB)")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.disk.usedMb)}})],1),t._v(" "),s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("total (GB)")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.disk.totalGb)}})],1),t._v(" "),s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("used (GB)")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.disk.usedGb)}})],1)],1),t._v(" "),s("el-col",{attrs:{span:12}},[s("el-progress",{attrs:{type:"dashboard",percentage:t.state.disk.usedPercent,color:t.colors}})],1)],1)],1)]):t._e()],1)],1),t._v(" "),s("el-row",{staticClass:"system_state",attrs:{gutter:15}},[s("el-col",{attrs:{span:12}},[t.state.cpu?s("el-card",{staticClass:"card_item",attrs:{"body-style":{height:"180px","overflow-y":"scroll"}}},[s("div",{attrs:{slot:"header"},slot:"header"},[t._v("CPU")]),t._v(" "),s("div",[s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("physical number of cores:")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.cpu.cores)}})],1),t._v(" "),t._l(t.state.cpu.cpus,(function(e,r){return[s("el-row",{key:r,attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("core "+t._s(r)+":")]),t._v(" "),s("el-col",{attrs:{span:12}},[s("el-progress",{attrs:{type:"line",percentage:+e.toFixed(0),color:t.colors}})],1)],1)]}))],2)]):t._e()],1),t._v(" "),s("el-col",{attrs:{span:12}},[t.state.ram?s("el-card",{staticClass:"card_item"},[s("div",{attrs:{slot:"header"},slot:"header"},[t._v("Ram")]),t._v(" "),s("div",[s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("total (MB)")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.ram.totalMb)}})],1),t._v(" "),s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("used (MB)")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.ram.usedMb)}})],1),t._v(" "),s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("total (GB)")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s(t.state.ram.totalMb/1024)}})],1),t._v(" "),s("el-row",{attrs:{gutter:10}},[s("el-col",{attrs:{span:12}},[t._v("used (GB)")]),t._v(" "),s("el-col",{attrs:{span:12},domProps:{textContent:t._s((t.state.ram.usedMb/1024).toFixed(2))}})],1)],1),t._v(" "),s("el-col",{attrs:{span:12}},[s("el-progress",{attrs:{type:"dashboard",percentage:t.state.ram.usedPercent,color:t.colors}})],1)],1)],1)]):t._e()],1)],1)],1)},a=[],o=(s("96cf"),s("3b8d")),l={name:"State",data:function(){return{timer:null,state:{},colors:[{color:"#5cb87a",percentage:20},{color:"#e6a23c",percentage:40},{color:"#f56c6c",percentage:80}],loading:!1,url:this.$store.state.urls.sys_state_url}},created:function(){var t=this;this.getData(),this.timer=setInterval((function(){t.getData()}),1e4)},beforeDestroy:function(){clearInterval(this.timer),this.timer=null},methods:{getData:function(){var t=Object(o["a"])(regeneratorRuntime.mark((function t(){var e,s;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return this.loading=!0,t.prev=1,e=this.url,t.next=5,this.$store.dispatch("common/Query",{url:e,data:{}});case 5:s=t.sent,this.state=s.data.data,t.next=11;break;case 9:t.prev=9,t.t0=t["catch"](1);case 11:return t.prev=11,this.loading=!1,t.finish(11);case 14:case"end":return t.stop()}}),t,this,[[1,9,11,14]])})));function e(){return t.apply(this,arguments)}return e}(),reload:function(){var t=Object(o["a"])(regeneratorRuntime.mark((function t(){var e,s;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,getSystemState();case 2:e=t.sent,s=e.data,this.state=s.server;case 5:case"end":return t.stop()}}),t,this)})));function e(){return t.apply(this,arguments)}return e}()}},n=l,c=(s("53d0"),s("2877")),i=Object(c["a"])(n,r,a,!1,null,"4051f96f",null);e["default"]=i.exports},7866:function(t,e,s){}}]);