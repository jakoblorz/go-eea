package eea

type ExtendedEuclidianRow struct {
	A int64
	B int64
	Q int64
	S int64
	T int64
}

type ExtendedEuclidianParameters struct {
	A int64 `env:"A"`
	B int64 `env:"B"`
	S int64
	T int64
}

func (parameters *ExtendedEuclidianParameters) Calculate() []*ExtendedEuclidianRow {

	// iterate using ggt-algorithm to detect total iterations required
	a := parameters.A
	b := parameters.B

	count := 0
	for b != 0 {

		t := b
		b = a % b
		a = t
		count++
	}

	// build simple point64er array for the calculated values
	values := make([]*ExtendedEuclidianRow, count+1)

	for i := range values {

		if i != 0 {

			t := parameters.B
			parameters.B = parameters.A % parameters.B
			values[i-1].Q = (parameters.A - parameters.B) / t
			parameters.A = t

			if i == count-1 {
				parameters.S = 1
			}
		}

		values[i] = &ExtendedEuclidianRow{
			A: parameters.A,
			B: parameters.B,
			Q: 0,
			S: parameters.S,
			T: parameters.T,
		}
	}

	// iterate in reverse and skip the first/last one
	// as s, t are know already there
	for i := len(values) - 2; i >= 0; i-- {
		values[i].S = values[i+1].T
		values[i].T = values[i+1].S - values[i].Q*values[i+1].T
	}

	return values

}
