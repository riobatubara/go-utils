package uuid

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"strings"
)

const delim = "||"

func (this UUIDBuilder) Stringbased(args ...interface{}) UUID {
	if len(args) > 0 {
		var stringbased []string
		for _, m := range args {
			for idx, val := range m.(map[string]interface{}) {
				stringbased = append(stringbased, idx+":"+fmt.Sprintf("%v", val))
			}
		}
		jn := strings.Join(stringbased, delim)

		hash := md5.New()
		hash.Write([]byte(jn))
		sum := hash.Sum(nil)

		uuid := New().Timebased()
		copy(uuid[:], sum)
		return uuid
	} else {
		var uuid UUID
		io.ReadFull(rand.Reader, uuid[:])

		uuid[6] &= 0x0F // clear version
		uuid[6] |= 0x40 // set version to 4 (random uuid)
		uuid[8] &= 0x3F // clear variant
		uuid[8] |= 0x80 // set to IETF variant

		return uuid
	}
}
