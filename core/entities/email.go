package entities

import (
	"time"

	"gitlab.com/meta-node/mail/core/pkgs/datatypes"
	"gorm.io/gorm"
)

type Email struct {
	ID       uint                  `gorm:"type:int(11);autoIncrement;primaryKey" json:"id"`
	Title    string                `gorm:"type:text;not null" json:"title"`
	Subject  string                `gorm:"type:text;not null" json:"subject"`
	WriterID uint                  `gorm:"type:int(11)" json:"writer_id"`
	To       datatypes.StringArray `gorm:"type:text;not null" json:"to"`
	From     string                `gorm:"type:varchar(191);not null" json:"from"`
	Content  string                `gorm:"type:text" json:"content"`
	Cc       datatypes.StringArray `gorm:"type:text" json:"cc"`
	Bcc      datatypes.StringArray `gorm:"type:text" json:"bcc"`
	Status   string                `gorm:"type:enum('Draft', 'Sending', 'Receiving', 'Approved', 'Sent', 'Seen', 'Declined', 'Dropped');default:Null" json:"status"`
	Files    []File                `gorm:"foreignkey:EmailID" json:"files"`

	CreatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	UpdatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"datetime(3)" json:"-"`
}

type File struct {
	ID      uint   `gorm:"type:int(11);primaryKey;autoIncrement" json:"id"`
	EmailID uint   `gorm:"type:int(11)" json:"email"`
	Path    string `gorm:"type:text" json:"path"`
}

type EmailInformation struct {
	HandlerID uint           `gorm:"int(11);primaryKey" json:"hanlder_id"`
	EmailID   uint           `gorm:"int(11);primaryKey" json:"email_id"`
	Status    string         `gorm:"type:enum('Draft', 'Sending', 'Receiving', 'Approved', 'Sent', 'Seen', 'Declined', 'Dropped');default:Null" json:"status"`
	CreatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	UpdatedAt time.Time      `gorm:"datetime(3)" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"datetime(3)" json:"-"`
}
