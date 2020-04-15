package ff_common_lib

type Status struct {
	State string `json:"state,omitempty"`
}

func Sum(x int, y int) int {
	return x + y
}
