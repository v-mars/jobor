package utils

//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"github.com/shirou/gopsutil/v3/cpu"
//	"github.com/shirou/gopsutil/v3/disk"
//	"github.com/shirou/gopsutil/v3/host"
//	"github.com/shirou/gopsutil/v3/mem"
//	"os"
//	"time"
//)
//
////获取内存信息
//func getMemInfo() (map[string]string) {
//	memdata := make(map[string]string)
//
//	v, _ := mem.VirtualMemory()
//
//	total := handerUnit(v.Total,NUM_GB,"G")
//	available := handerUnit(v.Available,NUM_GB,"G")
//	used := handerUnit(v.Used,NUM_GB,"G")
//	free := handerUnit(v.Free,NUM_GB,"G")
//	userPercent := fmt.Sprintf("%.2f",v.UsedPercent)
//
//	memdata["总量"] = total
//	memdata["可用"] = available
//	memdata["已使用"] = used
//	memdata["空闲"] = free
//	memdata["使用率"] = userPercent + "%"
//
//	return memdata
//}
//
////获取CPU信息
//func getCpuInfo(percent string) []map[string]string {
//	cpudatas := []map[string]string{}
//
//	infos, err := cpu.Info()
//	if err != nil {
//		fmt.Printf("get cpu info failed, err:%v", err)
//	}
//	for _, ci := range infos {
//		cpudata := make(map[string]string)
//		cpudata["型号"] = ci.ModelName
//		cpudata["数量"] = fmt.Sprint(ci.Cores)
//		cpudata["使用率"] = percent + "%"
//
//		cpudatas = append(cpudatas, cpudata)
//	}
//	return cpudatas
//}
//
////获取主机信息
//func getHostInfo() map[string]string {
//	hostdata := make(map[string]string)
//
//	hInfo, _ := host.Info()
//	hostdata["主机名称"] = hInfo.Hostname
//	hostdata["系统"] = hInfo.OS
//	hostdata["平台"] = hInfo.Platform + "-" + hInfo.PlatformVersion + " " + hInfo.PlatformFamily
//	hostdata["内核"] = hInfo.KernelArch
//
//	return hostdata
//	//fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
//}
//
//// disk info
//func getDiskInfo() {
//	parts, err := disk.Partitions(true)
//	if err != nil {
//		fmt.Printf("get Partitions failed, err:%v\n", err)
//		return
//	}
//	for _, part := range parts {
//		fmt.Printf("part:%v\n", part.String())
//		diskInfo, _ := disk.Usage(part.Mountpoint)
//		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
//	}
//
//	ioStat, _ := disk.IOCounters()
//	for k, v := range ioStat {
//		fmt.Printf("%v:%v\n", k, v)
//	}
//}
//
//const (
//	NUM_KB = 1000.0000
//	NUM_MIB = 1000000.0000
//	NUM_GB = 1000000000.0000
//)
//
//func handerUnit(num uint64, numtype float64,typename string) (newnum string) {
//
//	f :=fmt.Sprintf("%.2f", float64(num)/numtype)
//	return f + typename
//}
//
////处理日志
//func saveToLocalFile(filename string, data []byte) {
//	//line := []byte("\n")
//	//检查是否存在 存在则追加内容
//	file,err := os.OpenFile(filename,os.O_RDWR,0666)
//
//	if err != nil && os.IsNotExist(err) {
//		//println("文件不存在")//需要加[]
//
//		start := []byte("[")
//		end := []byte("]")
//		newdata := BytesCombine(start,data,end)
//
//		file, _ = os.Create(filename)
//		file.Write(newdata)
//		file.Seek(-1,2)
//	}else {
//		//println("文件存在 追加")
//		//fmt.Println(file.Seek(0, 2))
//
//		start := []byte(",")
//		end := []byte("]")
//		newdata := BytesCombine(start,data,end)
//
//		file.Seek(-1,2)
//		file.Write(newdata)
//	}
//
//	file.Close()
//}
//
////BytesCombine 多个[]byte数组合并成一个[]byte
//func BytesCombine(pBytes ...[]byte) []byte {
//	len := len(pBytes)
//	s := make([][]byte, len)
//	for index := 0; index < len; index++ {
//		s[index] = pBytes[index]
//	}
//	sep := []byte("")
//	return bytes.Join(s, sep)
//}
//
//func main(){
//	println("客户端监控任务开启...")//windows
//
//	var logdate string
//	// CPU使用率
//	for {
//		//限制检测时间
//		nowdatetime := time.Now()
//		hour := nowdatetime.Hour()
//		if hour >= 21 || hour <8{//设置空置时间
//			println("进入睡眠。。。")
//			time.Sleep(time.Minute)
//			continue
//		}
//
//		datas := make(map[string]interface{})
//		//获取内存使用率 同时定时
//		percent, _ := cpu.Percent(time.Second * 14, false)//设置间隔时间
//
//		nowtime := nowdatetime.Format("2006-01-02 15:04:05")//当前时间
//
//		nowdate := nowdatetime.Format("2006-01-02")
//		if logdate == "" || nowdate != logdate {
//			logdate = nowdate
//		}
//		datas["当前时间"] = nowtime
//
//		memdata := getMemInfo()
//		datas["内存信息"] = memdata
//
//		//getDiskInfo()
//
//		hostdata := getHostInfo()
//		datas["主机信息"] = hostdata
//
//		cpudata := getCpuInfo(fmt.Sprintf("%.2f",percent[0]))
//		datas["CPU信息"] = cpudata
//
//		//写入文件
//		jsonStr, err := json.Marshal(datas)
//
//		if err != nil {
//			fmt.Println("MapToJsonDemo err: ", err)
//		}
//		//println(time.Now().Format("2006-01-02 15:04:05"))
//		saveToLocalFile(logdate + ".syslog",jsonStr)
//		//fmt.Println(string(jsonStr))
//	}
//}
