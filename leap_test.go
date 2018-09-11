package leap

import "testing"

func Benchmark400(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		IsLeapYear(4)
	}
}
