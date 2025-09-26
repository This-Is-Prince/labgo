package main

/*
#include <stdio.h>

int sayHello(void) {
	printf("Helloooooo");
	return 0;
}
*/
import "C"

func main() {
	C.sayHello()
}
