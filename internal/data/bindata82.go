package data

import(
	"os"
	"time"
)


func dictUnidictDeBytes() ([]byte, error) {
	return _dictUnidictDe, nil
}

func dictUnidictDe() (*asset, error) {
	bytes, err := dictUnidictDeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/unidict.de", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1609246449, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
