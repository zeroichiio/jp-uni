package data

import(
	"os"
	"time"
)


func dictUnidictCrBytes() ([]byte, error) {
	return _dictUnidictCr, nil
}

func dictUnidictCr() (*asset, error) {
	bytes, err := dictUnidictCrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.cr", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
