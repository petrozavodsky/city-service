package city

import (
	"sort"
)

// Storage - хранилище городов
type Storage struct {
	Ids             []int
	Store           map[int]*City
	RegionIndex     map[string][]*City
	DistrictIndex   map[string][]*City
	PopulationIndex map[int][]*City
	FoundationIndex map[int][]*City
}

// MakeStorage - создает экземпляр хранилище
func MakeStorage() *Storage {
	return &Storage{
		make([]int, 0, 0),
		make(map[int]*City),
		make(map[string][]*City),
		make(map[string][]*City),
		make(map[int][]*City),
		make(map[int][]*City),
	}
}

// AddCity создание грода
func (s *Storage) AddCity(id int, city *City) {
	s.Ids = append(s.Ids, id)
	s.Store[id] = city

	s.generateIndex(id, city)
}

// DeleteCity удаление города по id
func (s *Storage) DeleteCity(id int) {

	s.Ids = make([]int, 0, 0)
	delete(s.Store, id)

	s.RegionIndex = map[string][]*City{}
	s.DistrictIndex = map[string][]*City{}
	s.PopulationIndex = map[int][]*City{}
	s.FoundationIndex = map[int][]*City{}

	for id := range s.Store {
		s.Ids = append(s.Ids, id)
		s.generateIndex(id, s.Store[id])
	}
}

// GetCity получениегорода по id
func (s *Storage) GetCity(id int) *City {

	return s.Store[id]
}

// GetByDistrict получение городов по району
func (s *Storage) GetByDistrict(key string) []*City {

	return s.DistrictIndex[key]
}

// GetByRegion получение городов по региону
func (s *Storage) GetByRegion(key string) []*City {

	return s.RegionIndex[key]
}

// GetByFoundation получение городов по интервалу численности
func (s *Storage) GetByFoundation(from int, to int) []*City {

	var out []*City

	for i := from; i <= to; i++ {
		if s.FoundationIndex[i] != nil {
			out = append(out, s.FoundationIndex[i]...)
		}
	}

	return out
}

// GetByPopulation получение городов по интервалу численности
func (s *Storage) GetByPopulation(from int, to int) []*City {

	var out []*City

	for i := from; i <= to; i++ {
		if s.PopulationIndex[i] != nil {
			out = append(out, s.PopulationIndex[i]...)
		}
	}

	return out
}

// GenerateId создание списка id
func (s *Storage) GenerateId() int {

	sort.Slice(s.Ids, func(i, j int) bool {
		return s.Ids[i] < s.Ids[j]
	})

	return s.Ids[len(s.Ids)-1] + 1
}

// generateIndex Создание индексов
func (s *Storage) generateIndex(id int, city *City) {

	s.RegionIndex[city.Region] = append(s.RegionIndex[city.Region], s.Store[id])
	s.DistrictIndex[city.District] = append(s.DistrictIndex[city.District], s.Store[id])
	s.PopulationIndex[city.Population] = append(s.PopulationIndex[city.Population], s.Store[id])
	s.FoundationIndex[city.Foundation] = append(s.FoundationIndex[city.Foundation], s.Store[id])
}

// UpdatePopulation Обновление численности
func (s *Storage) UpdatePopulation(id int, city *City) {

	s.DeleteCity(id)

	s.Ids = append(s.Ids, id)
	s.Store[id] = city

	s.generateIndex(id, city)
}

// StoreToStrings сериализация в строку
func (s *Storage) StoreToStrings() [][]string {
	var data [][]string

	for i := range s.Store {

		data = append(data, s.Store[i].ToStrings())
	}

	return data
}
