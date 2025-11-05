package validations

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	"github.com/josevitorrodriguess/whisper/server/internal/models"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func UserIsValid(u models.User) error {

	if strings.TrimSpace(u.ID) == "" {
		return errors.New("invalid uid")
	}

	if strings.TrimSpace(u.Email) == "" {
		return errors.New("empty email")
	}
	if !emailRegex.MatchString(strings.TrimSpace(u.Email)) {
		return errors.New("invalid email")
	}

	if strings.TrimSpace(u.Username) == "" {
		return errors.New("empty username")
	}
	if l := len(u.Username); l < 3 || l > 30 {
		return errors.New("username must be between 3 and 30 characters")
	}

	if strings.TrimSpace(u.PhotoURL) != "" {
		parsed, err := url.ParseRequestURI(u.PhotoURL)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" {
			return errors.New("invalid photoURL")
		}
		if parsed.Scheme != "http" && parsed.Scheme != "https" {
			return errors.New("photoURL must use http or https")
		}
	}

	return nil
}
