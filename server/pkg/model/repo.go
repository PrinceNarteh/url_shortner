package model

func GetAllUrls() ([]URL, error) {
	urls := make([]URL, 0)

	tx := db.Find(urls)
	if tx.Error != nil {
		return []URL{}, tx.Error
	}

	return urls, nil
}

func FindUrl(id int64) (URL, error) {
	var url URL

	tx := db.Where("id = ?", id).First(&url)

	if tx.Error != nil {
		return URL{}, nil
	}

	return url, nil
}

func CreateUrl(url URL) error {
	tx := db.Create(&url)
	return tx.Error
}

func UpdateUrl(url URL) error {
	tx := db.Save(url)
	return tx.Error
}

func DeleteUrl(id uint64) error {
	tx := db.Unscoped().Delete(&URL{}, id)
	return tx.Error
}

func FindByUrl(url URL) (URL, error) {
	var link URL

	tx := db.Where("link = ?", link).First(&link)
	if tx.Error != nil {
		return URL{}, tx.Error 
	}

	return link, nil
}
