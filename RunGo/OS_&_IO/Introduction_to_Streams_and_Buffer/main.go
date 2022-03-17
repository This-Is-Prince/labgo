package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
=================
ALL FUNCTIONS

1). io.Reader interface
2). strings.NewReader returns io.Reader interface
3). ioutil.ReadAll()
4). io.ReadFull()
5). io.ReadAtLeast()
6). io.LimitReader()

7). io.Writer interface
8).

=================
*/

func main() {
	fmt.Println("-------Introduction to Streams and Buffers-------")

	// IO_Reader()
	// Strings_NewReader()
	// IoUtil_ReadAll()
	// Io_ReadFull()
	// IO_ReadAtLeast()
	// Io_LimitReader()
	// Io_Writer()
	// Io_WriteString()
	// Fmt_Writer()
	// Io_Copy()
	// Io_CopyN()
	// Io_Pipe()
	// Bytes_NewBufferString()
	// Bufio_Writer()
	Bufio_Reader()
}

/*
Reading from a Data Source
*/
type DataSource struct {
	str   string
	index int
}

func (src *DataSource) Read(p []byte) (int, error) {
	lenSrcStr := len(src.str)
	if src.index >= lenSrcStr {
		return 0, io.EOF
	}
	n := len(p)
	end := src.index + n
	if end > lenSrcStr {
		end = lenSrcStr
		n = end - src.index
	}
	copy(p, []byte(src.str[src.index:end]))
	src.index = end
	return n, nil
}

func IO_Reader() {
	src := DataSource{str: "Hello Amazing World!"}

	p := make([]byte, 3)
	for {
		n, err := src.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])
		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occurred!", err)
			break
		}
	}
}

func Strings_NewReader() {
	src := strings.NewReader("Hello Amazing World!")
	p := make([]byte, 3)
	for {
		n, err := src.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])
		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occurred!", err)
			break
		}
	}
}

func IoUtil_ReadAll() {
	fmt.Println("+++===== IoUtil_ReadAll =====+++")
	src := strings.NewReader("Hello Amazing World!")
	data, err := ioutil.ReadAll(src)
	fmt.Println(string(data), err)
}

func Io_ReadFull() {
	fmt.Println("+++===== Io_ReadFull =====+++")

	src := strings.NewReader("Hello Amazing World!")
	buf := make([]byte, 14)

	bytesRead1, err1 := io.ReadFull(src, buf)
	fmt.Println("BytesRead1:", bytesRead1, "Err1:", err1, "Data:", string(buf[:bytesRead1]))

	bytesRead2, err2 := io.ReadFull(src, buf)
	fmt.Println("BytesRead2:", bytesRead2, "Err2:", err2, "Data:", string(buf[:bytesRead2]))

	bytesRead3, err3 := io.ReadFull(src, buf)
	fmt.Println("BytesRead3:", bytesRead3, "Err3:", err3, "Data:", string(buf[:bytesRead3]))
}

func IO_ReadAtLeast() {
	fmt.Println("+++===== ReadAtLeast =====+++")

	src := strings.NewReader("Hello Amazing World!")
	buf := make([]byte, 14)

	bytesRead1, err1 := io.ReadAtLeast(src, buf, 1)
	fmt.Println("BytesRead1:", bytesRead1, "Err1:", err1, "Data:", string(buf[:bytesRead1]))

	bytesRead2, err2 := io.ReadAtLeast(src, buf, 1)
	fmt.Println("BytesRead2:", bytesRead2, "Err2:", err2, "Data:", string(buf[:bytesRead2]))

	bytesRead3, err3 := io.ReadAtLeast(src, buf, 1)
	fmt.Println("BytesRead3:", bytesRead3, "Err3:", err3, "Data:", string(buf[:bytesRead3]))
}

func Io_LimitReader() {
	fmt.Println("+++===== Io_LimitReader =====+++")

	mainSrc := strings.NewReader("Hello Amazing World!")

	src := io.LimitReader(mainSrc, 10)

	p := make([]byte, 3)
	for {
		n, err := src.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])
		if err == io.EOF {
			fmt.Println("--end-of-life--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occurred!", err)
			break
		}
	}
}

/*
Writing to a Data store
*/

type DataStore struct {
	data []byte
}

func (ds *DataStore) Write(p []byte) (n int, err error) {
	lenStore := len(ds.data)
	if lenStore >= 10 {
		return 0, errors.New("store if full")
	}
	remCap := 10 - lenStore
	n = len(p)
	if remCap < n {
		n = remCap
	}
	ds.data = append(ds.data, p[:n]...)
	return
}

func Io_Writer() {
	fmt.Println("+++===== Io_Writer =====+++")

	ds := DataStore{}

	bytesWritten1, err1 := ds.Write([]byte("Hello!"))
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten1, err1)
	fmt.Printf("Value of ds.data: %s\n\n", ds.data)

	bytesWritten2, err2 := ds.Write([]byte(" Amazing"))
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten2, err2)
	fmt.Printf("Value of ds.data: %s\n\n", ds.data)

	bytesWritten3, err3 := ds.Write([]byte(" World!"))
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten3, err3)
	fmt.Printf("Value of ds.data: %s\n\n", ds.data)
}

func Io_WriteString() {
	fmt.Println("+++===== Io_WriteString =====+++")

	ds := &DataStore{}

	bytesWritten1, err1 := io.WriteString(ds, "Hello!")
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten1, err1)
	fmt.Printf("Value of ds.data: %s\n\n", ds.data)

	bytesWritten2, err2 := io.WriteString(ds, " Amazing")
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten2, err2)
	fmt.Printf("Value of ds.data: %s\n\n", ds.data)

	bytesWritten3, err3 := io.WriteString(ds, " World!")
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten3, err3)
	fmt.Printf("Value of ds.data: %s\n\n", ds.data)
}

func Fmt_Writer() {
	fmt.Println("+++===== Fmt_Writer =====+++")

	io.WriteString(os.Stdout, "Hello World!\n")

	os.Stdout.Write([]byte("Hello World!\n"))

	fmt.Fprint(os.Stdout, "Hello World!\n")
	fmt.Fprintln(os.Stdout, "Hello World!")
	fmt.Fprintf(os.Stdout, "%s World!\n", "Hello")
}

/*
CLosing I/O Operations
*/

func Io_Copy() {
	fmt.Println("+++===== Io_Copy =====+++")

	stringReader := strings.NewReader("Hello World! How are you?\n")
	io.Copy(os.Stdout, stringReader)
}

func Io_CopyN() {
	fmt.Println("+++===== Io_CopyN =====+++")

	stringReader := strings.NewReader("Hello World! How are you?\n")
	io.CopyN(os.Stdout, stringReader, 12)
}

func Io_Pipe() {
	fmt.Println("+++===== Io_Pipe =====+++")

	src, dst := io.Pipe()

	go func() {
		io.WriteString(dst, "Hello World.")
		io.WriteString(dst, "Amazing View From Santorini.")
		dst.Close()
	}()

	packet := make([]byte, 6)

	for {
		bytesRead, err := src.Read(packet)
		fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead, packet[:bytesRead], err)
		if err == io.EOF {
			break
		}
	}
}

/*
Buffered streams
*/
func Bytes_NewBufferString() {
	fmt.Println("+++===== Bytes_NewBufferString =====+++")

	buf := bytes.NewBufferString("[This is From bytes.NewBufferString(` `)]\n")

	fmt.Println("bytes written => ")
	strReader := strings.NewReader("[This is From strings.NewReader(` `)]\n")
	fmt.Println(buf.ReadFrom(strReader))
	fmt.Println(buf.WriteString("[This is From buf.WriteString(` `)]\n[This is From buf.WriteByte(byte(` `)) "))
	fmt.Println(buf.WriteByte(byte('@')))
	fmt.Println(buf.Write([]byte(" ]\n[This is From buf.Write([]byte(` `))]\n [This is From buf.WriteRune(rune(` `)) ")))
	fmt.Println(buf.WriteRune(rune('#')))
	fmt.Println(buf.WriteString(" ]\n"))

	fmt.Println("bytes read => ")
	packet := make([]byte, 12)
	n, err := buf.Read(packet)
	fmt.Printf("%v %v %s\n", n, err, packet)
	data, err := buf.ReadByte()
	fmt.Printf("%d %s %v\n", data, string(data), err)
	r, size, err := buf.ReadRune()
	fmt.Printf("%d %s %v %v\n", r, string(r), size, err)
	bytesData, err := buf.ReadBytes(0)
	fmt.Printf("%s %v\n", bytesData, err)
	str, err := buf.ReadString(0)
	fmt.Printf("%s %v\n", str, err)
	n1, err := buf.WriteTo(os.Stdout)
	fmt.Printf("%v %v\n", n1, err)
}

func Bufio_Writer() {
	fmt.Println("+++===== Bytes_NewBufferString =====+++")

	dst := bufio.NewWriter(os.Stdout)

	fmt.Println(dst.WriteString("Hello World"))
	fmt.Println(dst.Write([]byte(" How are you?\n")))

	fmt.Println(dst.Flush(), "FLushed!")

	fmt.Println(dst.WriteString("Good\n"))
	fmt.Println(dst.Flush(), "FLushed!")
}

func Bufio_Reader() {
	fmt.Println("+++===== Bufio_Reader =====+++")

	src := bufio.NewReader(os.Stdin)
	line, _, _ := src.ReadLine()
	fmt.Println(string(line))
}
