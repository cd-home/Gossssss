//
// Created by apple on 2021/6/13.
//
#include "stdio.h"

int main() {
    int c;

    printf("%d\n", EOF);

    // EOF 结束
    printf("%d\n", getchar() != EOF);

    while ((c = getchar()) != EOF) {
        putchar(c);
    }
}