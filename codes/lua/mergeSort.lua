-- l = left, r = right, m = mid
local function merge(A, l, m, r)
    local nL, nR = m + 1 - l, r - m
    local tmp, i = {}, 1
    -- copy A[l:r] into tmp; and L = tmp[0:nL-1], R = tmp[nL:r]   // index from 1 to len
    for j = l, r, 1 do
        tmp[i] = A[j]
        i = i + 1
    end
    -- merge
    local iL, iR, k = 1, 1, l
    while iL <= nL and iR <= nR do
        if tmp[iL] < tmp[nL + iR] then
            A[k] = tmp[iL]
            iL = iL + 1
        else
            A[k] = tmp[nL + iR]
            iR = iR + 1
        end
        k = k + 1
    end
    -- copy reminder
    while iL <= nL do
        A[k] = tmp[iL]
        iL = iL + 1
        k = k + 1
    end
    while iR <= nR do
        A[k] = tmp[nL + iR]
        iR = iR + 1
        k = k + 1
    end
end

function MergeSort(A, p, r) 
    if p >= r then
        return
    end
    local m = math.floor((p + r) / 2)
    MergeSort(A, p, m)
    MergeSort(A, m + 1, r)
    merge(A, p, m, r)
end

a = { 9, 5, 3, 4, 8, 7, 6, 1, 2 }
MergeSort(a, 2, 9)
for _, value in ipairs(a) do
    print(value)
end
