package schema

import "github.com/liukeshao/echo-admin/pkg/uid"

func IdGenerator() uint64 {
	id, err := uid.GenUint64()
	if err != nil {
		panic(err)
	}
	return id
}
