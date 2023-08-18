package middleware

import (
	"context"
	"net/http"

	pb "github.com/michael-m-truong/fitness-tracker/pb" // Import the generated gRPC code
	"github.com/michael-m-truong/fitness-tracker/services"
)

// maybe add some encapsulation soon
var UserCtxKey = &contextKey{"user"}

// also works as type contextKey string
type contextKey struct {
	name string
}

type IAuthMiddleware interface {
	Auth() func(http.Handler) http.Handler
}

type AuthMiddleware struct {
	AuthService services.IAuthService
}

func (mw AuthMiddleware) Auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Validate JWT token using the auth service client
			tokenStr := header
			user, err := mw.AuthService.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// Convert the gRPC user response to the appropriate struct
			authUser := &pb.User{
				Username: user.Username,
			}

			// Put user information in the context
			ctx := context.WithValue(r.Context(), UserCtxKey, authUser)

			// Call the next handler with the new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *pb.User {
	raw, _ := ctx.Value(UserCtxKey).(*pb.User)
	return raw
}
