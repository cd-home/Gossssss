//
// Created by apple on 2021/6/13.
//
#include "stdio.h"

// 符号常量， 不需要分号
#define LOWER 0
#define UPPER 300
#define STEP 20

int main() {

    int fahr;

    for (fahr = LOWER; fahr <= UPPER; fahr += STEP) {
        printf("%3d %6.1f\n", fahr, (5.0 / 9.0) * (fahr - 32));
    }

    return 0;
}
