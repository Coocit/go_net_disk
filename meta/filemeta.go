package meta

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

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMeats[fileSha1]
}
