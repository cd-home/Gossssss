//
// Created by apple on 2021/6/15.
//

#include "stdio.h"

int main() {

    int c;
    int inSpace;
    inSpace = 0;
    while ((c = getchar()) != EOF) {
        if (c == ' ' || c == '\t' || c == '\n') {
            if (inSpace == 0) {
                inSpace = 1;
                putchar('\n');
            }
        } else {
            inSpace = 0;
            putchar(c);
        }
    }
    return 0;
}