package other

import "fmt"

type Type int

const (
	Null Type = iota

	False
	Number
	String
	True

	JSON
)
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

func iotas() {

	fmt.Printf("Ldate:%d Ltime: %d Lmicroseconds: %d LstdFlags: %d", Ldate, Ltime, Lmicroseconds, LstdFlags)
	//a := uint16(1 << 15)
	//fmt.Printf("2进制 %016b, 左移8:%016b", a, a>>8)
	// a := []byte("ab")
	// println("len:", len(a), "byte 1:", a[0], "byte 2:", a[1])
	// fmt.Printf("a:%016b \n", a)
	// fmt.Printf(" bigEndian:%016b a", binary.BigEndian.Uint16([]byte(a)))

	// beBuf := make([]byte, 4)
	// frameType := 123456789
	// fmt.Printf("a:%b %b %b\n", Ldate, Ltime, LstdFlags)
	// binary.LittleEndian.PutUint32(beBuf, uint32(frameType))
	// fmt.Println(binary.BigEndian.Uint32(beBuf))
	// fmt.Println(binary.LittleEndian.Uint32(beBuf))

	// var a []byte = []byte{0, 1, 2, 3}
	// fmt.Println(a)
	// fmt.Println(binary.BigEndian.Uint32(a))
	// fmt.Println(binary.LittleEndian.Uint32(a))

	// fmt.Println(Null, False, Number, String, True, JSON)
	// files, err := ioutil.ReadDir(".")
	// if err != nil {
	//	log.Fatal(err)
	// }
	// for _, file := range files {
	//	if file.IsDir() {
	//		continue
	//	}
	//	fmt.Println(file.Name())
	//	if file.Name() == "issue.go" {
	//		content, err := ioutil.ReadFile(file.Name())
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		fmt.Printf("File contents: %s", content)
	//	}
	// }
}
