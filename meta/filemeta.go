package meta

import (
	"go_net_disk/dao"
	"sort"
)

// FileMeta 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMeats map[string]FileMeta

func init() {
	fileMeats = make(map[string]FileMeta)
}

// UpdateFileMeta 新增/更新文件元信息
func UpdateFileMeta(fileMeta FileMeta) {
	fileMeats[fileMeta.FileSha1] = fileMeta
}

// UpdateFileMetaDB 新增更新文件元信息到数据库中
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return dao.OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMeats[fileSha1]
}

// GetFileMetaDB 从MySQL获取元信息
func GetFileMetaDB(fileSha1 string) (FileMeta, error) {
	tfile, err := dao.GetFileMeta(fileSha1)
	if err != nil {
		return FileMeta{}, err
	}
	fmeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return fmeta, nil
}

// GetLastFileMetas 获取批量的文件元信息列表
func GetLastFileMetas(count int) []FileMeta {
	fMetaArray := make([]FileMeta, len(fileMeats))
	for _, v := range fileMeats {
		fMetaArray = append(fMetaArray, v)
	}
	sort.Sort(ByUploadTime(fMetaArray))
	return fMetaArray[0:count]
}

// RemoveFileMeta 删除元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMeats, fileSha1)
}
