package repositories

import "lang-portal-backend/internal/models"

type WordRepository struct {
	// Add necessary fields
}

func (r *WordRepository) GetWords() ([]models.Word, error) {
	// Implement the logic to get words
	return nil, nil
}

func (r *WordRepository) GetWordByID(id int) (*models.Word, error) {
	// Implement the logic to get a word by ID
	return nil, nil
}
