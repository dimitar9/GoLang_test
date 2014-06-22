package try_testing

import "testing"
import "fmt"

func Test1(t *testing.T) {
    if dimi_minus(3,2) != 1 {
        t.Error("dimi minus is wrong!")
    }
}


func BenchmarkHello(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}

