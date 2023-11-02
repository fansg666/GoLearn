package main

import (
	"fmt"
	"reflect"
)

// import (
//
//	"fmt"
//	"math/rand"
//	"time"
//
// )
//
// type kelvin float64
//
//	func measureTemperature(samples int, sensor func() kelvin) {
//		for i := 0; i < samples; i++ {
//			k := sensor()
//			fmt.Printf("%v K\n", k)
//			time.Sleep(time.Second)
//		}
//	}
//
//	func fakeSensor() kelvin {
//		return kelvin(rand.Intn(151) + 150)
//	}
//
//	func main() {
//		measureTemperature(3, fakeSensor)
//	}
//
//	func main() {
//		plans := [...]int{
//			1,
//			2,
//			3,
//		}
//		for i := 0; i < 3; i++ {
//			fmt.Println(plans[i])
//		}
//	}

//	func main() {
//		tmmp := []float64{
//			-1.0, 2, -3.0, -4.2,
//		}
//		groups := make(map[float64][]float64)
//		for _, t := range tmmp {
//			g := math.Trunc(t/10) * 10
//			groups[g] = append(groups[g], t)
//		}
//		fmt.Println(groups)
//	}
//
//	type suu struct {
//		s int
//		g float64
//	}
//
// func (c suu) deli() float64 {
//
// }
//
//	func main() {
//		var monster []map[string]string
//		monster = make([]map[string]string, 2)
//		monster[0] = make(map[string]string, 2)
//		monster[1] = make(map[string]string, 2)
//		monster[0]["sss"] = "111"
//		monster[0]["ddd"] = "222"
//		monster[1]["ssss"] = "111"
//		monster[1]["ddds"] = "222"
//		fmt.Println(monster)
//	}
//
//	type Student struct {
//		name string
//		id   int
//	}
//
//	type teacher struct {
//		Student
//		score int
//	}
//
//	func main() {
//		teachers := teacher{
//			Student{"fan", 1},
//			1,
//		}
//		fmt.Println(teachers)
//
//	type fan interface {
//		run()
//	}
//
//	type student struct {
//		name string
//	}
//
//	type teacher struct {
//		name1 string
//	}
//
// type gunju struct {
// }
//
//	func (s student) run() {
//		fmt.Println(s.name)
//	}
//
//	func (s student) fly() {
//		fmt.Println("我镇在非")
//	}
//
//	func (t teacher) run() {
//		fmt.Println(t.name1)
//	}
//
//	func (g gunju) runn(fans fan) {
//		fans.run()
//		if stu, ok := fans.(student); ok {
//			stu.fly()
//		}
//
// }
//
//	func main() {
//		s := student{"fan"}
//		t := teacher{"liu"}
//		var fans [2]fan
//		fans[0] = s
//		fans[1] = t
//		var gun gunju
//		for _, v := range fans {
//			gun.runn(v)
//		}
//	}
//
// var (
//
//	myMap = make(map[int]int, 10)
//	lock  sync.Mutex
//
// )
//
//	func test(n int) {
//		res := 1
//		for i := 1; i <= n; i++ {
//			res *= i
//		}
//		lock.Lock()
//		myMap[n] = res
//		lock.Unlock()
//	}
//
//	func main() {
//		for i := 0; i <= 5; i++ {
//			go test(i)
//		}
//		time.Sleep(time.Second * 10)
//		lock.Lock()
//		for i, v := range myMap {
//			fmt.Printf("map[%d]=%d\n", i, v)
//		}
//		lock.Unlock()
//	}
//
//	func main() {
//		var intChant chan int
//		intChant = make(chan int, 3)
//		num := 211
//		intChant <- num
//		fmt.Println(len(intChant), cap(intChant))
//	}
//
//	func writeData(intChan chan int) {
//		for i := 1; i <= 50; i++ {
//			//放入数据
//			intChan <- i
//			fmt.Println("writeData ", i)
//			//time.Sleep(time.Second)
//		}
//		close(intChan) //关闭
//	}
//
//	func readData(intChan chan int, exitChan chan bool) {
//		for {
//			v, ok := <-intChan
//			if !ok {
//				break
//			}
//			//time.Sleep(time.Second)
//			fmt.Printf("readData 读到数据=%v\n", v)
//		}
//		//readData 读取完数据后，即任务完成
//		exitChan <- true
//		close(exitChan)
//	}
//
//	func main() {
//		//创建两个管道
//		intChan := make(chan int, 50)
//		exitChan := make(chan bool, 1)
//		go writeData(intChan)
//		go readData(intChan, exitChan)
//		//time.Sleep(time.Second * 10)
//		for {
//			_, ok := <-exitChan
//			if !ok {
//				break
//			}
//		}
//	}
func reFlestTest01(b interface{}) {
	//rTyp := reflect.TypeOf(b)
	//fmt.Println(rTyp)
	num1 := reflect.ValueOf(b)
	//num1.SetInt(66)
	fmt.Printf("%v", num1.Kind())
	//num2 := num1.Int() + 1
	//fmt.Printf("%v\n", num1.Int())
	//fmt.Println(num2)
	//iv := num1.Interface()
	//num3 := iv.(int)
	//fmt.Println(num3)
}
func main() {
	var num int = 100
	reFlestTest01(num)
}
