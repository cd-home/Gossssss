//
// Created by apple on 2021/6/14.
//
#include "stdio.h"

int main() {

    int c, d;

    while ((c = getchar()) != EOF) {
        switch (c) {
            case '\\':
                putchar('\\');
                putchar('\\');
                break;
            case '\b':
                putchar('\\');
                putchar('b');
                break;
            case '\t':
                putchar('\\');
                putchar('t');
                break;
            default:
                putchar(c);
                break;
        }
    }

    return 0;
}
