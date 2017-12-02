package main

import "fmt"

type ixxx int
type xsdf string
type sdffj float32

type PCer interface { //1.接口命名习惯以er结尾
	GetBrand() string
	Memoryer //2.接口可以嵌入接口
	Cpuer    //3.接口只是方法集，不含数据字段,interface中只有方法，没有数据字段
	PrintInfo()
}

//内存条
type Memoryer interface {
	GetMemory() int
}

//CPU
type Cpuer interface {
	GetCpu() int
}

type HighEndPC struct {
	Brand        string
	MemoryCap    int
	CpuKernlCnt  int
	GpuMemoryCap int
}

type LowhEndPC struct {
	Brand       string
	MemoryCap   int
	CpuKernlCnt int
}

func (self *HighEndPC) GetBrand() string {
	return self.Brand
}

func (self *HighEndPC) GetMemory() int {
	return self.MemoryCap
}

func (self *HighEndPC) GetCpu() int {
	return self.CpuKernlCnt
}

func (self *HighEndPC) GetGpu() int {
	return self.GpuMemoryCap
}

func (self *HighEndPC) PrintInfo() {
	fmt.Printf("高端电脑 品牌:%s,内存大小%d,处理器核心数%d,显存大小%d \n", self.Brand, self.MemoryCap, self.CpuKernlCnt, self.GpuMemoryCap)
}
func (self *LowhEndPC) GetBrand() string {
	return self.Brand
}

func (self *LowhEndPC) GetMemory() int {
	return self.MemoryCap
}

func (self *LowhEndPC) GetCpu() int {
	return self.CpuKernlCnt
}

func (self *LowhEndPC) PrintInfo() {
	fmt.Printf("低端电脑 品牌:%s,内存大小%d,处理器核心数%d \n", self.Brand, self.MemoryCap, self.CpuKernlCnt)
}

func main() {
	companyPC := LowhEndPC{"dell", 8, 2}
	homePC := HighEndPC{"diy", 16, 4, 11}
	myPC := []PCer{&companyPC, &homePC} //4.PCer是companyPC和homePC的抽象，只要实现了接口中的方法，就可以塞进去
	for _, pc := range myPC {
		pc.PrintInfo()
	}
}
