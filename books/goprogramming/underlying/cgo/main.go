package main

//void SayHello(const char* s);
import "C"

func main() {
	C.SayHello(C.CString("Hello, World"))
}

/*
#include <stdio.h>

static void SayHello(const char* s) {
    puts(s);
}
*/
