<template>
    <div>
      <panel_group :tasks="data.tasks" :today_task_logs="data.today_task_logs"
                   :workers="data.workers" :task_logs="data.task_logs"></panel_group>
      <el-card class="jobor-dashboard" size="mini">
        <label>时间范围：</label>
        <datetime-range :value.sync="dateRange" @change="changeDateRange"></datetime-range>
        <el-button size="mini" @click="getData" type="primary" style="position: absolute;right: 10px">刷新</el-button>
      </el-card>
      <el-card class="jobor-dashboard" size="mini">
        <label style="font-size: 20px">任务状态趋势</label>
        <el-row :gutter="20" class="">
          <el-col :xs="24" :sm="24" :lg="16" class="card-panel-col">
            <line_v1 :data.sync="taskLogStatusDay" ref="line_task_status" ele-id="line-task-status"></line_v1>
          </el-col>
          <el-col :xs="24" :sm="24" :lg="8" class="card-panel-col">
            <el-table size="mini" key="user-deployment" :data="data.task_run||[]" height="230">
<!--              <el-table-column type="index" :index="index" width="50" label="排名" align="center"></el-table-column>-->
              <el-table-column label="排行" prop="ord_num" width="50" align="center">
                <template slot-scope="scope">
                  <span style="color: #ff4d4f" v-if="scope.row.ord_num===1">1</span>
                  <span style="color: #ffa940" v-else-if="scope.row.ord_num===2">2</span>
                  <span style="color: #ffc53d" v-else-if="scope.row.ord_num===3">3</span>
                  <span style="" v-else>{{scope.row.ord_num}}</span>
                </template>
              </el-table-column>
              <el-table-column label="任务" prop="task" width=""></el-table-column>
              <el-table-column label="运行次数" prop="count" width=""></el-table-column>
            </el-table>
          </el-col>
        </el-row>
      </el-card>
    </div>
</template>

<script>
  import panelgroup from "@/components/jobor/dashboard/panelgroup";
  import DatetimeRange from "@/components/jobor/dashboard/DatetimeRange";
  import line_v1 from "@/components/echart/line_v1";
  import {echart_color} from "@/components/js/echart";
    export default {
        name: "dashboard",
      data(){return{
        url: this.$store.state.urls.jobor_dash_url,
        dateRange: [],
        data: {
          tasks: 0,
          task_logs: 0,
          today_task_logs: 0,
          workers: 0,
          project_deployments: [],
          task_run: [],
          task_log_status_day:[],
        },
        taskLogStatusDay: {
          legendData: [],
          xAxisData: [],
          series: [],
        },
      }},
      methods: {
        index(val){
          return val + 1
        },

        getData: async function () {
          let data = {start_time: this.dateRange[0], end_time:this.dateRange[1]}
          this.loading = true
          try {
            let response = await this.$store.dispatch("common/Query", {url: this.url, data: data})
            this.data = response.data.data
            // let pdd = this.data.project_deployments
            // for(let i=0;i<pdd.length;i++){
            //   if(this.projectDayDeployment.xAxisData.indexOf(pdd[i].time)===-1){
            //     this.projectDayDeployment.xAxisData.push(pdd[i].time)
            //   }
            //   if(this.projectDayDeployment.legendData.indexOf(pdd[i].project_name)===-1){
            //     this.projectDayDeployment.legendData.push(pdd[i].project_name)
            //   }
            // }
            // for(let i=0;i<this.projectDayDeployment.legendData.length;i++){
            //   let temp = {
            //     name:this.projectDayDeployment.legendData[i],
            //     type:'line',
            //     symbolSize:1,
            //     smooth: false,
            //     itemStyle: {normal: {areaStyle: {
            //           // type: 'default',
            //           color: echart_color('rgba(86,144,132,0.32)', 1),
            //         }}},
            //     emphasis: {
            //       focus: 'series'
            //     },
            //     data:[]
            //   }
            //   for(let i1=0;i1<this.projectDayDeployment.xAxisData.length;i1++){
            //     temp.data.push(0)
            //     for(let i2=0;i2<pdd.length;i2++){
            //       if(this.projectDayDeployment.xAxisData[i1]===pdd[i2].time &&
            //         this.projectDayDeployment.legendData[i]===pdd[i2].project_name){
            //         temp.data[i1] = pdd[i2].count
            //         break
            //       }
            //     }
            //   }
            //   this.projectDayDeployment.series.push(temp)
            // }
            // this.$refs.pdd.drawLine("line-pdd",8)

            let dsf = this.data.task_log_status_day
            for(let i=0;i<dsf.length;i++){
              if(this.taskLogStatusDay.xAxisData.indexOf(dsf[i].time)===-1){
                this.taskLogStatusDay.xAxisData.push(dsf[i].time)
              }
            }
            this.taskLogStatusDay.legendData = ["总计", "成功", "失败"]
            this.taskLogStatusDay.series = [
              {
                name:'总计',
                type:'line',
                symbolSize:1,
                smooth: false,
                lineStyle: {
                  color: 'rgb(41,195,215)'
                },
                itemStyle: {normal: {color:'rgb(41,195,215)',areaStyle: {
                      // type: 'default',
                      color: echart_color('rgb(41,195,215)', 1),
                    }}},
                emphasis: {
                  focus: 'series'
                },
                data:[]
              },
              {
                name:'成功',
                type:'line',
                symbolSize:1,
                smooth: false,
                lineStyle: {
                  color: 'rgb(0, 199, 82)'
                },
                itemStyle: {normal: {color:'rgb(0, 199, 82)',areaStyle: {
                      // type: 'default',
                      color: echart_color('rgb(0, 199, 82)', 1),
                    }}},
                emphasis: {
                  focus: 'series'
                },
                data:[]
              },
              {
                name:'失败',
                type:'line',
                symbolSize:1,
                smooth: false,
                lineStyle: {
                  color: 'rgba(244,81,108,0.89)'
                },
                itemStyle: {normal:
                    {
                      color:'rgba(244,81,108,0.89)',
                      areaStyle: {
                        // type: 'default',
                        color: echart_color('rgba(244,81,108,0.89)', 1),
                      }}},
                emphasis: {
                  focus: 'series'
                },
                data:[]
              }
            ]
            for(let i1=0;i1<this.taskLogStatusDay.xAxisData.length;i1++){
              for(let i2=0;i2<dsf.length;i2++){
                if(this.taskLogStatusDay.xAxisData[i1]===dsf[i2].time){
                  this.taskLogStatusDay.series[0].data.push(dsf[i2].count)
                  this.taskLogStatusDay.series[1].data.push(dsf[i2].success_count)
                  this.taskLogStatusDay.series[2].data.push(dsf[i2].failed_count)
                  break
                }
              }
            }
            this.$refs.line_task_status.drawLine("line-task-status",1)


            // this.total = response.data.data.total
          } catch (e) {
            this.$message.warning(e)
          } finally {
            this.loading = false
          }
        },

        changeDateRange: function (val) {
          if(val){
            this.dateRange = val
            this.getData()
          }
        }
      },
      async mounted() {
        await this.getData()
      },
      components:{
        panel_group:panelgroup,
        DatetimeRange,line_v1,

      }
    }
</script>

<style lang="scss" scoped>
  .jobor-dashboard{
    margin-bottom: 10px;
  }
</style>
