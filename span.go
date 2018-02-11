package concurrency

type Span struct {
	Start int
	Count int
}

func Spans(total, spanSize int) []Span {
	spans := make([]Span, 0, (total+spanSize-1)/spanSize)
	var c int
	for i := 0; i < total; i += c {
		if i+spanSize <= total {
			c = spanSize
			spans = append(spans, Span{Start: i, Count: c})
		} else {
			c = total - i
			if c > spanSize/2 || len(spans) == 0 {
				spans = append(spans, Span{Start: i, Count: c})
			} else {
				spans[len(spans)-1].Count += c
			}
		}
	}
	return spans
}
