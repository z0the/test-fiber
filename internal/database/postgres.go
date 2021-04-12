package database

import (
	"test-job/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresDB() *PostgresDB {
	dsn := "host=localhost user=postgres password=p@ssw0rd dbname=testdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		logrus.Fatal("Error on connection to db, err: ", err)
	}
	db.AutoMigrate(&model.Organization{})
	logrus.Info("Successful database connection!")
	return &PostgresDB{db: db}
}

func (d *PostgresDB) GetOrganization(id int) (model.Organization, error) {
	var org model.Organization
	queryRes := d.db.First(&org, id)
	return org, queryRes.Error
}
func (d *PostgresDB) GetAllOrganizations() ([]model.Organization, error) {
	var orgs []model.Organization
	queryRes := d.db.Find(&orgs)
	return orgs, queryRes.Error
}
func (d *PostgresDB) CreateOrganization(organization model.Organization) error {
	queryRes := d.db.Create(&organization)
	return queryRes.Error
}
func (d *PostgresDB) UpdateOrganization(organization model.Organization) error {
	queryRes := d.db.Model(&organization).Updates(organization)
	return queryRes.Error
}
func (d *PostgresDB) DeleteOrganization(id int) error {
	var organization model.Organization
	queryRes := d.db.Delete(&organization, id)
	return queryRes.Error
}
