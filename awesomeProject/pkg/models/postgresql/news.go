package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"lina.net/aitunewstask/pkg/models"
)

type ModelNews struct {
	DB *sql.DB
}

func (m *ModelNews) Insert(title, content, expires, category string) (int, error) {

	stmt := `INSERT INTO news (title, content, created, expires, category)
       VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + $3::interval, $4)
       RETURNING id`

	fmt.Println("Query:", stmt)
	fmt.Println("Values:", title, content, expires, category)
	row := m.DB.QueryRow(stmt, title, content, expires, category)
	var id int

	err := row.Scan(&id)
	if err != nil {
		fmt.Println("Error scanning row:", err)
		return 0, err
	}
	return int(id), nil
}

func (m *ModelNews) Latest() ([]*models.News, error) {
	stmt := `SELECT id, title, content, created, expires, category FROM news
             WHERE expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	news := []*models.News{}

	for rows.Next() {
		s := &models.News{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires, &s.Category)
		if err != nil {
			return nil, err
		}

		news = append(news, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return news, nil
}
func (m *ModelNews) Get(id int) (*models.News, error) {
	stmt := `SELECT id, title, content, created, expires, category FROM news WHERE id = $1`
	row := m.DB.QueryRow(stmt, id)

	n := &models.News{}

	err := row.Scan(&n.ID, &n.Title, &n.Content, &n.Created, &n.Expires, &n.Category)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return n, nil
}
func (m *ModelNews) LatestForCreatePage(limit int) ([]*models.News, error) {
	stmt := `SELECT id, title, content, created, expires, category FROM news
             WHERE expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT $1`

	rows, err := m.DB.Query(stmt, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	news := []*models.News{}

	for rows.Next() {
		s := &models.News{}

		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires, &s.Category)
		if err != nil {
			return nil, err
		}

		news = append(news, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return news, nil
}
func (m *ModelNews) LatestByCategory(category string) ([]*models.News, error) {
	stmt := `SELECT id, title, content, created, expires, category FROM news
             WHERE category = $1 AND expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	news := []*models.News{}

	for rows.Next() {
		s := &models.News{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires, &s.Category)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		news = append(news, s)
		fmt.Println(news)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return news, nil
}
