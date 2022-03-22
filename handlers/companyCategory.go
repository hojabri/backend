package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/hojabri/backend/models"
	"github.com/hojabri/backend/repository"
	"net/http"
	"strconv"
)

var companyCategoryRepository repository.CompanyCategoryRepository

func init() {
	companyCategoryRepository = repository.NewCompanyCategoryRepository()
}

// GetAllCompanyCategories gets all repository information
func GetAllCompanyCategories(c *fiber.Ctx) error {
	companyCategories := companyCategoryRepository.FindAll()
	
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    companyCategories,
		Title:   "GetAllCompanyCategories",
		Message: "All CompanyCategories",
	}
	
	return c.Status(resp.Code).JSON(resp)
}

// GetSingleCompanyCategory Gets single companyCategory information
func GetSingleCompanyCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in getting companyCategory information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	companyCategory, err := companyCategoryRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting companyCategory information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	if companyCategory == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("companyCategory with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding companyCategory",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    companyCategory,
		Title:   "OK",
		Message: "CompanyCategory information",
	}
	return c.Status(resp.Code).JSON(resp)
	
}

// AddNewCompanyCategory adds new companyCategory
func AddNewCompanyCategory(c *fiber.Ctx) error {
	companyCategory := &models.CompanyCategory{}
	
	err := c.BodyParser(companyCategory)
	
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in parsing companyCategory body information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	id, err := companyCategoryRepository.Save(*companyCategory)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in adding new companyCategory",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	companyCategory, err = companyCategoryRepository.FindByID(id)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly added companyCategory",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	if companyCategory == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("companyCategory with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding companyCategory",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    companyCategory,
		Title:   "OK",
		Message: "new companyCategory added successfully",
	}
	return c.Status(resp.Code).JSON(resp)
	
}

// UpdateCompanyCategory updates a companyCategory by companyCategory id
func UpdateCompanyCategory(c *fiber.Ctx) error {
	companyCategory := &models.CompanyCategory{}
	
	err := c.BodyParser(companyCategory)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in parsing companyCategory body information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in parsing companyCategory ID. (it should be an integer)",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	updatingCompanyCategory, err := companyCategoryRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting companyCategory information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	if updatingCompanyCategory == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("companyCategory with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding companyCategory",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	companyCategory.ID = uint(id)
	
	err = companyCategoryRepository.Update(*companyCategory)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in updating companyCategory information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	companyCategory, err = companyCategoryRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly updated companyCategory",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	if companyCategory == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("companyCategory with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding companyCategory",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    companyCategory,
		Title:   "UpdateCompanyCategory",
		Message: "companyCategory updated successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}

// DeleteCompanyCategory deletes the companyCategory from db
func DeleteCompanyCategory(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in getting companyCategory information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	companyCategory, err := companyCategoryRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding companyCategory",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	if companyCategory == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("companyCategory with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding companyCategory",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	err = companyCategoryRepository.Delete(*companyCategory)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in deleting companyCategory object",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    "companyCategory deleted successfully",
		Title:   "OK",
		Message: "companyCategory deleted successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}
