package easyatomics

import (
	"math/rand"
	"sync"
	"testing"
)

func TestUint64OverFlow(t *testing.T) {
	u64 := AtomicUint64{}
	u64.Dec()
	if u64.Get() != 18446744073709551615 {
		t.Errorf("ERR %d", u64.Get())
	}
	u64.Dec()
	if u64.Get() != 18446744073709551614 {
		t.Errorf("ERR %d", u64.Get())
	}
	u64.Inc()
	if u64.Get() != 18446744073709551615 {
		t.Errorf("ERR %d", u64.Get())
	}
	u64.Inc()
	if u64.Get() != 0 {
		t.Errorf("ERR %d", u64.Get())
	}
}

func TestParallelUint64(t *testing.T) {
	start := uint64(rand.Intn(10000) + 55)
	cycles := rand.Intn(10000) + 55
	incUp := uint64(rand.Intn(9999999)) + 55
	u64 := AtomicUint64{}
	u64.Set(start)
	var waitgroup sync.WaitGroup
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			u64.Inc()
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if u64.Get() != uint64(cycles)+start {
		t.Errorf("ERR %d", u64.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			u64.Dec()
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if u64.Get() != start {
		t.Errorf("ERR %d", u64.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			u64.IncBy(incUp)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if u64.Get() != (incUp*uint64(cycles))+start {
		t.Errorf("ERR %d", u64.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			u64.DecBy(incUp)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if u64.Get() != start {
		t.Errorf("ERR %d", u64.Get())
	}
}

func TestInt64OverFlow(t *testing.T) {
	i64 := AtomicInt64{}
	i64.Dec()
	if i64.Get() != -1 {
		t.Errorf("ERR %d", i64.Get())
	}
	i64.Dec()
	if i64.Get() != -2 {
		t.Errorf("ERR %d", i64.Get())
	}
	i64.Inc()
	if i64.Get() != -1 {
		t.Errorf("ERR %d", i64.Get())
	}
	i64.Inc()
	if i64.Get() != 0 {
		t.Errorf("ERR %d", i64.Get())
	}
}

func TestParallelInt64(t *testing.T) {
	start := int64(rand.Intn(10000) + 55)
	cycles := rand.Intn(10000) + 55
	incUp := int64(rand.Intn(9999999)) + 55

	i64 := AtomicInt64{}
	i64.Set(start)
	var waitgroup sync.WaitGroup
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			i64.Inc()
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if i64.Get() != int64(cycles)+start {
		t.Errorf("ERR %d", i64.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			i64.Dec()
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if i64.Get() != start {
		t.Errorf("ERR %d", i64.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			i64.IncBy(incUp)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if i64.Get() != (int64(cycles)*incUp)+start {
		t.Errorf("ERR %d", i64.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			i64.DecBy(incUp)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if i64.Get() != start {
		t.Errorf("ERR %d", i64.Get())
	}
}

func TestUint32OverFlow(t *testing.T) {
	u32 := AtomicUint32{}
	u32.Dec()
	if u32.Get() != 4294967295 {
		t.Errorf("ERR %d", u32.Get())
	}
	u32.Dec()
	if u32.Get() != 4294967294 {
		t.Errorf("ERR %d", u32.Get())
	}
	u32.Inc()
	if u32.Get() != 4294967295 {
		t.Errorf("ERR %d", u32.Get())
	}
	u32.Inc()
	if u32.Get() != 0 {
		t.Errorf("ERR %d", u32.Get())
	}
}

func TestParallelUint32(t *testing.T) {
	start := uint32(rand.Intn(10000) + 55)
	cycles := rand.Intn(10000) + 55
	incUp := uint32(rand.Intn(9999999)) + 55
	u32 := AtomicUint32{}
	u32.Set(start)
	var waitgroup sync.WaitGroup
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			u32.Inc()
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if u32.Get() != uint32(cycles)+start {
		t.Errorf("ERR %d", u32.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			u32.Dec()
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if u32.Get() != start {
		t.Errorf("ERR %d", u32.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			u32.IncBy(incUp)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if u32.Get() != (incUp*uint32(cycles))+start {
		t.Errorf("ERR %d", u32.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			u32.DecBy(incUp)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if u32.Get() != start {
		t.Errorf("ERR %d", u32.Get())
	}
}

func TestInt32OverFlow(t *testing.T) {
	i32 := AtomicInt32{}
	i32.Dec()
	if i32.Get() != -1 {
		t.Errorf("ERR %d", i32.Get())
	}
	i32.Dec()
	if i32.Get() != -2 {
		t.Errorf("ERR %d", i32.Get())
	}
	i32.Inc()
	if i32.Get() != -1 {
		t.Errorf("ERR %d", i32.Get())
	}
	i32.Inc()
	if i32.Get() != 0 {
		t.Errorf("ERR %d", i32.Get())
	}
}

func TestParallelInt32(t *testing.T) {
	start := int32(rand.Intn(10000) + 55)
	cycles := rand.Intn(10000) + 55
	incUp := int32(rand.Intn(9999999)) + 55

	i32 := AtomicInt32{}
	i32.Set(start)
	var waitgroup sync.WaitGroup
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			i32.Inc()
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if i32.Get() != int32(cycles)+start {
		t.Errorf("ERR %d", i32.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			i32.Dec()
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if i32.Get() != start {
		t.Errorf("ERR %d", i32.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			i32.IncBy(incUp)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if i32.Get() != (int32(cycles)*incUp)+start {
		t.Errorf("ERR %d", i32.Get())
	}
	for i := 0; i < cycles; i++ {
		waitgroup.Add(1)
		go func() {
			i32.DecBy(incUp)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	if i32.Get() != start {
		t.Errorf("ERR %d", i32.Get())
	}
}
