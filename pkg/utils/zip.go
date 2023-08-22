package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func IsZip(src string) (bool, error) {
	f, err := os.Stat(src)
	if err != nil {
		return false, nil
	}

	if f.IsDir() {
		return false, nil
	}

	return ExtZip == filepath.Ext(f.Name()), nil

}

func IsDir(src string) (bool, error) {
	f, err := os.Stat(src)
	if err != nil {
		return false, err
	}

	return f.IsDir(), nil
}

func IsFileExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func GetSrcZips(src string) ([]string, error) {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return nil, err
	}
	fileNames := make([]string, 0)
	for _, f := range files {
		ext := filepath.Ext(f.Name())
		if ext != ExtZip {
			continue
		}

		fileNames = append(fileNames, strings.TrimSuffix(f.Name(), ext))
	}

	return fileNames, nil
}

// ZipCompress
// zip压缩
func ZipCompress(srcDir, dstZipPath string) error {
	if filepath.Ext(filepath.Base(dstZipPath)) != ExtZip {
		return fmt.Errorf("not a zip file")
	}

	dstDir := filepath.Dir(dstZipPath)
	isDstDir, err := IsDir(dstDir)
	if err != nil {
		return err
	}
	if !isDstDir {
		return fmt.Errorf("dstDir is not a directory")
	}

	isSrcDir, err := IsDir(srcDir)
	if err != nil {
		return err
	}
	if !isSrcDir {
		return fmt.Errorf("srcDirs is not a directory")
	}

	f, err := os.Create(dstZipPath)
	if err != nil {
		return err
	}
	defer f.Close()

	zw := zip.NewWriter(f)
	defer zw.Close()

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, fi := range files {
		if err := compress(fi, srcDir, "", zw); err != nil {
			return err
		}
	}

	return nil
}

func compress(fi os.FileInfo, fileDir string, subName string, zw *zip.Writer) error {
	if fi.IsDir() {
		fileDir := filepath.Join(subName, fi.Name())
		if subName != "" {
			subName = filepath.Join(subName, fi.Name())
		} else {
			subName = fi.Name()
		}

		header, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}
		header.Name = subName + "/"

		_, err = zw.CreateHeader(header)
		if err != nil {
			return err
		}

		files, err := ioutil.ReadDir(fileDir)
		if err != nil {
			return err
		}
		for _, fi := range files {
			if err := compress(fi, fileDir, subName, zw); err != nil {
				return err
			}
		}
	} else {
		filePath := filepath.Join(fileDir, fi.Name())
		f, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer f.Close()

		header, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}
		if subName != "" {
			header.Name = filepath.Join(subName, fi.Name())
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, f)
		if err != nil {
			return err
		}
	}

	return nil
}

const (
	ExtZip = ".zip"
	// maxFileCount 压缩文件最大文件数
	maxFileCount = 15000
	// singleFileMaxSize 单个文件解压和总文件解压不能超过10G
	singleFileMaxSize uint64 = 10737418240
	fileMaxSize       int64  = 10737418240
)

// UnZipCompress
// zip解压
func UnZipCompress(src, dst string) error {
	isZip, err := IsZip(src)
	if err != nil {
		return err
	}
	if !isZip {
		return fmt.Errorf("src is not a zip file")
	}

	isDir, err := IsDir(dst)
	if err != nil {
		return err
	}
	if !isDir {
		return fmt.Errorf("dst is not a directory")
	}

	srcReader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer srcReader.Close()

	if len(srcReader.File) > maxFileCount {
		return fmt.Errorf("zip file have too many files")
	}

	for _, f := range srcReader.File {
		// safety requirements
		if strings.Contains(f.Name, "./") ||
			strings.Contains(f.Name, ".\\") ||
			strings.Contains(f.Name, "..") {
		}
		if f.UncompressedSize64 > singleFileMaxSize {
			return fmt.Errorf("single file exceeds the maximum size")
		}
	}
	err = writeReader(srcReader, dst)
	if err != nil {
		return err
	}

	return nil
}

func writeReader(srcReader *zip.ReadCloser, dst string) error {
	var totalSize int64
	for _, f := range srcReader.File {
		if totalSize > fileMaxSize {
			return fmt.Errorf("the total size has exceeds the upper limit")
		}
		fileName := f.Name
		targetFilePath := filepath.Join(dst, fileName)

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(targetFilePath, f.Mode())
			if err != nil {
				return err
			}
			continue
		}

		isExist, err := IsFileExist(targetFilePath)
		if err != nil {
			return err
		}
		if isExist {
			return fmt.Errorf("the target path has exist")
		}

		if err := os.MkdirAll(path.Dir(targetFilePath), 0755); err != nil {
			return err
		}

		writeSize, err := writeFile(f, targetFilePath)
		if err != nil {
			return err
		}
		totalSize += writeSize
	}

	return nil
}

func writeFile(f *zip.File, targetFilePath string) (int64, error) {
	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		return 0, err
	}
	defer targetFile.Close()
	file, err := f.Open()
	if err != nil {
		return 0, err
	}
	defer file.Close()

	writeSize, err := io.Copy(targetFile, file)
	if err != nil {
		return 0, err
	}

	return writeSize, nil
}
