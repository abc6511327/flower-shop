package services

import (
    "flower-shop/internal/models"
    "flower-shop/pkg/db"
)

type FlowerService struct {}

func (fs *FlowerService) GetAllFlowers() ([]models.Flower, error) {
    rows, err := db.DB.Query("SELECT id, name, description, price FROM flowers")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var flowers []models.Flower
    for rows.Next() {
        var f models.Flower
        if err := rows.Scan(&f.ID, &f.Name, &f.Description, &f.Price); err != nil {
            return nil, err
        }
        flowers = append(flowers, f)
    }
    return flowers, nil
}
