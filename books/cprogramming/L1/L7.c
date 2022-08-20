//
// Created by apple on 2021/6/13.
//

#include "stdio.h"

int main() {

    int c;
    long nl;
    long ns;
    long nt;

    nl = 0;
    ns = 0;
    nt = 0;
    while ((c = getchar()) != EOF) {
        // 输入的是换行才++
        if (c == '\n') {
            ++nl;
        }
        if (c == '\b') {
            ++ns;
        }
        if (c == '\t') {
            ++nt;
        }
    }
    printf("%ld\n", nl);
    printf("%ld\n", ns);
    printf("%ld\n", nt);
    return 0;
}