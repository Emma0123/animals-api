package handlers

import (
	"animals-api/animal"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type animalHandler struct {
	animalService animal.Service
}

func NewAnimalHandler(animalService animal.Service) *animalHandler {
	return &animalHandler{animalService}
}

func (h *animalHandler) GetAllAnimalsDetail(c *gin.Context) {
	animals, err := h.animalService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var animalsResponse []animal.AnimalResponse

	for _, a := range animals {
		animalResponse := convertToAnimalResponse(a)

		animalsResponse = append(animalsResponse, animalResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": animalsResponse,
	})
}

func (h *animalHandler) GetOneAnimalDetail(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid ID format",
		})
		return
	}

	a, err := h.animalService.FindByID(int(id))

	if a.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
		return
	}

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	animalResponse :=
		convertToAnimalResponse(a)

	c.JSON(http.StatusOK, gin.H{
		"data": animalResponse,
	})

}

func (h *animalHandler) CreateAnimal(c *gin.Context) {
	var animalRequest animal.AnimalRequest

	err := c.ShouldBindJSON(&animalRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return

	}

	// Checking duplicate animal
	existingAnimal, err := h.animalService.FindByName(animalRequest.Name)
	if err == nil && existingAnimal.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"errors": "Animal with this name already exists",
		})
		return
	}

	createdAnimal, err := h.animalService.Create(animalRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    convertToAnimalResponse(createdAnimal),
		"message": "Animal details added successfully",
	})
}

func (h *animalHandler) UpdateAnimal(c *gin.Context) {
	var updateAnimalRequest animal.UpdateAnimalRequest

	err := c.ShouldBindJSON(&updateAnimalRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	currentAnimal, err := h.animalService.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newAnimal, err := h.animalService.Create(animal.AnimalRequest{
				Name:  updateAnimalRequest.Name,
				Class: updateAnimalRequest.Class,
				Legs:  updateAnimalRequest.Legs,
			})

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"errors": err,
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"data": convertToAnimalResponse(newAnimal),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	if updateAnimalRequest.Name != "" {
		currentAnimal.Name = updateAnimalRequest.Name
	}
	if updateAnimalRequest.Class != "" {
		currentAnimal.Class = updateAnimalRequest.Class
	}
	if updateAnimalRequest.Legs != 0 {
		currentAnimal.Legs = updateAnimalRequest.Legs
	}

	updatedAnimal, err := h.animalService.Update(id, animal.UpdateAnimalRequest{
		Name:  currentAnimal.Name,
		Class: currentAnimal.Class,
		Legs:  currentAnimal.Legs,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    convertToAnimalResponse(updatedAnimal),
		"message": "Animal details updated successfully",
	})
}

func (h *animalHandler) DeleteAnimalDetail(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "Invalid ID format",
		})
		return
	}

	a, err := h.animalService.Delete(int(id))

	if a.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
		return
	}

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})

		return
	}

	animalResponse :=
		convertToAnimalResponse(a)

	c.JSON(http.StatusOK, gin.H{
		"data":    animalResponse,
		"message": "Animal details deleted successfully",
	})

}

func convertToAnimalResponse(a animal.Animals) animal.AnimalResponse {
	return animal.AnimalResponse{
		ID:    a.ID,
		Name:  a.Name,
		Class: a.Class,
		Legs:  a.Legs,
	}
}
