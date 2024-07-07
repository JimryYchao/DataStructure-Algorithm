#pragma once
#include <vector>

namespace cppA
{
    // MergeSort
    static void merge(std::vector<int> &A, size_t l, size_t m, size_t r);
    void MergeSort(std::vector<int> &A, size_t begin, size_t end);

    // InsertSort
    void insertSort(int A[], size_t n);    
}
