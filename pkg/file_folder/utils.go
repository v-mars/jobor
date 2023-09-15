package file_folder

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

func CheckOrCreateFolder(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		if err = os.Mkdir(folderPath, 0755); err != nil {
			panic(any(err))
		}
		// 再修改权限
		//err = os.Chmod(folderPath, 0777)
		//if err != nil {
		//	return "", ""
		//}
		//if err = os.Chown("",1,1);err != nil {
		//	return err
		//}
	}
	return nil
}

func RemoveFileOrFolder(filePath string) error {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		err = nil
	} else {
		err = os.RemoveAll(filePath)
		if err != nil {
			return fmt.Errorf("filepath [%s] remove conf file err, %s", filePath, err)
		}
	}
	return nil
}

func CopyFileOrFolder(from, to string) error {
	if from == to {
		return fmt.Errorf("from and to is not same")
	}
	f, err := os.Stat(from)
	if err != nil {
		return err
	}
	if f.IsDir() {
		// from 是目录，那么定义的to也是目录
		var list []fs.FileInfo
		if _, err = os.Stat(to); err != nil && os.IsNotExist(err) {
			if err = os.MkdirAll(to, f.Mode()); err != nil {
				return err
			}
		}
		if list, err = ioutil.ReadDir(from); err == nil {
			for _, item := range list {
				if err = CopyFileOrFolder(filepath.Join(from, item.Name()),
					filepath.Join(to, item.Name())); err != nil {
					return err
				}
			}
		}
	} else {
		// from 是文件，那么定义的to也是文件
		p := filepath.Dir(to)
		if _, err = os.Stat(p); err != nil && os.IsNotExist(err) {
			if err = os.MkdirAll(p, 0755); err != nil {
				return err
			}
		}
		// 读取原文件
		var file *os.File
		file, err = os.Open(from)
		if err != nil {
			return err
		}
		defer file.Close()
		bufReader := bufio.NewReader(file)
		//创建一个文件用于保存
		var out *os.File
		if _, err = os.Stat(to); err != nil && os.IsNotExist(err) {
			out, err = os.Create(to)
			if err != nil {
				return err
			}
		} else {
			out, err = os.Open(to)
			if err != nil {
				return err
			}
		}

		defer out.Close()
		// 然后将文件流与文件流对接起来
		_, err = io.Copy(out, bufReader)
		if err != nil {
			return err
		}
	}
	return nil
}

type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	//return fis[i].ModTime().Before(fis[j].ModTime())   // 从旧往新排
	return fis[j].ModTime().Before(fis[i].ModTime()) // 从新往旧排
}

// KeepLastedFolder
/**
保留最新的文件、文件数量
param: path
param: keepFileNum
param: keepFolderNum
*/
func KeepLastedFolder(path string, keepFileNum, keepFolderNum int) error {
	if len(path) == 0 || path == "/" {
		return fmt.Errorf("path【%s】 信息不合规", path)
	}
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	fis, err := f.Readdir(-1)
	if err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	sort.Sort(ByModTime(fis))

	var folders []os.FileInfo
	var files []os.FileInfo
	for _, fi := range fis {
		//fmt.Println(fi.Name(), fi.IsDir())
		if fi.IsDir() {
			folders = append(folders, fi)
		} else {
			files = append(files, fi)
		}
	}

	if len(folders) > keepFolderNum {
		for _, fi := range folders[keepFolderNum:] {
			//fmt.Println("folder del:", fi.Name(), filepath.Join(path, fi.Name()))
			if err = os.RemoveAll(filepath.Join(path, fi.Name())); err != nil {
				return err
			}
		}
	}
	if len(files) > keepFileNum {
		for _, fi := range files[keepFileNum:] {
			//fmt.Println("file del:", fi.Name(), fi.Sys())
			if err = os.Remove(filepath.Join(path, fi.Name())); err != nil {
				return err
			}
		}
	}
	return nil
}
