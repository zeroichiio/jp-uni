package data

import(
	"os"
	"time"
)


func dictUnidictBfBytes() ([]byte, error) {
	return _dictUnidictBf, nil
}

func dictUnidictBf() (*asset, error) {
	bytes, err := dictUnidictBfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.bf", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
