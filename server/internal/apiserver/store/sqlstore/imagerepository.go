package sqlstore

import (
	Constants "github.com/Gugush284/Go-server.git/internal/apiserver"
	ModelImage "github.com/Gugush284/Go-server.git/internal/apiserver/model/image"
)

type ImageRepository struct {
	store *SqlStore
}

func (r *ImageRepository) Upload(i *ModelImage.Image) error {
	if err := i.Validate(); err != nil {
		return err
	}

	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return err
		}
	}

	statement, err := r.store.Db.Exec(
		"INSERT INTO images (image, image_name, txt) VALUES (?, ?, ?)",
		i.Image,
		i.ImageName,
		i.Txt,
	)
	if err != nil {
		return err
	}

	id, err := statement.LastInsertId()
	if err != nil {
		return err
	}
	if id == 0 {
		return Constants.ErrSqlIdNil
	}

	i.ImageId = int(id)

	return nil
}
