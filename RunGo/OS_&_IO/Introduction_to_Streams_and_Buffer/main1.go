package main

/*
 ================
 Reading From a Data Source
 ================
*/
/* ============ 1) io.Reader ============ */
/*
// MyStringData - simple struct to hold string data
type MyStringData struct {
	str       string
	readIndex int // default: 0
}

// add `Read` method (to implement io.Reader)
func (myStringData *MyStringData) Read(p []byte) (n int, err error) {
	// convert `str` string to a slice of bytes
	strBytes := []byte(myStringData.str)

	// if `readIndex` is GTE source length, return `EOF` error
	if myStringData.readIndex >= len(strBytes) {
		return 0, io.EOF // `0` bytes read
	}

	// get next readable limit (exclusive)
	nextReadLimit := myStringData.readIndex + len(p)

	// if `nextReadLimit` is GTE source length
	// set `nextReadLimit` to source length and `err` to `EOF`
	if nextReadLimit >= len(strBytes) {
		nextReadLimit = len(strBytes)
		err = io.EOF
	}

	// get next bytes to copy and set `n` to its length
	nextBytes := strBytes[myStringData.readIndex:nextReadLimit]
	n = len(nextBytes)

	// copy all bytes of `nextBytes` into `p` slice
	copy(p, nextBytes)
	// increment `readIndex` to `nextReadLimit`
	myStringData.readIndex = nextReadLimit

	// return values
	return
}

func main() {
	// create data source
	src := MyStringData{
		str: "Hello Amazing World!",
	}

	// create a packet
	p := make([]byte, 3) // slice of length `3`

	// read `src` util an error is returned
	for {
		// read `p` bytes from `src`
		n, err := src.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

		// handle error
		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occurred!", err)
			break
		}
	}
}
*/

/* ============ 2) strings.NewReader ============ */
/*
func main() {
	// create data source
	src := strings.NewReader("Hello Amazing World!")

	// create a packet
	p := make([]byte, 3) //slice of length `3`

	// read `src` until an error is returned
	for {
		// read `p` bytes from `src`
		n, err := src.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

		// handle error
		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occurred!", err)
			break
		}
	}
}
*/

/* ============ 3) ioutil.ReadAll ============ */
/*
func main() {
	// create data source
	src := strings.NewReader("Hello Amazing World!")

	// read all data from `src`
	data, err := ioutil.ReadAll(src)
	if err != nil {
		log.Fatal(err)
		return
	}

	// print `data`
	fmt.Printf("Read data of length %d : %s\n", len(data), data)
}
*/

/* ============ 4) ioutil.ReadFull ============ */
/*
func main() {
	// create data source
	src := strings.NewReader("Hello Amazing World!") // 20 characters

	// create buffer of length 14
	buf := make([]byte, 14)

	// call 1: read from `src`
	bytesRead1, err1 := io.ReadFull(src, buf)
	fmt.Printf("Bytes read: %d, value: `%s`, err: %v\n", bytesRead1, buf[:bytesRead1], err1)

	// call 2: read from `src`
	bytesRead2, err2 := io.ReadFull(src, buf)
	fmt.Printf("Bytes read: %d, value: `%s`, err: %v\n", bytesRead2, buf[:bytesRead2], err2)

	// call 3: read from `src`
	bytesRead3, err3 := io.ReadFull(src, buf)
	fmt.Printf("Bytes read: %d, value: `%s`, err: %v\n", bytesRead3, buf[:bytesRead3], err3)
}
*/

/* ============ 4.1) io.ReadAtLeast ============ */
/*
func main() {
	src := strings.NewReader("Hello Amazing World!")
	buf := make([]byte, 14)

	// call 1: read from `src`
	bytesRead1, err1 := io.ReadAtLeast(src, buf, 1)
	fmt.Printf("Bytes read: %d, value: `%s`, err: %v\n", bytesRead1, buf[:bytesRead1], err1)

	// call 2: read from `src`
	bytesRead2, err2 := io.ReadAtLeast(src, buf, 1)
	fmt.Printf("Bytes read: %d, value: `%s`, err: %v\n", bytesRead2, buf[:bytesRead2], err2)

	// call 3: read from `src`
	bytesRead3, err3 := io.ReadAtLeast(src, buf, 1)
	fmt.Printf("Bytes read: %d, value: `%s`, err: %v\n", bytesRead3, buf[:bytesRead3], err3)
} */

/* ============ 5) io.LimitReader ============ */
/*
func main() {
	// create a main data source
	mainSrc := strings.NewReader("Hello Amazing World!") // 20 characters

	// create data source from `mainSrc` with cap of `10` bytes
	src := io.LimitReader(mainSrc, 10)

	// create a packet
	p := make([]byte, 3) // slice of length `3`

	// read `src` until an error is returned
	for {
		// read `p` bytes from `src`
		n, err := src.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

		// handle error
		if err == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occurred!", err)
			break
		}
	}
}
*/

/*
 ================
 Writing to a Data Store
 ================
*/
/* ============ 1) io.Writer ============ */
/*
// SampleStore - sample store type
type SampleStore struct {
	data []byte
}

// implement `io.Writer` interface
func (ss *SampleStore) Write(p []byte) (n int, err error) {
	// check if `10` bytes has been written
	if len(ss.data) == 10 {
		return 0, io.EOF // end of limit error
	}

	// get remaining capacity of the `ss.data`
	remainingCap := 10 - len(ss.data)

	// get length of data to write
	writeLength := len(p)
	if remainingCap <= writeLength {
		writeLength = remainingCap
		err = io.EOF
	}

	// append `writeLength` of data from `p` to `ss.data`
	ss.data = append(ss.data, p[:writeLength]...)

	// set number of bytes written and return
	n = writeLength
	return
}

func main() {
	ss := SampleStore{}

	// write 1: "Hello!"
	bytesWritten1, err1 := ss.Write([]byte("Hello!"))
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten1, err1)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

	// write 2: " Amazing"
	bytesWritten2, err2 := ss.Write([]byte(" Amazing"))
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten2, err2)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

	// write 3: " World!"
	bytesWritten3, err3 := ss.Write([]byte(" World!"))
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten3, err3)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)
}
*/

/* ============ 2) io.WriteString ============ */
/*
// SampleStore - sample store type
type SampleStore struct {
	data []byte
}

// implement `io.Writer` interface
func (ss *SampleStore) Write(p []byte) (n int, err error) {
	// check if `10` bytes has been written
	if len(ss.data) == 10 {
		return 0, io.EOF // end of limit error
	}

	// get remaining capacity of the `ss.data`
	remainingCap := 10 - len(ss.data)

	// get length of data to write
	writeLength := len(p)
	if remainingCap <= writeLength {
		writeLength = remainingCap
		err = io.EOF
	}

	// append `writeLength` of data from `p` to `ss.data`
	ss.data = append(ss.data, p[:writeLength]...)

	// set number of bytes written and return
	n = writeLength
	return
}

func main() {
	ss := &SampleStore{}

	// write 1: "Hello!"
	bytesWritten1, err1 := io.WriteString(ss, "Hello!")
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten1, err1)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

	// write 2: " Amazing"
	bytesWritten2, err2 := io.WriteString(ss, " Amazing")
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten2, err2)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)

	// write 3: " World!"
	bytesWritten3, err3 := io.WriteString(ss, " World!")
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten3, err3)
	fmt.Printf("Value of ss.data: %s\n\n", ss.data)
}
*/

/*
 ================
 Standard I/O Streams
 ================
*/
/*
func main() {
	// use `io.WriteString` to write to a `io.Writer`
	io.WriteString(os.Stdout, "Hello World!\n")

	// call `Write` method of a `io.Writer`
	os.Stdout.Write([]byte("Hello World!\n"))

	// use `fmt` package function to write to a `io.Writer`
	fmt.Fprint(os.Stdout, "Hello World!\n")
	fmt.Fprintln(os.Stdout, "Hello World!") // adds new line
	fmt.Fprintf(os.Stdout, "%s, World!\n", "Hello")
}
*/

/*
 ================
 Transferring Data between streams
 ================
*/

/* ============ io.Copy ============ */
/*
func main() {
	// create a string `io.Reader` object
	stringReader := strings.NewReader("Hello World! How are you?\n")

	// copy data from `stringsReader` to `os.Stdout` (`io.Writer`)
	io.Copy(os.Stdout, stringReader)
}
*/

/* ============ io.CopyN ============ */
/*
func main() {
	// create a string `io.Reader` object
	stringReader := strings.NewReader("Hello World! How are you?\n")

	// copy `12 bytes` data from `stringsReader`
	// to `os.Stdout` (`io.Writer`)
	io.CopyN(os.Stdout, stringReader, 12)
}
*/

/* ============ io.Pipe ============ */
/*
func main() {
	// create a pipe
	src, dst := io.Pipe()

	// start goroutine that writes data to `dst`
	go func() {
		dst.Write([]byte("DATA_1")) // write and block
		dst.Write([]byte("DATA_2")) // write and block
		dst.Close()                 // indicate EOF
	}()

	// data transfer packet
	packet := make([]byte, 6)

	// read from `src`
	bytesRead1, err1 := src.Read(packet)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead1, packet, err1)

	// read from `src`
	bytesRead2, err2 := src.Read(packet)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead2, packet, err2)

	// read from `src`
	bytesRead3, err3 := src.Read(packet)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead3, packet, err3)
}
*/
