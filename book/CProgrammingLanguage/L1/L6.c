//
// Created by apple on 2021/6/13.
//
#include "stdio.h"

int main() {

    double nc;

    for (nc = 0; getchar() != EOF; ++nc);

    printf("%.0f\n", nc);

    return 0;
}