package models

import (
	"fmt"
	"mlt-go/module/users/models"
	"unsafe"
)

type SliceMock struct {
	Addr uintptr
	Len  int
	Cap  int
}

func main() {
	users := new(models.Users)
	users.Id = 1
	fmt.Println(users)
	//buf := &bytes.Buffer{}
	//err = binary.Write(buf, binary.BigEndian, users)
	//if err != nil {
	//	panic(err)
	//}
	Len := unsafe.Sizeof(*users)
	testBytes := &SliceMock{
		Addr: uintptr(unsafe.Pointer(users)),
		Cap:  int(Len),
		Len:  int(Len),
	}
	//struct 转  []byte
	data := *(*[]byte)(unsafe.Pointer(testBytes))
	//byte 转  struct
	var ptestStruct *models.Users = *(**models.Users)(unsafe.Pointer(&data))
	fmt.Println("ptestStruct.data is : ", ptestStruct)
}
