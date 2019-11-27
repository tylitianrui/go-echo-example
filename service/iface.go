package service

type IService interface {
	Get(id uint) (data interface{}, err error)
	Update(interface{}) error
	Create(data interface{}) error
	Delete()
}
