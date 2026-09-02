package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/awsbuilderslpu/Community-Day---Go-Backend/internal/module/auth"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserContextKey contextKey = "user"

type Claims struct {
	UserID uint      `json:"user_id"`
	Role   auth.Role `json:"role"`
	jwt.RegisteredClaims
}

func AuthMiddleware(jwtSecret string, requiredRoles ...auth.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Authorization header required", http.StatusUnauthorized)
				return
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
				http.Error(w, "Invalid token format", http.StatusUnauthorized)
				return
			}

			tokenStr := bearerToken[1]
			claims := &Claims{}

			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(jwtSecret), nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			// RBAC Check
			if len(requiredRoles) > 0 {
				hasRole := false
				for _, role := range requiredRoles {
					if claims.Role == role {
						hasRole = true
						break
					}
				}
				if !hasRole {
					http.Error(w, "Forbidden: insufficient permissions", http.StatusForbidden)
					return
				}
			}

			ctx := context.WithValue(r.Context(), UserContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
