package main

import (
	"family/pattern"
	"family/service"
	"fmt"
	"reflect"
)

type BaseLocation struct {
	city     string
	district string
	ward     string
	address  string
}

type School struct {
	BaseLocation
	name string
}

func handleExtend() {
	school := &School{}
	school.city = "Ha Noi"
	school.district = "Hoai Duc"
	school.ward = "An Khanh"
	school.address = "Van Lung"
	school.name = "An Khanh A"
	fmt.Println("School Information: ", *school)
	fmt.Println("School BaseLocation: ", school.BaseLocation)
}

type Position struct {
	lat float64
	lon float64
}

func (p *Position) SpecialMove(x, y float64) {
	p.lat = x + x
	p.lon = y + y
}

// định nghĩa hàm Move trong object Position
// p Position: mỗi lần gọi sẽ copy 1 đối tượng mới vào thay đổi trên đối tượng copy sẽ không thay đổi trên đối tượng gốc
// p *Position: mỗi lần gọi sẽ gọi trên con trỏ trỏ vào đối tượng gốc
func (p *Position) Move(lat, lon float64) {
	p.lat += lat
	p.lon += lon
}

// định nghĩa hàm Teleport trong object Position
func (p *Position) Teleport(lat, lon float64) {
	p.lat = lat
	p.lon = lon
}

type SpecialPosition struct {
	*Position
}

// Override SpecialMove của Position
func (p *SpecialPosition) SpecialMove(x, y float64) {
	p.lat = x * x
	p.lon = y * y
}

// Tạo struct player kể thừa thuộc tính từ Position và có 1 biến position kiểu con trỏ
type Player struct {
	*Position
}

type Enmery struct {
	*SpecialPosition
}

func createPlayer() *Player {
	return &Player{
		Position: &Position{
			lat: 6,
			lon: 8,
		},
	}
}

func createEnmery() *Enmery {
	return &Enmery{
		SpecialPosition: &SpecialPosition{
			Position: &Position{
				lat: 6,
				lon: 6,
			},
		},
	}
}

func handleComposition() {

	player := createPlayer()
	fmt.Println("Player: ", *player.Position)
	player.Move(5, 5)
	fmt.Println("Player Move: ", *player.Position)
	player.Teleport(5, 5)
	fmt.Println("Player Teleport: ", *player.Position)

	enmery := createEnmery()
	fmt.Println("Enmery: ", *enmery.Position)
	enmery.SpecialMove(5, 5)
	fmt.Println("Enmery SpecialMove: ", *enmery.Position)
}

func main() {

	// a := []int{1, 2, 3} // slice a trong đó chứa con trỏ trỏ đến phần tử đầu của mảng
	// b := a              // khởi tạo slice b chứa con trỏ trỏ đến mảng của slice a
	// c := *&a[2] + 5
	// fmt.Println(&a[1], &b[2], c) // các ô nhớ trong mảng xếp gần nhau

	m := map[string]string{"vuong": "MALE"}
	m1 := m
	fmt.Println(&m, &m1)
	m["vuong"] = "male"
	fmt.Println(m, m1)

	// position := new(Position)
	// fmt.Println(position)
	// (*position).lat = 5
	// fmt.Println((*position).lat)

	// handleComposition()
	// handleExtend()
	// anynomousStruct()
	// handleStruct()
	// handleVariable()
	// handleArray()
	// handleSlice()
	// handleMap()
	// deferEx()
}

// Thưc hiện lệnh sau defer trước khi hàm return (đưa lệnh vào stack)
// Panic là từ khóa để ném ra lỗi, tương tự throw trong java
// sử dụng defer và panic gần giống với việc sử dụng final try catch trong java
func deferEx() {

	//call Ham ần danh
	defer func() {
		err := recover() // Kiểm tra có lỗi trong hệ thống không , lỗi xảy ra từ pannic có thì tổng hợp vào biến err
		if err != nil {
			println("Loi la: ", err)
		}
	}()
	panic("Lỗi chia cho 0") // Ném lỗi này ra khỏi function, hàm nay tương đương với throw exception trong java

	// a := 10
	// b := 0
	// c := a / b
	// println("C: ", c)

	// value := 5
	// defer changeValue(&value) // Goi hàm này cuối cùng
	// value = value + 1
	// println("Value: ", value)

	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// fmt.Println("3")
}

func changeValue(value *int) {
	println("Value Prepare Change: ", *value)
	*value = *value + 1
}

func loop() {
	for i := 1; i <= 10; i++ {
		println("value: ", i)
	}

	m := map[string]int{"vuong": 30, "toan": 22}
	for _, value := range m {
		fmt.Println(value)
	}
}

// Truyền tham trị ( truyền giá trị vào sv, sv thay đổi thì param truyền vào ko đổi)
func printSv(sv service.Student) {
	sv.Name = "AAAA"
	fmt.Printf("SV: %v \n", sv)
}

func dependency() {
	rc := pattern.CreateRockClimber(&pattern.NOPafetyPlacer{})
	for i := 1; i < 11; i++ {
		rc.ClimbRock()
	}
}

func decorator() {
	s := &pattern.Store{}
	// http.HandleFunc("/", pattern.MakeHttpFunc(s, pattern.Handler))
	pattern.Execute(pattern.MyExecuteFunction(s))
}

func ifEsleSwitchCase() {
	var i interface{} = 5.8
	switch i.(type) {
	case int:
		fmt.Println("Type int")
		break

	case float32:
		fmt.Println("Type float32")
		break

	case float64:
		fmt.Println("Type float64")
		break
	}
}

// func increment(a *int) {
// 	*a++
// }

// func anynomousStruct() {
// 	// studentAnynomouse := struct{ name string }{} // Cach 1
// 	var studentAnynomouse struct{ name string } = struct{ name string }{}   // cach 2
// 	var structAnynomousePointer *struct{ name string } = &studentAnynomouse // khoi tao con tro
// 	fmt.Printf("studentAnynomouse: %v, %T\n", studentAnynomouse, studentAnynomouse)
// 	fmt.Printf("studentAnynomouse: %v, %T\n", structAnynomousePointer, structAnynomousePointer)
// }

func handleStruct() {
	var studentVuong = service.Student{
		Id:      1,
		Name:    "Vuong",
		Subject: []string{"Ly"},
		Address: service.Address{Location: "aaaa"},
		Point:   100,
	}
	fmt.Printf("STUDENT VUONG: %v, %T\n", studentVuong, studentVuong)
	var studentNam = &studentVuong // sử dụng con trỏ studentNam trỏ đến studentVuong => studentdentNam có thể thay đổi studentVuong

	fmt.Printf("STUDENT Nam: %v, %T\n", *studentNam, *studentNam)
	studentNam.Name = "Nam"
	fmt.Printf("STUDENT VUONG: %v, %T\n", studentVuong, studentVuong)
	fmt.Printf("STUDENT Nam: %v, %T\n", *studentNam, *studentNam)

	var typeOfStudent reflect.Type = reflect.TypeOf(studentVuong)
	field, _ := typeOfStudent.FieldByName("Point")
	fmt.Printf("typeOfStudent: %v, %T\n", typeOfStudent, typeOfStudent)
	fmt.Printf("Anotaiton Point: %v", field)
}

// func handleVariable() {
// 	var j float64 = 34.4
// 	fmt.Println("Hellow World")
// 	var i int = int(j)
// 	var c string = strconv.Itoa(i)
// 	fmt.Println(i, c)
// }

// func handleArray() {
// 	//ARRAY
// 	points := [3]int{1, 2, 4}      // Khai bao mang
// 	var pointsCopy [3]int = points // Taọ ra mảng mới, copy dữ liệu từ mảng points sang pointCoppy. Thay đổi ở pointCoppy thì points không đổi
// 	pointsCopy[1] = 20
// 	var pointsRef *[3]int = &points // Tạo ra con trỏ mảng trỏ đến địa chỉ của points
// 	pointsRef[1] = 60

// 	fmt.Printf("POINT: %v, %T\n", points, points)
// 	fmt.Printf("POINT COPY: %v, %T\n", pointsCopy, pointsCopy)
// 	fmt.Printf("POINT REF: %v, %T\n", pointsRef, pointsRef)

// 	colors := [...]string{"blue", "red"}
// 	fmt.Printf("COLOR: %v, %T\n", colors, colors)

// 	var persons [3]string
// 	persons[0] = "vuong"
// 	persons[1] = "hung"
// 	fmt.Printf("PERSION: %v, %T\n", persons, persons)
// }

// func handleSlice() {
// 	// SLICE
// 	//sự khác biệt giữa khởi tạo slice và mảng là khởi tạo slice không chỉ định tham số number trong []
// 	// Slice như 1 con trỏ , trỏ đến 1 mảng. add phần tử vào slice chính là add thêm phẩn tử vào mảng
// 	//Hãy nhớ rằng, nếu bạn khởi tạo giá trị bên trong toán tử [ ], nghĩa là bạn đang tạo một mảng. Nếu bạn không chỉ rõ giá trị, nghĩa là bạn đang tạo một slice.

// 	var slicePoints []int = make([]int, 2) //
// 	// slicePoints := []int{1, 2, 4} // không cần chỉ định tham số là độ dài của mảng
// 	slicePoints[1] = 7
// 	fmt.Printf("SLICE POINT: %v, %v %v, %T\n", len(slicePoints), cap(slicePoints), slicePoints, slicePoints)
// 	slicePoints = append(slicePoints, 4)
// 	slicePoints = append(slicePoints, 5)
// 	slicePoints = append(slicePoints, 6)
// 	slicePoints = append(slicePoints, []int{6, 7, 8}...)
// 	fmt.Printf("SLICE POINT: %v, %v %v, %T\n", len(slicePoints), cap(slicePoints), slicePoints, slicePoints)

// 	slicePointRef := slicePoints[:2]
// 	slicePointRef[0] = 1000
// 	fmt.Printf("SLICE POINT: %v, %v %v, %T\n", len(slicePoints), cap(slicePoints), slicePoints, slicePoints)
// 	fmt.Printf("SLICE POINT REF: %v, %v %v, %T\n", len(slicePointRef), cap(slicePointRef), slicePointRef, slicePointRef)

// 	slicePoints = append(slicePoints, 1001)
// 	slicePointRef = append(slicePointRef, 1002)

// 	fmt.Printf("SLICE POINT: %v, %v %v, %T\n", len(slicePoints), cap(slicePoints), slicePoints, slicePoints)
// 	fmt.Printf("SLICE POINT REF: %v, %v %v, %T\n", len(slicePointRef), cap(slicePointRef), slicePointRef, slicePointRef)
// }

func handleMap() {
	// 	// MAP

	// 	// Khởi tạo các  giá trị mặc định
	// 	// var mapStudent map[int64]string = map[int64]string{
	// 	// 	1: "Lan",
	// 	// 	2: "Hung",
	// 	// }
	// 	// Cách khởi tạo map rỗng
	// 	// mapStudent = make(map[int64]string)
	// 	// mapStudent = map[int64]string{}

	// Khởi tạo map với key là mảng
	// var mapStudent2 = map[[3]int64]string{} // Không thể khởi tạo map với key là slice , phải sử dụng mảng
	// fmt.Printf("MAP STUDENT2: %v, %T\n", mapStudent2, mapStudent2)
	var mapStudent map[int64]string = map[int64]string{
		1: "Lan",
		2: "Hung",
	}
	mapStudent[2] = "Lam"
	mapStudent[3] = "Tuan"
	fmt.Printf("MAP STUDENT: %v, %T\n", mapStudent, mapStudent)
	delete(mapStudent, 1)
	_, contain := mapStudent[1]
	fmt.Printf("MAP STUDENT 1: %v, %T\n", mapStudent[1], mapStudent[1])
	fmt.Printf("KEY 1: %v", contain)

	mapCoppyStudent := mapStudent // Như 1 con trỏ, trỏ đến map origin
	fmt.Printf("MapCoppy: %v", mapCoppyStudent)
	fmt.Printf("MapCoppy: %v", mapStudent)
}
