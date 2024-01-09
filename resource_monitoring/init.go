package resource_monitoring

var lastCurrentBytesSent int32
var lastCurrentBytesRecv int32

func init(){
	lastCurrentBytesSent = 0
	lastCurrentBytesRecv = 0
}