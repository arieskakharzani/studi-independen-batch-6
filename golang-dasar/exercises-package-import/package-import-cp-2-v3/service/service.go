package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	if quantity <= 0 {
		return errors.New("invalid quantity")
	}
	product, err := s.database.GetProductByName(productName)
	if err != nil {
		return err
	}

	cartItems, _ := s.database.GetCartItems()

	// Mengecek apakah produk sudah ada dalam cart
	for i, item := range cartItems {
		if item.ProductName == productName {
			cartItems[i].Quantity += quantity
			err := s.database.SaveCartItems(cartItems)
			if err != nil {
				return err
			}
			return nil
		}
	}

	newCartItem := entity.CartItem{
		ProductName: productName,
		Price:       product.Price,
		Quantity:    quantity,
	}
	cartItems = append(cartItems, newCartItem)
	err = s.database.SaveCartItems(cartItems)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) RemoveCart(productName string) error {
	cartItems, _ := s.database.GetCartItems()

	for i, item := range cartItems {
		if item.ProductName == productName {
			cartItems = append(cartItems[:i], cartItems[i+1:]...)
			err := s.database.SaveCartItems(cartItems)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("product not found")
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.GetCartItems()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {
	var emptyCart []entity.CartItem
	err := s.database.SaveCartItems(emptyCart)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	return s.database.GetProductData(), nil
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	totalPrice := 0
	for _, item := range cartItems {
		totalPrice += item.Price * item.Quantity
	}

	if money < totalPrice {
		return entity.PaymentInformation{}, errors.New("money is not enough")
	}

	change := money - totalPrice

	paymentInfo := entity.PaymentInformation{
		ProductList: cartItems,
		TotalPrice:  totalPrice,
		MoneyPaid:   money,
		Change:      change,
	}

	// Reset the cart after payment
	err = s.ResetCart()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	return paymentInfo, nil
}
