package auth

type Role string
type Version int
type TimeInt int

const (
	Admin Role = "ADMIN"
	User  Role = "USER"
)

const (
	AccessTokenExpired  TimeInt = 15           // 15 mins
	RefreshTokenExpired TimeInt = 60 * 24 * 30 // 30 days
)
const (
	InitialSessionVersion Version = 1
)

func ParseRole(role string) (Role, error) {
	switch Role(role) {
	case Admin, User:
		return Role(role), nil
	default:
		return "", ErrRoleInvalid
	}
}

func (r Role) String() string {
	return string(r)
}

func (r Role) Compare(diff Role) bool {
	return r == diff
}
