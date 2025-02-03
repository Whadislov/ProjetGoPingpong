package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key to sign the token
var jwtSecret = []byte("my_secret_jwt")

// generateJWT generates a JWT token
func generateJWT(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Handler de connexion : génère un token et le retourne en JSON
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// 🚨 Simule une connexion réussie avec un ID utilisateur "12345"
	userID := 12345

	// Générer le token JWT
	token, err := generateJWT(userID)
	if err != nil {
		http.Error(w, "Erreur lors de la génération du token", http.StatusInternalServerError)
		return
	}

	// Retourner le token en JSON
	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Middleware to verify JSON Web Token (JWT)
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Extract the token after "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract User ID from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Impossible to read the token", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(string) // User ID needs to be a string
		if !ok {
			http.Error(w, "Invalid user ID ", http.StatusUnauthorized)
			return
		}

		// Add user ID in the header
		r.Header.Set("X-User-ID", userID)

		// Call next header
		next(w, r)
	}
}
