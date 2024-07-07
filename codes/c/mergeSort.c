#include "cppA.h"
#include <stdio.h>

static void merge(int *A, size_t l, size_t m, size_t r)
{
    size_t nL = m + 1 - l;
    size_t nR = r - m;
    int tmp[r - l + 1];
    // copy A[l:r] into tmp; and L = tmp[0:nL-1], R = tmp[nL:r]
    for (size_t i = 0; i < r - l + 1; i++)
        tmp[i] = A[l + i];
    // merge
    int iL = 0, iR = 0, k = l;
    while (iL < nL && iR < nR)
    {
        if (tmp[iL] < tmp[nL + iR])
            A[k] = tmp[iL++];
        else
            A[k] = tmp[nL + iR++];
        k++;
    }
    // copy reminder
    while (iL < nL)
        A[k++] = tmp[iL++];
    while (iR < nR)
        A[k++] = tmp[nL + iR++];
}
void mergeSort(int *A, size_t l, size_t r)
{
    if (l >= r)
        return;
    size_t m = (l + r) / 2;
    mergeSort(A, l, m);     // recursively sort A[p:m]
    mergeSort(A, m + 1, r); // recursively sort A[m+1:r]
    merge(A, l, m, r);      // merge L and R
}

int main()
{
    printf("Hello\n");
    int a[9] = {9, 5, 3, 4, 8, 7, 6, 1, 2};
    mergeSort(a, 0, 8);
    for (size_t i = 0; i < 9; i++)
        printf("%d,", a[i]);
}