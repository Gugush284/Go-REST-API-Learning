package sqlstore

import (
	"database/sql"

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

func (r *ImageRepository) Download(id int) (*ModelImage.Image, error) {
	if err := r.store.Db.Ping(); err != nil {
		if err := r.store.Open(); err != nil {
			return nil, err
		}
	}

	i := &ModelImage.Image{}

	row := r.store.Db.QueryRow(
		"SELECT image, image_name, txt FROM images WHERE image_id = (?)",
		id)
	if err := row.Scan(&i.Image, &i.ImageName, &i.Txt); err != nil {
		if err == sql.ErrNoRows {
			return nil, Constants.ErrRecordNotFound
		}

		return nil, err
	}

	i.ImageId = id

	return i, nil
}
