package ports

import (
	"github.com/cyneptic/letsgo/internal/core/entities"
	"github.com/google/uuid"
)

type ReserveRepositoryContract interface {
	FilterFlightList(userId uuid.UUID) error
	SortFlightList(reservation entities.Reservation) error
}
