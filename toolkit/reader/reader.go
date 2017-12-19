package reader

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

type MultipleReader interface {
	//多重读取器的接口 后者会持有多重读取其中的数据
	Reader() io.ReadCloser
}

//多重读取器的实现类型
type myMultipleReader struct {
	data []byte
}

func (m *myMultipleReader) Reader() io.ReadCloser {
	// 通过 []byte 创建一个 bytes.Reader 对象
	// NopCloser 将 r 包装为一个 ReadCloser 类型，但 Close 方法不做任何事情。
	return ioutil.NopCloser(bytes.NewReader(m.data))
}

func NewMultipleReader(reader io.Reader) (MultipleReader, error) {
	var data []byte
	var err error
	if reader != nil {
		data, err = ioutil.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("multiple reader: couldn't create a new one: %s", err)
		}
	} else {
		data = []byte{}
	}
	return &myMultipleReader{
		data: data,
	}, nil
}
