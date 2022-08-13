// Create by GodYao 2021/6/15 22:10
#include "stdio.h"


int main() {
    int c, i, nwhite, nother;
    int ndigit[10];

    i = 0;
    while (i < 10)
        ndigit[i++] = 0;


    while ((c = getchar()) != EOF) {
        if (c >= '0' && c <= '9')
            ++ndigit[c - '0'];
        else if (c == ' ' || c == '\n' || c == '\t')
            ++nwhite;
        else
            ++nother;
    }

    int j;
    j = 0;
    while (j < 10)
        printf("%d ", ndigit[j++]);

    printf("nwhite = %d, nother = %d\n", nwhite, nother);

    return 0;
}