package main

import "fmt"

type SystemMetrics struct {
	CPUUsage    int
	MemoryUsage int
	DiskIO      int
}

func (s *SystemMetrics) ChangeCPU(value int) {
	s.CPUUsage = value
	fmt.Printf("address in func1:%p\n", s)
}
func (s *SystemMetrics) ChangeMemory(value int) {
	s.MemoryUsage = value
	fmt.Printf("address in func2:%p\n", s)
}
func (s *SystemMetrics) ChangeDisk(value int) {
	s.DiskIO = value
	fmt.Printf("address in func3:%p\n", s)
}

func main() {
	 sys := SystemMetrics{
		CPUUsage:    12,
		MemoryUsage: 20,
		DiskIO:      89,
	}
	fmt.Printf("main address: %p\n",&sys)
	fmt.Println("---------------------------------------")
	sys.ChangeCPU(23)
	sys.ChangeMemory(10)
	sys.ChangeDisk(50)
	fmt.Println("---------------------------------------")
	fmt.Println("values after changing :",sys)

}
