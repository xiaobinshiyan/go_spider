package module

import (
	"math"
	"sync"
)

type SNGenertor interface {
	Start() uint64      //用于获取预设的最小序列号
	Max() uint64        //获取预设的最大序列号
	Next() uint64       //用于获取下一个序列号
	CycleCount() uint64 //获取循环计数
	Get() uint64        // Get 用于获得一个序列号并准备下一个序列号。
}
type mySNGenertor struct {
	start      uint64
	max        uint64
	next       uint64
	cycleCount uint64
	lock       sync.RWMutex
}

// NewSNGenertor 会创建一个序列号生成器。
// 参数start用于指定第一个序列号的值。
// 参数max用于指定序列号的最大值。
func NewSNGenertor(start uint64, max uint64) SNGenertor {
	if max == 0 {
		//一个常量 int最大值
		max = math.MaxUint64
	}
	return &mySNGenertor{
		start: start,
		max:   max,
		next:  start,
	}
}

func (m *mySNGenertor) Start() uint64 {
	return m.start
}

func (m *mySNGenertor) Max() {
	return m.max
}
func (m *mySNGenertor) Next() {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.next
}

func (m *mySNGenertor) CycleCount() {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.cycleCount
}

func (m *mySNGenertor) Get() {
	m.lock.Lock()
	defer m.lock.Unlock()
	id := m.next
	if id == m.max {
		m.next = m.start
		m.cycleCount++
	} else {
		m.next++
	}
	return id
}
