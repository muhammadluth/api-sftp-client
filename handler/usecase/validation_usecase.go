package usecase

import (
	"api-sftp-client/handler"
	"fmt"
	"regexp"

	"github.com/muhammadluth/log"
)

type ValidationUsecase struct {
}

func NewValidationUsecase() handler.IValidationUsecase {
	return &ValidationUsecase{}
}

func (u *ValidationUsecase) ValidatePathDirectory(traceId, pathDirectory string) bool {
	if matched, err := regexp.MatchString(`(.*)?(?:$|(.+?)(?:(\.*$)|$))`, pathDirectory); err != nil {
		log.Error(err, traceId)
		return false
	} else if !matched {
		fmt.Println(matched)
		return false
	}
	return true
}
