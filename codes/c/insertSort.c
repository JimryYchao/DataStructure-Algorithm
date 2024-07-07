#include "cppA.h"
#include <stdio.h>

void insertSort(int *A, size_t n)
{
    int key = 0, j = 0;
    for (size_t i = 1; i < n; i++)
    {
        key = A[i];
        // insert A[i] into the sorted subArray A[0:i-1]
        j = i - 1;
        while (j>=0 && A[j] > key)
        {
            A[j+1] = A[j];
            j--;
        }
        A[j+1] = key;
    }
}

int main() {
    int a[9] = {9, 5, 3, 4, 8, 7, 6, 1, 2};
    insertSort(a, 9);
    for (size_t i = 0; i < 9; i++)
        printf("%d,", a[i]);
}