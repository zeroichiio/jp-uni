package data

import(
	"os"
	"time"
)


func dictUnidictAjBytes() ([]byte, error) {
	return _dictUnidictAj, nil
}

func dictUnidictAj() (*asset, error) {
	bytes, err := dictUnidictAjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.aj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
