function InsertSort(A, n)
    local j, key = 0, 0
    for i = 2, n, 1 do
        key = A[i]
        j = i - 1
        while j > 0 and A[j] > key do
            A[j + 1] = A[j]
            j = j - 1
        end
        A[j + 1] = key
    end
end

a = { 9, 5, 3, 4, 8, 7, 6, 1, 2 }
InsertSort(a, 9)
for _, value in ipairs(a) do
    print(value)
end
