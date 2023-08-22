package utils

import (
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

func GetLocalIPv4Address() (string, error) {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addr {

		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				return ipv4.String(), nil
			}
		}
	}
	return "", fmt.Errorf("not found ipv4 address")
}

// GetOutBoundIP net.Dial("udp", "8.8.8.8:53")
func GetOutBoundIP(network, addr string) (ip string) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		//log.Errorf("get out bound ip err: %s\n", err)
		panic(any(err))
	}
	var localAddr net.Addr
	if network == "tcp" {
		localAddr = conn.LocalAddr().(*net.TCPAddr) // .(*net.TCPAddr)
	} else {
		localAddr = conn.LocalAddr().(*net.UDPAddr) // .(*net.UDPAddr)
	}
	//fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func EnsureDirExist(name string) error {
	if !FileExists(name) {
		return os.MkdirAll(name, os.ModePerm)
	}
	return nil
}

func GzipCompressFile(srcPath, dstPath string) error {
	sf, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer func(sf *os.File) {
		err := sf.Close()
		if err != nil {

		}
	}(sf)
	df, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer func(df *os.File) {
		err := df.Close()
		if err != nil {

		}
	}(df)
	writer := gzip.NewWriter(df)
	writer.Name = dstPath
	writer.ModTime = time.Now().UTC()
	_, err = io.Copy(writer, sf)
	if err != nil {
		return err
	}
	if err := writer.Close(); err != nil {
		return err
	}
	return nil
}

func Sum(i []int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CurrentUTCTime() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05 +0000")
}
