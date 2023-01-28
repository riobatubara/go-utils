package uuid

import (
	"time"
)

type UUID [16]byte

type UUIDBuilder struct {
	Timestamp *int64
}

func New() UUIDBuilder {
	now := time.Now().Unix()
	return UUIDBuilder{Timestamp: &now}
}

// func New(args ...interface{}) ([]byte, error) {
// fmt.Println(len(args))
// var stringbased []string
// for _, m := range args {
// 	for idx, val := range m.(map[string]interface{}) {
// 		stringbased = append(stringbased, idx+":"+fmt.Sprintf("%v", val))
// 	}
// }
// jn := strings.Join(stringbased, delim)

// hash := md5.New()
// _, err := hash.Write([]byte(jn))
// if err != nil {
// 	return nil, err
// }
// sum := hash.Sum(nil)

// uuid := make([]byte, 16)
// copy(uuid, sum)
// fmt.Println(jn)
// now := time.Now().Unix()
// fmt.Println("timestamp: ", now)
// namebased :=
// return uuid, nil
// return nil, nil
// }
