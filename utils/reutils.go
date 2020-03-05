package utils
import(
	"crypto/md5"
	"fmt"
	"time"
)
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2019-10-24 15:04:05")
}