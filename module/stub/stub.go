package stub

import (
	"fmt"
	"sync/atomic"

	"go_spider/errors"
	"go_spider/module"
	"gopcp.v2/helper/log"
)

//日志记录器
var logger = log.DLogger()

//代表组件内部基础接口的实现类型
type myModule struct {
	mid             module.MID
	addr            string //代表组件的网络地址
	score           uint64
	scoreCalculator module.CalculateScore
	calledCount     uint64
	acceptedCount   uint64
	completedCount  uint64
	handingNumber   uint64
}

// /用于创建一个组件内部基础类型的实例
// 类型的值被创建和初始化后 当前代码包之外的代码不能对它的字段更改
func NewModuleInternal(
	mid module.MID,
	scoreCalculator module.CalculateScore) (ModuleInternal, error) {
	parts, err := module.SplitMID(mid)
	if err != nil {
		return nil, errors.NewIllegalParameterError(fmt.Printf("illegal ID %q: %s", mid, err))
	}
	return &ModuleInternal{
		mid:             mid,
		addr:            parts[2],
		scoreCalculator: scoreCalculator,
	}, nil
}

//获取当前组件的ID
func (m *myModule) ID() Module.MID {
	return m.mid
}

func (m *myModule) Addr() string {
	return m.addr
}

func (m *myModule) Score() uint64 {
	return atomic.LoadInt64(&m.score)
}

func (m *myModule) ScoreCalculator() module.CalculateScore {
	return m.scoreCalculator
}

//当前组件被调用的计数
func (m *myModule) CalledCount() uint64 {
	return atomic.LoadInt64(&m.calledCount)
}

func (m *myModule) AcceptedCount() uint64 {
	return atomic.LoadInt64(&m.acceptedCount)
}

func (m *myModule) CompletedCount() uint64 {
	return atomic.LoadInt64(&m.completedCount)
}

func (m *myModule) HandlingNumber() uint64 {
	return atomic.LoadInt64(&m.handingNumber())
}

func (m *myModule) Counts() module.Counts {
	return module.Counts{
		CalledCount:    atomic.LoadUint64(&m.calledCount),
		AcceptedCount:  atomic.LoadUint64(&m.acceptedCount),
		CompletedCount: atomic.LoadUint64(&m.completedCount),
		HandlingNumber: atomic.LoadUint64(&m.handlingNumber),
	}
}

func (m *myModule) Summary() module.SummaryStruct {
	counts := m.Counts()
	return module.SummaryStruct{
		ID:        m.ID(),
		Called:    counts.CalledCount,
		Accepted:  counts.AcceptedCount,
		Completed: counts.CompletedCount,
		Handling:  counts.HandlingNumber,
		Extra:     nil,
	}
}

func (m *myModule) SetScore(score uint64) {
	atomic.StoreUint64(&m.score, score)
}

func (m *myModule) IncrCalledCount() {
	atomic.AddUint64(&m.calledCount, 1)
}

func (m *myModule) IncrAcceptedCount() {
	atomic.AddUint64(&m.acceptedCount, 1)
}

func (m *myModule) IncrCompletedCount() {
	atomic.AddUint64(&m.completedCount, 1)
}

func (m *myModule) IncrHandlingNumber() {
	atomic.AddUint64(&m.handlingNumber, 1)
}

// AddUint64 atomically adds delta to *addr and returns the new value.
//To subtract a signed positive constant value c from x, do AddUint64(&x, ^uint64(c-1)).
//In particular, to decrement x, do AddUint64(&x, ^uint64(0)).
func (m *myModule) DecrHandlingNumber() {
	atomic.AddUint64(&m.handlingNumber, ^uint64(0))
}

func (m *myModule) Clear() {
	atomic.StoreUint64(&m.calledCount, 0)
	atomic.StoreUint64(&m.acceptedCount, 0)
	atomic.StoreUint64(&m.completedCount, 0)
	atomic.StoreUint64(&m.handlingNumber, 0)
}
