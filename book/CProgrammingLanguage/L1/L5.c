//
// Created by apple on 2021/6/13.
//
#include "stdio.h"

int main() {
    // 声明
    int nc;
    // 初始化
    nc = 0;
    while (getchar() != EOF) {
        ++nc;
    }

    printf("%d\n", nc);

    return 0;
}
