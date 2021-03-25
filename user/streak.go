package user

type Streak interface {
	Increment()
	String() string
}
