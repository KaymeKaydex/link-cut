package model

type URLContainer struct {
	ShortUrl        string    `gorm:"primaryKey;column:short_url"`
	LongUrl    		string    `gorm:"column:long_url;not null;index"`
}
