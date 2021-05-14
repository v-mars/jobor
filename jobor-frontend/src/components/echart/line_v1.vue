<template>
  <div>
    <div :id="EleId" :style="{'width':'100%','height':height}"></div>
  </div>
</template>

<script>

  import echarts from 'echarts'
  import macarons from 'echarts/theme/macarons'

    export default {
      name: "line_echart_v1",
      props: {
        data: {required: true},
        EleId: {default: "jobor_status_main"},
        height: {default: "230px"},
      },
      data() {
        return {
          charts: '',
        }
      },
      methods: {
        drawLine(id,interval) {
          let myChart = echarts.init(document.getElementById(id),'macarons')
          // let myChart = this.charts
          if(interval===undefined||interval===null||!interval){
            interval = 4
          }
          if(this.data.xAxisData && this.data.xAxisData.length>4){
            interval = Number(this.data.xAxisData.length/interval)
          }
          this.charts=myChart
          myChart.clear();
          myChart.setOption({
            tooltip: {
              trigger: 'axis',
              confine: false,
            },
            legend: {
              data:this.data.legendData
            },
            // grid: {left: '1%', right: '2%', bottom: '1%', containLabel: true},
            grid:{ x:40, y:35, x2:40, y2:20, borderWidth:1 },

            toolbox: {},
            xAxis: [
              {
                type : 'category',
                boundaryGap : false,
                axisTick: {show: false},
                axisLine: {
                  lineStyle: {
                    color: 'rgba(201,196,196,0.92)'
                  }
                },
                axisLabel:{
                  interval: interval  //控制坐标轴刻度标签的显示间隔.设置成 0 强制显示所有标签。设置为 1，隔一个标签显示一个标签。设置为2，间隔2个标签。以此类推
                },
                data : this.data.xAxisData
              }
            ],
            yAxis: [
              {
                type : 'value',
                axisTick: {show: false},
                axisLine: {
                  lineStyle: {
                    color: 'rgba(201,196,196,0.92)'
                  }
                },
                splitLine: {
                  lineStyle: {
                    color: 'rgba(201,196,196,0.92)',
                    type:'dashed'  //虚线
                  }
                },
                name: ''
              }
            ],

            series: this.data.series
          })
          // this.changeColor(myChart)
          // console.log(myChart.getModel().getSeriesByIndex(0).getData().getVisual('color'))
          window.onresize = function () {
            //重置容器高宽
            if(myChart){
              // setContentSize(worldMapContainer);
              myChart.resize()
            }
          }
        },

      },
      //调用
      mounted() {
        // this.$nextTick(function() {
        //   this.drawLine(this.EleId)
        // })

      }
    }
</script>

<style scoped>

</style>
