(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-3884e73c"],{"0c45":function(t,a,e){"use strict";e("d865")},"229c":function(t,a,e){"use strict";e("bb6b")},"493c":function(t,a,e){"use strict";e.r(a);var s=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",[e("panel_group",{attrs:{tasks:t.data.tasks,today_task_logs:t.data.today_task_logs,workers:t.data.workers,task_logs:t.data.task_logs}}),t._v(" "),e("el-card",{staticClass:"jobor-dashboard",attrs:{size:"mini"}},[e("label",[t._v("时间范围：")]),t._v(" "),e("datetime-range",{attrs:{value:t.dateRange},on:{"update:value":function(a){t.dateRange=a},change:t.changeDateRange}}),t._v(" "),e("el-button",{staticStyle:{position:"absolute",right:"10px"},attrs:{size:"mini",type:"primary"},on:{click:t.getData}},[t._v("刷新")])],1),t._v(" "),e("el-card",{staticClass:"jobor-dashboard",attrs:{size:"mini"}},[e("label",{staticStyle:{"font-size":"20px"}},[t._v("任务状态趋势")]),t._v(" "),e("el-row",{attrs:{gutter:20}},[e("el-col",{staticClass:"card-panel-col",attrs:{xs:24,sm:24,lg:16}},[e("line_v1",{ref:"line_task_status",attrs:{data:t.taskLogStatusDay,"ele-id":"line-task-status"},on:{"update:data":function(a){t.taskLogStatusDay=a}}})],1),t._v(" "),e("el-col",{staticClass:"card-panel-col",attrs:{xs:24,sm:24,lg:8}},[e("el-table",{key:"user-deployment",attrs:{size:"mini",data:t.data.task_run||[],height:"230"}},[e("el-table-column",{attrs:{label:"排名",prop:"ord_num",width:"50",align:"center"},scopedSlots:t._u([{key:"default",fn:function(a){return[1===a.row.ord_num?e("span",{staticStyle:{color:"#ff4d4f"}},[t._v("1")]):2===a.row.ord_num?e("span",{staticStyle:{color:"#ffa940"}},[t._v("2")]):3===a.row.ord_num?e("span",{staticStyle:{color:"#ffc53d"}},[t._v("3")]):e("span",{},[t._v(t._s(a.row.ord_num))])]}}])}),t._v(" "),e("el-table-column",{attrs:{label:"任务",prop:"task",width:""}}),t._v(" "),e("el-table-column",{attrs:{label:"运行次数",prop:"count",width:""}})],1)],1)],1)],1)],1)},n=[],i=(e("96cf"),e("3b8d")),r=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("el-row",{staticClass:"panel-group",attrs:{gutter:20}},[e("el-col",{staticClass:"card-panel-col",attrs:{xs:12,sm:12,lg:6}},[e("div",{staticClass:"card-panel",on:{click:function(a){return t.handleSetLineChartData("messages")}}},[e("div",{staticClass:"card-panel-icon-wrapper icon-ic-today"},[e("svg-icon",{attrs:{"icon-class":"ic-today","class-name":"card-panel-icon"}})],1),t._v(" "),e("div",{staticClass:"card-panel-description"},[e("div",{staticClass:"card-panel-text"},[t._v("\n            今天执行任务数\n          ")]),t._v(" "),e("count-to",{staticClass:"card-panel-num",attrs:{"start-val":0,"end-val":t.today_task_logs,duration:3e3}})],1)])]),t._v(" "),e("el-col",{staticClass:"card-panel-col",attrs:{xs:12,sm:12,lg:6}},[e("div",{staticClass:"card-panel",on:{click:function(a){return t.handleSetLineChartData("app_v1")}}},[e("div",{staticClass:"card-panel-icon-wrapper icon-app_v1"},[e("svg-icon",{attrs:{"icon-class":"app_v1","class-name":"card-panel-icon"}})],1),t._v(" "),e("div",{staticClass:"card-panel-description"},[e("div",{staticClass:"card-panel-text"},[t._v("\n            节点数\n          ")]),t._v(" "),e("count-to",{staticClass:"card-panel-num",attrs:{"start-val":0,"end-val":t.workers,duration:3600}})],1)])]),t._v(" "),e("el-col",{staticClass:"card-panel-col",attrs:{xs:12,sm:12,lg:6}},[e("div",{staticClass:"card-panel",on:{click:function(a){return t.handleSetLineChartData("guide")}}},[e("div",{staticClass:"card-panel-icon-wrapper icon-guide"},[e("svg-icon",{staticStyle:{color:"#00A2CA"},attrs:{"icon-class":"guide","class-name":"card-panel-icon"}})],1),t._v(" "),e("div",{staticClass:"card-panel-description"},[e("div",{staticClass:"card-panel-text"},[t._v("\n            总计任务数\n          ")]),t._v(" "),e("count-to",{staticClass:"card-panel-num",attrs:{"start-val":0,"end-val":t.tasks,duration:1800}})],1)])]),t._v(" "),e("el-col",{staticClass:"card-panel-col",attrs:{xs:12,sm:12,lg:6}},[e("div",{staticClass:"card-panel",on:{click:function(a){return t.handleSetLineChartData("purchases")}}},[e("div",{staticClass:"card-panel-icon-wrapper icon-build_v1"},[e("svg-icon",{attrs:{"icon-class":"build_v1","class-name":"card-panel-icon"}})],1),t._v(" "),e("div",{staticClass:"card-panel-description"},[e("div",{staticClass:"card-panel-text"},[t._v("\n            总计执行任务数\n          ")]),t._v(" "),e("count-to",{staticClass:"card-panel-num",attrs:{"start-val":0,"end-val":t.task_logs,duration:3e3}})],1)])])],1)},l=[],o=e("ec1b"),c=e.n(o),d={name:"panelgroup",props:{tasks:{default:0},task_logs:{default:0},workers:{default:0},today_task_logs:{default:0}},methods:{handleSetLineChartData:function(t){this.$emit("handleSetLineChartData",t)}},components:{CountTo:c.a}},u=d,p=(e("229c"),e("2877")),h=Object(p["a"])(u,r,l,!1,null,"38cb5094",null),g=h.exports,m=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("el-date-picker",{attrs:{type:"datetimerange","picker-options":t.pickerOptions,"range-separator":"至","start-placeholder":"开始日期","end-placeholder":"结束日期",size:t.size,clearable:!1,align:"left"},on:{change:t.valueChange},model:{value:t.valueCurr,callback:function(a){t.valueCurr=a},expression:"valueCurr"}})},v=[],_={name:"DatetimeRange",props:{size:{default:"mini"},value:{Type:Array}},data:function(){return{pickerOptions:{shortcuts:[{text:"最近一周",onClick:function(t){var a=new Date,e=new Date;e.setTime(e.getTime()-6048e5),t.$emit("pick",[e,a])}},{text:"最近一个月",onClick:function(t){var a=new Date,e=new Date;e.setTime(e.getTime()-2592e6),t.$emit("pick",[e,a])}},{text:"最近三个月",onClick:function(t){var a=new Date,e=new Date;e.setTime(e.getTime()-7776e6),t.$emit("pick",[e,a])}},{text:"最近6个月",onClick:function(t){var a=new Date,e=new Date;e.setTime(e.getTime()-15552e6),t.$emit("pick",[e,a])}},{text:"最近1年",onClick:function(t){var a=new Date,e=new Date;e.setTime(e.getTime()-31536e6),t.$emit("pick",[e,a])}}]},valueCurr:[new Date((new Date).getTime()-2592e6),new Date]}},methods:{valueChange:function(t){this.$emit("update:value",this.valueCurr),this.$emit("change",t)}},mounted:function(){this.valueChange()}},f=_,k=Object(p["a"])(f,m,v,!1,null,"e124598a",null),b=k.exports,y=function(){var t=this,a=t.$createElement,e=t._self._c||a;return e("div",[e("div",{style:{width:"100%",height:t.height},attrs:{id:t.EleId}})])},x=[],w=(e("c5f6"),e("313e")),D=e.n(w),C=(e("817d"),{name:"line_echart_v1",props:{data:{required:!0},EleId:{default:"jobor_status_main"},height:{default:"230px"}},data:function(){return{charts:""}},methods:{drawLine:function(t,a){var e=D.a.init(document.getElementById(t),"macarons");void 0!==a&&null!==a&&a||(a=4),this.data.xAxisData&&this.data.xAxisData.length>4&&(a=Number(this.data.xAxisData.length/a)),this.charts=e,e.clear(),e.setOption({tooltip:{trigger:"axis",confine:!1},legend:{data:this.data.legendData},grid:{x:40,y:35,x2:40,y2:20,borderWidth:1},toolbox:{},xAxis:[{type:"category",boundaryGap:!1,axisTick:{show:!1},axisLine:{lineStyle:{color:"rgba(201,196,196,0.92)"}},axisLabel:{interval:a},data:this.data.xAxisData}],yAxis:[{type:"value",axisTick:{show:!1},axisLine:{lineStyle:{color:"rgba(201,196,196,0.92)"}},splitLine:{lineStyle:{color:"rgba(201,196,196,0.92)",type:"dashed"}},name:""}],series:this.data.series}),window.onresize=function(){e&&e.resize()}}},mounted:function(){}}),S=C,L=Object(p["a"])(S,y,x,!1,null,"bb3cece4",null),T=L.exports;function $(t,a){return new D.a.graphic.LinearGradient(0,0,0,1,[{offset:0,color:t},{offset:a,color:"rgba(61,234,255,0)"}])}var A={name:"dashboard",data:function(){return{url:this.$store.state.urls.jobor_dash_url,dateRange:[],data:{tasks:0,task_logs:0,today_task_logs:0,workers:0,project_deployments:[],task_run:[],task_log_status_day:[]},taskLogStatusDay:{legendData:[],xAxisData:[],series:[]}}},methods:{index:function(t){return t+1},getData:function(){var t=Object(i["a"])(regeneratorRuntime.mark((function t(){var a,e,s,n,i,r;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return a={start_time:this.dateRange[0],end_time:this.dateRange[1]},this.loading=!0,t.prev=2,t.next=5,this.$store.dispatch("common/Query",{url:this.url,data:a});case 5:for(e=t.sent,this.data=e.data.data,s=this.data.task_log_status_day,n=0;n<s.length;n++)-1===this.taskLogStatusDay.xAxisData.indexOf(s[n].time)&&this.taskLogStatusDay.xAxisData.push(s[n].time);this.taskLogStatusDay.legendData=["总计","成功","失败"],this.taskLogStatusDay.series=[{name:"总计",type:"line",symbolSize:1,smooth:!1,lineStyle:{color:"rgb(41,195,215)"},itemStyle:{normal:{color:"rgb(41,195,215)",areaStyle:{color:$("rgb(41,195,215)",1)}}},emphasis:{focus:"series"},data:[]},{name:"成功",type:"line",symbolSize:1,smooth:!1,lineStyle:{color:"rgb(0, 199, 82)"},itemStyle:{normal:{color:"rgb(0, 199, 82)",areaStyle:{color:$("rgb(0, 199, 82)",1)}}},emphasis:{focus:"series"},data:[]},{name:"失败",type:"line",symbolSize:1,smooth:!1,lineStyle:{color:"rgba(244,81,108,0.89)"},itemStyle:{normal:{color:"rgba(244,81,108,0.89)",areaStyle:{color:$("rgba(244,81,108,0.89)",1)}}},emphasis:{focus:"series"},data:[]}],i=0;case 12:if(!(i<this.taskLogStatusDay.xAxisData.length)){t.next=26;break}r=0;case 14:if(!(r<s.length)){t.next=23;break}if(this.taskLogStatusDay.xAxisData[i]!==s[r].time){t.next=20;break}return this.taskLogStatusDay.series[0].data.push(s[r].count),this.taskLogStatusDay.series[1].data.push(s[r].success_count),this.taskLogStatusDay.series[2].data.push(s[r].failed_count),t.abrupt("break",23);case 20:r++,t.next=14;break;case 23:i++,t.next=12;break;case 26:this.$refs.line_task_status.drawLine("line-task-status"),t.next=31;break;case 29:t.prev=29,t.t0=t["catch"](2);case 31:return t.prev=31,this.loading=!1,t.finish(31);case 34:case"end":return t.stop()}}),t,this,[[2,29,31,34]])})));function a(){return t.apply(this,arguments)}return a}(),changeDateRange:function(t){t&&(this.dateRange=t,this.getData())}},mounted:function(){var t=Object(i["a"])(regeneratorRuntime.mark((function t(){return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,this.getData();case 2:case"end":return t.stop()}}),t,this)})));function a(){return t.apply(this,arguments)}return a}(),components:{panel_group:g,DatetimeRange:b,line_v1:T}},R=A,z=(e("0c45"),Object(p["a"])(R,s,n,!1,null,"15443772",null));a["default"]=z.exports},bb6b:function(t,a,e){},d865:function(t,a,e){}}]);