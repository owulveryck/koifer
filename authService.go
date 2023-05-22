package koifer

import (
	"encoding/json"
	"net/http"
)

// AuthService is a structure that allows authentication of a user through a UserRepository
type AuthService struct {
	userRepository UserRepository
}

// authService implements http.Handler
func (authservice *AuthService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Password == "" {
		http.Error(w, "Name and password are required.", http.StatusUnauthorized)
		return
	}

	authenticate, err := authservice.authenticate(user.Name, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !authenticate {
		http.Error(w, "Unauthorized access.", http.StatusUnauthorized)
		return
	}

	token := generateRandomToken()
	response := map[string]string{"token": token}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func NewAuthService(userRepository UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

// Authenticate the user
func (s *AuthService) authenticate(name, password string) (bool, error) {
	user, err := s.userRepository.GetUserByName(name)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil
	}
	return user.Password == password, nil
}
