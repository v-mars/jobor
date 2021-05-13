import echarts from "echarts";


export function echart_color(color, offset) {
  return new echarts.graphic.LinearGradient(0, 0, 0, 1, [
    {offset: 0, color: color},         // rgb(50, 109, 230)
    {offset: offset, color: 'rgba(61,234,255,0)'}    // 1
  ])

}


export default {
  echart_color,
}
