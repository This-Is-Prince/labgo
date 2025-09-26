package main

/*
#include <stdio.h>
#include "libhey.h"

int sayHello(void) {
	printf("Helloooooo");
	return 0;
}
*/
import "C"

func main() {
	C.sayHello()
	C.Hey()
}
