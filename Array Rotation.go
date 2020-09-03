package main 
import (
    f "fmt"
    "bufio"
    "os"
)

func right_rotate(a []int, N, K int) {
    r := append(a[(N-K):N], a[:(N-K)]...)
    for i:=0; i<len(r); i++ {
        f.Printf("%d ", r[i])
    }
    f.Println("\n")
}

func left_rotate(a []int, N, K int) {
    r:=append(a[-K:N], a[:-K]...)
    for i:=0; i<len(r); i++ {
        f.Printf("%d ", r[i])
    }
    f.Println("\n")
}

func main() {
    in := bufio.NewReader(os.Stdin)
    mat := make([]int,100000)
    var T, N, K int
    f.Fscan(in, &T)
    for ; T>0; T-- {
        f.Fscan(in, &N, &K)
        arr := mat[:N]
        for i:=0; i<N; i++ {
            f.Fscan(in, &arr[i])
        }
        K%=N
        if K>=0{
            right_rotate(arr, N, K)
        } else {
            left_rotate(arr, N, K)    
        }
    }
}
