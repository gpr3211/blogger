package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gpr3211/blogger/internal/clog"
)

func GetApiHead(r http.Header) (string, error) {
	head := r.Get("Authorization")
	if len(head) == 0 {
		clog.Printf("Failed to parse header %s", head)
		return "", errors.New("could not fetch apikey")
	}
	stripper := strings.Fields(head)
	if stripper[0] != "ApiKey" || len(stripper) != 2 {
		clog.Println("Failed parsing key")
		return "", errors.ErrUnsupported
	}
	return stripper[1], nil

}
