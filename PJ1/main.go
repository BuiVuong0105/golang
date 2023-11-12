package main

import (
	"family/pattern"
	"fmt"
)

func main() {
	// ifEsleSwitchCase()
	dependency()
	// decorator()
	// anynomousStruct()
	// handleStruct()
	// handleVariable()
	// handleArray()
	// handleSlice()
	// handleMap()
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

// func handleStruct() {
// 	var studentVuong = service.Student{
// 		Id:      1,
// 		Name:    "Vuong",
// 		Subject: []string{"Ly"},
// 		Address: service.Address{Location: "aaaa"},
// 		Point:   100,
// 	}
// 	fmt.Printf("STUDENT VUONG: %v, %T\n", studentVuong, studentVuong)
// 	var studentNam = &studentVuong // sử dụng con trỏ studentNam trỏ đến studentVuong => studentdentNam có thể thay đổi studentVuong

// 	fmt.Printf("STUDENT Nam: %v, %T\n", *studentNam, *studentNam)
// 	studentNam.Name = "Nam"
// 	fmt.Printf("STUDENT VUONG: %v, %T\n", studentVuong, studentVuong)
// 	fmt.Printf("STUDENT Nam: %v, %T\n", *studentNam, *studentNam)

// 	var typeOfStudent reflect.Type = reflect.TypeOf(studentVuong)
// 	field, _ := typeOfStudent.FieldByName("Point")
// 	fmt.Printf("typeOfStudent: %v, %T\n", typeOfStudent, typeOfStudent)
// 	fmt.Printf("Anotaiton Point: %v", field)
// }

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

// func handleMap() {
// 	// MAP

// 	// Khởi tạo các  giá trị mặc định
// 	// var mapStudent map[int64]string = map[int64]string{
// 	// 	1: "Lan",
// 	// 	2: "Hung",
// 	// }
// 	// Cách khởi tạo map rỗng
// 	// mapStudent = make(map[int64]string)
// 	// mapStudent = map[int64]string{}

// 	// Khởi tạo map với key là mảng
// 	// var mapStudent2 = map[[3]int64]string{} // Không thể khởi tạo map với key là slice , phải sử dụng mảng
// 	// fmt.Printf("MAP STUDENT2: %v, %T\n", mapStudent2, mapStudent2)
// 	var mapStudent map[int64]string = map[int64]string{
// 		1: "Lan",
// 		2: "Hung",
// 	}
// 	mapStudent[2] = "Lam"
// 	mapStudent[3] = "Tuan"
// 	fmt.Printf("MAP STUDENT: %v, %T\n", mapStudent, mapStudent)
// 	delete(mapStudent, 1)
// 	_, contain := mapStudent[1]
// 	fmt.Printf("MAP STUDENT 1: %v, %T\n", mapStudent[1], mapStudent[1])
// 	fmt.Printf("KEY 1: %v", contain)

// 	mapCoppyStudent := mapStudent // Như 1 con trỏ, trỏ đến map origin
// 	fmt.Printf("MapCoppy: %v", mapCoppyStudent)
// 	fmt.Printf("MapCoppy: %v", mapStudent)
// }
