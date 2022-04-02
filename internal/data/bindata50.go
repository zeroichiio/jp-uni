package data

import(
	"os"
	"time"
)


func dictUnidictByBytes() ([]byte, error) {
	return _dictUnidictBy, nil
}

func dictUnidictBy() (*asset, error) {
	bytes, err := dictUnidictByBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.by", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
