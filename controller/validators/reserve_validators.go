package validators

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
)

func ValidateReservationParams(*entities.ReservationRequest) error {

	return nil
}

func ValidateUserId(userId uuid.UUID) error {
	return nil
}
