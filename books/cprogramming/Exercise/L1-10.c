//
// Created by apple on 2021/6/14.
//
#include "stdio.h"

int main() {

    int c, d;

    while ((c = getchar()) != EOF) {
        d = 0;
        if (c == '\\') {
            putchar('\\');
            putchar('\\');
            d = 1;
        } else if (c == '\t') {
            putchar('\\');
            putchar('t');
            d = 1;
        } else if (c == '\b') {
            putchar('\\');
            putchar('b');
            d = 1;
        }
        if (d == 0) {
            putchar(c);
        }
    }

    return 0;
}
