package stock

type Service interface {
	Create(stock *CreateRequest) (*Stock, error)
	Get(id string) (*Stock, error)
}

type serviceImpl struct {
	repository Repository
}

func NewService() (Service, error) {
	repository, err := newRepository()
	if err != nil {
		return nil, err
	}

	return &serviceImpl{
		repository,
	}, nil
}

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (s *serviceImpl) Create(sr *CreateRequest) (*Stock, error) {
	stock := NewStock(sr.Name)

	stock, err := s.repository.Insert(stock)
	if err != nil {
		return nil, err
	}

	return stock, nil
}

func (s *serviceImpl) Get(id string) (*Stock, error) {
	return s.repository.FindByID(id)
}
