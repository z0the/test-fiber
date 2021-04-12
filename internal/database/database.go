package database

import (
	"test-job/internal/model"
)

type Database interface {
	GetOrganization(id int) (model.Organization, error)
	GetAllOrganizations() ([]model.Organization, error)
	CreateOrganization(organization model.Organization) error
	UpdateOrganization(organization model.Organization) error
	DeleteOrganization(id int) error
}
