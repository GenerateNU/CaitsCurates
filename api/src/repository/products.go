package repository

import (
 "CaitsCurates/backend/config"
 "pCaitsCurates/backend/model"

 _ "github.com/go-sql-driver/mysql"
)

func GetAllProducts(products *[]model.Products) (err error) {
 if err = config.DB.Find(products).Error; err != nil {
  return err
 }
 return nil
}

func CreateProduct(product *model.Products) (err error) {
 if err = config.DB.Create(product).Error; err != nil {
  return err
 }
 return nil
}

func GetAProduct(product *model.Products, id string) (model.Products, error) {
 err := config.DB.Where("id = ?", id).First(product).Error
 return *product, err
}