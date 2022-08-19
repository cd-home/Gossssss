//
// Created by apple on 2021/6/11.
//
#include "stdio.h"

int main() {
    // 先声明
    int lower, upper, step;
    float fahr, celsius;

    // 赋值
    lower = 0;
    upper = 300;
    step = 20;
    fahr = lower;

    // 循环
    while (fahr <= upper) {
        celsius = (5.0 / 9.0) * (fahr - 32.0);
        printf("%3.0f %6.1f\n", fahr, celsius);
        fahr += step;
    }

    return 0;
}
