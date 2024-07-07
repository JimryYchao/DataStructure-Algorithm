namespace csA
{
    public class MergeSort
    {
        private static void merge(int[] A, int l, int m, int r)
        {
            int[] tmp = new int[r - l + 1];
            // copy A[l:r] into tmp; and L = A[l:m], R = A[m+1:r]
            for (int i = 0; i < r - l + 1; i++)
                tmp[i] = A[l + i];
            // merge
            int iL = 0, iR = m + 1 - l, k = l;
            while (iL <= m - l && iR <= r - l)
            {
                if (tmp[iL] < tmp[iR])
                    A[k] = tmp[iL++];
                else
                    A[k] = tmp[iR++];
                k++;
            }
            // copy reminder
            while (iL <= m - l)
                A[k++] = tmp[iL++];
            while (iR <= r - l)
                A[k++] = tmp[iR++];
        }

        public static void Sort(int[] A, int p, int r)
        {
            if (p >= r)
                return;
            var m = (p + r) / 2;
            Sort(A, p, m);          // recursively sort A[p:m]
            Sort(A, m + 1, r);      // recursively sort A[m+1:r]
            merge(A, p, m, r);      // merge L and R
        }
    }
}




