#include "cppA.h"
#include <iostream>
using namespace std;

namespace cppA
{
    void InsertSort(int A[], size_t n)
    {
        int key = 0, j = 0;
        for (size_t i = 1; i < n; i++)
        {
            key = A[i];
            // insert A[i] into the sorted subArray A[0:i-1]
            j = i - 1;
            while (j >= 0 && A[j] > key)
            {
                A[j + 1] = A[j];
                j--;
            }
            A[j + 1] = key;
        }
    }
}


int main()
{
    using namespace std;
    int a[] = {9, 5, 3, 4, 8, 7, 6, 1, 2};
    cppA::InsertSort(a, 9);
    for (auto &e : a)
        cout << e << ",";
}