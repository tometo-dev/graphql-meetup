package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tsuki42/graphql-meetup/models"
)

type MeetupRepo struct {
	DB *gorm.DB
}

func (m *MeetupRepo) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Table("MEETUP").Order("name").Find(&meetups).Error
	if err != nil {
		return nil, err
	} else {
		return meetups, nil
	}
}

func (m *MeetupRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	meetup.ID = uuid.New().String()
	if err := m.DB.Table("MEETUP").Create(meetup).Error; err != nil {
		return nil, err
	} else {
		return meetup, nil
	}
}

func (m *MeetupRepo) GetMeetupsByUser(userID string) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Table("MEETUP").Where("user_id = ?", userID).Find(&meetups).Error
	if err != nil {
		return nil, err
	}
	return meetups, nil
}

func (m *MeetupRepo) GetMeetupByID(id string) (*models.Meetup, error) {
	var meetup models.Meetup
	err := m.DB.Table("MEETUP").Where("id = ?", id).Find(&meetup).Error
	return &meetup, err
}

func (m *MeetupRepo) UpdateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	err := m.DB.Table("MEETUP").Where("id = ?", meetup.ID).Update(meetup).Error
	return meetup, err
}

func (m *MeetupRepo) DeleteMeetup(id string) error {
	err := m.DB.Table("MEETUP").Where("id = ?", id).Delete(&models.Meetup{}).Error
	if err != nil {
		return fmt.Errorf("error occured while deleting: %v", err)
	}
	return nil
}
