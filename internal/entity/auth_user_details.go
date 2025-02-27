package entity

import (
	"fmt"
	"time"

	"github.com/photoprism/photoprism/pkg/rnd"
)

const (
	GenderMale    = "male"
	GenderFemale  = "female"
	GenderOther   = "other"
	GenderUnknown = ""
)

// UserDetails represents user profile information.
type UserDetails struct {
	UserUID      string    `gorm:"type:VARBINARY(42);primary_key;auto_increment:false;" json:"-" yaml:"-"`
	IdURL        string    `gorm:"type:VARBINARY(512);column:id_url;" json:"IdURL,omitempty" yaml:"IdURL,omitempty"`
	SubjUID      string    `gorm:"type:VARBINARY(42);index;" json:"SubjUID,omitempty" yaml:"SubjUID,omitempty"`
	SubjSrc      string    `gorm:"type:VARBINARY(8);default:'';" json:"-" yaml:"SubjSrc,omitempty"`
	PlaceID      string    `gorm:"type:VARBINARY(42);index;default:'zz'" json:"-" yaml:"-"`
	PlaceSrc     string    `gorm:"type:VARBINARY(8);" json:"-" yaml:"PlaceSrc,omitempty"`
	CellID       string    `gorm:"type:VARBINARY(42);index;default:'zz'" json:"-" yaml:"CellID,omitempty"`
	BirthYear    int       `gorm:"default:-1;" json:"BirthYear" yaml:"BirthYear,omitempty"`
	BirthMonth   int       `gorm:"default:-1;" json:"BirthMonth" yaml:"BirthMonth,omitempty"`
	BirthDay     int       `gorm:"default:-1;" json:"BirthDay" yaml:"BirthDay,omitempty"`
	NamePrefix   string    `gorm:"size:32;" json:"NamePrefix" yaml:"NamePrefix,omitempty"`
	GivenName    string    `gorm:"size:64;" json:"GivenName" yaml:"GivenName,omitempty"`
	MiddleName   string    `gorm:"size:64;" json:"MiddleName" yaml:"MiddleName,omitempty"`
	FamilyName   string    `gorm:"size:64;" json:"FamilyName" yaml:"FamilyName,omitempty"`
	NameSuffix   string    `gorm:"size:32;" json:"NameSuffix" yaml:"NameSuffix,omitempty"`
	NickName     string    `gorm:"size:64;" json:"NickName" yaml:"NickName,omitempty"`
	UserGender   string    `gorm:"size:16;" json:"Gender" yaml:"Gender,omitempty"`
	UserBio      string    `gorm:"size:512;" json:"Bio" yaml:"Bio,omitempty"`
	UserLocation string    `gorm:"size:512;" json:"Location" yaml:"Location,omitempty"`
	UserCountry  string    `gorm:"type:VARBINARY(2);" json:"Country" yaml:"Country,omitempty"`
	UserPhone    string    `gorm:"size:32;" json:"Phone" yaml:"Phone,omitempty"`
	SiteURL      string    `gorm:"type:VARBINARY(512);column:site_url" json:"SiteURL" yaml:"SiteURL,omitempty"`
	ProfileURL   string    `gorm:"type:VARBINARY(512);column:profile_url" json:"ProfileURL" yaml:"ProfileURL,omitempty"`
	FeedURL      string    `gorm:"type:VARBINARY(512);column:feed_url" json:"FeedURL,omitempty" yaml:"FeedURL,omitempty"`
	AvatarURL    string    `gorm:"type:VARBINARY(512);column:avatar_url" json:"AvatarURL,omitempty" yaml:"AvatarURL,omitempty"`
	OrgName      string    `gorm:"size:128;" json:"OrgName" yaml:"OrgName,omitempty"`
	OrgTitle     string    `gorm:"size:64;" json:"OrgTitle" yaml:"OrgTitle,omitempty"`
	OrgEmail     string    `gorm:"size:255;index;" json:"OrgEmail" yaml:"OrgEmail,omitempty"`
	OrgPhone     string    `gorm:"size:32;" json:"OrgPhone" yaml:"OrgPhone,omitempty"`
	OrgURL       string    `gorm:"type:VARBINARY(512);column:org_url" json:"OrgURL" yaml:"OrgURL,omitempty"`
	CreatedAt    time.Time `json:"CreatedAt" yaml:"-"`
	UpdatedAt    time.Time `json:"UpdatedAt" yaml:"-"`
}

// TableName returns the entity table name.
func (UserDetails) TableName() string {
	return "auth_users_details"
}

// NewUserDetails creates new user details.
func NewUserDetails(uid string) *UserDetails {
	return &UserDetails{UserUID: uid}
}

// CreateUserDetails creates new user details or returns nil on error.
func CreateUserDetails(user *User) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}

	if user.UID() == "" {
		return fmt.Errorf("empty user uid")
	}

	user.UserDetails = NewUserDetails(user.UID())

	if err := Db().Where("user_uid = ?", user.UID()).First(user.UserDetails).Error; err == nil {
		return nil
	}

	return user.UserDetails.Create()
}

// HasID tests if the entity has a valid uid.
func (m *UserDetails) HasID() bool {
	return rnd.IsUID(m.UserUID, UserUID)
}

// Create new entity in the database.
func (m *UserDetails) Create() error {
	return Db().Create(m).Error
}

// Save updates the record in the database or inserts a new record if it does not already exist.
func (m *UserDetails) Save() error {
	return Db().Save(m).Error
}

// Updates multiple properties in the database.
func (m *UserDetails) Updates(values interface{}) error {
	return UnscopedDb().Model(m).Updates(values).Error
}
