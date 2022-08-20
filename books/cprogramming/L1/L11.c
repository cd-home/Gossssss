// Create by GodYao 2021/6/15 23:10
#include "stdio.h"

#define MAX_CHARS 100

void copy(char from[], char to[]);

int get_line(char s[], int limit);

int main() {
    int len;
    int max;
    char line[MAX_CHARS];
    char longest[MAX_CHARS];

    max = 0;
    while ((len = get_line(line, MAX_CHARS)) > 0) {
        if (len > max) {
            max = len;
            copy(line, longest);
        }
    }
    if (max > 0) {
        printf("%s", longest);
    }
    return 0;
}


int get_line(char s[], int limit){
    int c, i;
    for (i = 0; i < limit - 1 && (c = getchar()) != EOF && c != '\n'; ++i)
        s[i] = c;
    if (c == '\n') {
        s[i] = c;
        ++i;
    }
    // '\n' 最后一个
    s[i] = '\0';
    return i;
}


void copy(char from[], char to[]) {
    int i;
    i = 0;
    // '\0' 代表末尾
    while ((to[i] = from[i]) != '\0')
        ++i;
}