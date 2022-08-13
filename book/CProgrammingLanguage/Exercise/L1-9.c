//
// Created by apple on 2021/6/14.
//
#include "stdio.h"

int main() {
    int c;
    int inspace;
    inspace = 0;
    while ((c = getchar()) != EOF) {
        // 遇到' ' 就只输出一次
        if (c == ' ') {
            if (inspace == 0) {
                // 修改标志位
                inspace = 1;
                putchar(c);
            }
        }
        // 如果是字符，就恢复
        if (c != ' ') {
            inspace = 0;
            putchar(c);
        }
    }
}