#include "cppA.h"
#include <iostream>
using namespace std;

namespace cppA
{

    static void merge(vector<int> &A, size_t l, size_t m, size_t r)
    {
        size_t nL = m + 1 - l;
        size_t nR = r - m;
        int *tmp = new int[r - l + 1];
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

        delete[] tmp;
    }

    void MergeSort(vector<int> &A, size_t begin, size_t end)
    {
        if (begin >= end)
            return;
        auto m = (begin + end) / 2;
        MergeSort(A, begin, m);   // recursively sort A[p:m]
        MergeSort(A, m + 1, end); // recursively sort A[m+1:r]
        merge(A, begin, m, end);  // merge L and R
    }
}

int main()
{
    using namespace std;
    vector<int> a = {9, 5, 3, 4, 8, 7, 6, 1, 2};
    cppA::MergeSort(a, 2, 8);
    for (auto &e : a)
        cout << e << ",";
}