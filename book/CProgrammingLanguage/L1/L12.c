// Create by GodYao 2021/6/28 22:20
#include "stdio.h"


int main() {
//    printf("%d", 15 >> 1);
    printf("%d", (15 >> 2) & ~(~0 << 3));
    return 0;
}