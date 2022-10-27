package city

import (
	"strconv"
)

type City struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	District   string `json:"district"`
	Population int    `json:"population"`
	Foundation int    `json:"foundation"`
}

// SetId установка значения
func (c *City) SetId(id int) {
	c.Id = id
}

// SetName установка значения
func (c *City) SetName(name string) {
	c.Name = name
}

// SetRegion установка значения
func (c *City) SetRegion(region string) {
	c.Region = region
}

// SetDistrict установка значения
func (c *City) SetDistrict(district string) {
	c.District = district
}

// SetPopulation установка значения
func (c *City) SetPopulation(population int) {
	c.Population = population
}

// SetFoundation установка значения
func (c *City) SetFoundation(foundation int) {
	c.Foundation = foundation
}

// GetId получение значения
func (c *City) GetId() int {
	return c.Id
}

// GetName получение значения
func (c *City) GetName() string {
	return c.Name
}

// GetRegion получение значения
func (c *City) GetRegion() string {
	return c.Region
}

// GetDistrict получение значения
func (c *City) GetDistrict() string {
	return c.District
}

// GetPopulation получение значения
func (c *City) GetPopulation() int {
	return c.Population
}

// GetFoundation получение значения
func (c *City) GetFoundation() int {
	return c.Foundation
}

// ToStrings сериализация в массив строк
func (c *City) ToStrings() []string {
	var values []string

	values = append(values, strconv.Itoa(c.Id))
	values = append(values, c.Name)
	values = append(values, c.Region)
	values = append(values, c.District)
	values = append(values, strconv.Itoa(c.Population))
	values = append(values, strconv.Itoa(c.Foundation))

	for i := range values {
		if len(values) < i {
			values[i] = values[i] + ","
		}
	}

	return values
}
