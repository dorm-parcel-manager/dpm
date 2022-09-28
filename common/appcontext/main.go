package appcontext

type AppContext interface {
	RequireLoggedIn() error

	RequireStudent() error
	RequireStaff() error
	RequireAdmin() error
}
