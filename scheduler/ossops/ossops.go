package ossops

import "github.com/interesting1113/video-server.git/web/config"

var EP string
var AK string
var SK string

func init() {
	AK = ""
	SK = ""
	EP = config.GetOssAddr()
}

func UploadToOss(filename string, path string, bn string) bool {}

func DeleteObject(filename string, bn string) bool {}
