/*
 * Author: fasion
 * Created time: 2019-06-27 15:11:03
 * Last Modified by: fasion
 * Last Modified time: 2019-08-21 17:45:54
 */

package runtime

import (
	"runtime"
	"time"
)

type GoPerf struct {
	Goroutines float64

	GoMemoryVirtual float64

	GoHeapVirtual float64
	GoHeapUsed float64
	GoHeapIdle float64
	GoHeapFree float64

	GoHeapAllocated float64
	GoHeapObjects float64

	GoStackVirtual float64
}

type GoPerfSample struct {
	Time time.Time
	Perf GoPerf
}

type GoPerfSampler struct {
	lastMemStats *runtime.MemStats
	lastMemStatsTime time.Time
}

func NewGoPerfSampler() (*GoPerfSampler) {
	return &GoPerfSampler{}
}

func (self *GoPerfSampler) Sample() (*GoPerfSample, error) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	sample := GoPerfSample{
		Time: time.Now(),
		Perf: GoPerf{
			Goroutines: float64(runtime.NumGoroutine()),
			GoMemoryVirtual: float64(memStats.Sys),
			GoHeapVirtual: float64(memStats.HeapSys),
			GoHeapUsed: float64(memStats.HeapInuse),
			GoHeapIdle: float64(memStats.HeapIdle - memStats.HeapReleased),
			GoHeapFree: float64(memStats.HeapReleased),
			GoHeapAllocated: float64(memStats.HeapAlloc),
			GoHeapObjects: float64(memStats.HeapObjects),
			GoStackVirtual: float64(memStats.StackSys),
		},
	}

	return &sample, nil
}
