namespace csA;

public class InsertSort
{
    public static void Sort(int[] A, int n)
    {
        int key = 0, j = 0;
        for (int i = 1; i < n; i++)
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
