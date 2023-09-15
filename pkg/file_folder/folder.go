package file_folder

import (
	"os"
	"path/filepath"
)

func CreateDir(basePath, folderName string) (dirPath, dataString string) {
	folderPath := filepath.Join(basePath, folderName)
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
	}
	return folderPath, folderName
}
