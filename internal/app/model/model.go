package model

// Модель хранилища URL для базы данных

type URLContainer struct {
	ShortUrl string `gorm:"primaryKey;column:short_url"`    // Хранит символы после доменного имени
	LongUrl  string `gorm:"column:long_url;not null;index"` // Хранит полный URL который необходимо скоратить
}

// Функция которую требует gorm для того чтобы обращаться по имени таблицы

func (URLContainer) TableName() string {
	return "url_container"
}
