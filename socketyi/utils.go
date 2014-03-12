package socketyi

import (
	"github.com/golang/glog"
	"fmt"
    "time"
    "math/rand"
    "crypto/md5"
    "strconv"
    "io"
)

func CheckError(obj interface{}, err error) {
    if err != nil {
        glog.V(0).Infoln(obj, "Fatal error: %s", err.Error())
    }
}

func CreatSessionID() string {
	nano := time.Now().UnixNano()
    rand.Seed(nano)
    rndNum := rand.Int63()
    return Md5(Md5(strconv.FormatInt(nano, 10))+Md5(strconv.FormatInt(rndNum, 10)))
}

func Md5(text string) string {
	hashMd5 := md5.New()
    io.WriteString(hashMd5, text)
    return fmt.Sprintf("%x", hashMd5.Sum(nil))
}
