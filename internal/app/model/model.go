package model

// URLContainer - модель хранилища URL для базы данных
type URLContainer struct {
	ShortUrl string `gorm:"primaryKey;column:short_url;type:varchar(10);not null"` // Хранит символы после доменного имени
	LongUrl  string `gorm:"column:long_url;not null;unique"`                       // Хранит полный URL который необходимо скоратить
}

// TableName функция, которую требует gorm, для того чтобы обращаться по имени таблицы
func (URLContainer) TableName() string {
	return "url_container"
}
