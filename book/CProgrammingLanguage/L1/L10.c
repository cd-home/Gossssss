// Create by GodYao 2021/6/15 22:52
#include "stdio.h"

int power(int, int);

int power(int base, int n) {
    int p;
    p = 1;
    while (n-- > 0)
        p *= base;
    return p;
}

int main() {
    printf("%d\n", power(2, 3));
    printf("%d\n", power(2, 4));
    printf("%d\n", power(4, 3));
    return 0;
}