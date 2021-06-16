package user

type User struct {
	Name
	streak Streak
}

func (u *User) CompleteDaily() {
	u.streak.Increment()
}

func (u *User) DecreaseDaily() {

}
