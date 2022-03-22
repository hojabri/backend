package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/hojabri/backend/models"
	"github.com/hojabri/backend/repository"
	"net/http"
	"strconv"
)

var companyRepository repository.CompanyRepository

func init() {
	companyRepository = repository.NewCompanyRepository()
}

// GetAllCompanies gets all repository information
func GetAllCompanies(c *fiber.Ctx) error {
	companies := companyRepository.FindAll()
	
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    companies,
		Title:   "GetAllCompanies",
		Message: "All Companies",
	}
	
	return c.Status(resp.Code).JSON(resp)
}

// GetSingleCompany Gets single company information
func GetSingleCompany(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in getting company information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	company, err := companyRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting company information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	if company == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("company with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding company",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    company,
		Title:   "OK",
		Message: "Company information",
	}
	return c.Status(resp.Code).JSON(resp)
	
}

// AddNewCompany adds new company
func AddNewCompany(c *fiber.Ctx) error {
	company := &models.Company{}
	
	err := c.BodyParser(company)
	
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in parsing company body information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	id, err := companyRepository.Save(*company)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in adding new company",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	company, err = companyRepository.FindByID(id)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly added company",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	if company == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("company with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding company",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    company,
		Title:   "OK",
		Message: "new company added successfully",
	}
	return c.Status(resp.Code).JSON(resp)
	
}

// UpdateCompany updates a company by company id
func UpdateCompany(c *fiber.Ctx) error {
	company := &models.Company{}
	
	err := c.BodyParser(company)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in parsing company body information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in parsing company ID. (it should be an integer)",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	updatingCompany, err := companyRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    err.Error(),
			Title:   "NotFound",
			Message: "Error in getting company information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	if updatingCompany == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("company with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding company",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	company.ID = uint(id)
	
	err = companyRepository.Update(*company)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in updating company information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	company, err = companyRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding newly updated company",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	if company == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("company with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding company",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    company,
		Title:   "UpdateCompany",
		Message: "company updated successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}

// DeleteCompany deletes the company from db
func DeleteCompany(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)
	
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "Error",
			Message: "Error in getting company information",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	company, err := companyRepository.FindByID(uint(id))
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusInternalServerError,
			Body:    err.Error(),
			Title:   "InternalServerError",
			Message: "Error in finding company",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	if company == nil {
		errorResp := models.Response{
			Code:    http.StatusNotFound,
			Body:    fmt.Sprintf("company with id %d could not be found", id),
			Title:   "NotFound",
			Message: "Error in finding company",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	err = companyRepository.Delete(*company)
	if err != nil {
		errorResp := models.Response{
			Code:    http.StatusNotAcceptable,
			Body:    err.Error(),
			Title:   "NotAcceptable",
			Message: "Error in deleting company object",
		}
		
		return c.Status(errorResp.Code).JSON(errorResp)
	}
	
	resp := models.Response{
		Code:    http.StatusOK,
		Body:    "company deleted successfully",
		Title:   "OK",
		Message: "company deleted successfully",
	}
	return c.Status(resp.Code).JSON(resp)
}
