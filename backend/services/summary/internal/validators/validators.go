package validators

import (
	"errors"
	"log"

	pb "github.com/Azat201003/eduflow/backend/libs/gen/go/summary"
)

type Validator struct{}

func (v *Validator) Validate(obj interface{}) error {
	switch i := obj.(type) {
	case *pb.FilterRequest:
		return v.ValidatePageSize(i.Page.Size)
	default:
		return nil
	}
}

func (v *Validator) ValidatePageSize(obj uint32) error {
	log.Println(obj)
	if obj == 0 {
		return errors.New("Page size is 0")
	}
	return nil
}
