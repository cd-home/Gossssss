//
// Created by apple on 2021/6/13.
//

#include "stdio.h"

int main() {
    int blanks, tabs, newlines;
    int c;
    int done = 0;
    int last_char = 0;

    blanks = 0;
    tabs = 0;
    newlines = 0;

    while (done == 0) {
        c = getchar();
        if (c == ' ')
            ++blanks;
        if (c == '\t')
            ++tabs;
        if (c == '\n')
            ++newlines;
        if (c == EOF) {
            if (last_char != '\n') {
                ++newlines;
            }
            done = 1;
        }
        last_char = c;
    }
    printf("Blanks: %d\nTabs: %d \nNewlines: %d\n", blanks, tabs, newlines);
    return 0;
}