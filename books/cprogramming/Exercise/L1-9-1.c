//
// Created by apple on 2021/6/14.
//

#include "stdio.h"

int main() {

    int c, pc; // previous character

    pc = EOF;

    while ((c = getchar()) != EOF) {
        if (c == ' ') {
            if (pc != ' ') {
                putchar(c);
            }
        }
        if (c != ' ') {
            putchar(c);
        }
        pc = c;
    }

    return 0;
}
