package data

import(
	"os"
	"time"
)


func dictUnidictBkBytes() ([]byte, error) {
	return _dictUnidictBk, nil
}

func dictUnidictBk() (*asset, error) {
	bytes, err := dictUnidictBkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bk", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
