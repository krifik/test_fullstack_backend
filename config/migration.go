package config

import "github.com/krifik/test_fullstack_backend/entity"

type Entity struct {
	Entity interface{}
}

func RegisterEntities() []Entity {
	return []Entity{
		{Entity: entity.Student{}},
		{Entity: entity.User{}},
		// {Entity: entity.Order{}},
		// {Entity: entity.Driver{}},
		// {Entity: entity.Payment{}},
		// {Entity: entity.Product{}},
		// {Entity: entity.ProductImage{}},
		// {Entity: entity.Category{}},
		// {Entity: entity.Partner{}},
	}
}
