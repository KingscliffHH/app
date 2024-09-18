package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Infinities-ICT-Solutions/project-dashboard/data"
	st "github.com/Infinities-ICT-Solutions/project-dashboard/internal/storage"
)

type benchmarkHandler struct {
	storage st.BenchmarkStorage
}

func NewBenchmarkHandler(storage st.BenchmarkStorage) *benchmarkHandler {
	return &benchmarkHandler{storage: storage}
}

func (p *benchmarkHandler) ListBenchmarks(c echo.Context) error {
	res, err := p.storage.GetAll()
	if err != nil {
		return c.String(500, "error getting benchmarks")
	}

	return c.JSON(http.StatusOK, res)
}

func (p *benchmarkHandler) GetBenchmark(c echo.Context) error {
	id := c.Param("id")
	res, err := p.storage.GetById(id)
	if err != nil {
		return &echo.HTTPError{Code: 404, Message: "benchmark not found"}
	}

	return c.JSON(http.StatusOK, res)
}

func (p *benchmarkHandler) CreateBenchmark(c echo.Context) error {
	// uses echo bind to get form values
	benchmark := &data.Benchmark{}
	c.Bind(benchmark)
	errors, err := benchmark.Validate()
	if err != nil {
		return c.JSON(400, errors)
	}

	res, err := p.storage.Create(benchmark)
	if err != nil {
		return c.String(500, "error creating benchmark")
	}

	return c.JSON(http.StatusCreated, res)
}

func (p *benchmarkHandler) UpdateBenchmark(c echo.Context) error {
	id := c.Param("id")
	benchmark := &data.Benchmark{}
	c.Bind(benchmark)
	errors, err := benchmark.Validate()
	if err != nil {
		return c.JSON(400, errors)
	}

	res, err := p.storage.Update(id, benchmark)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(http.StatusAccepted, res)
}

func (p *benchmarkHandler) DeleteBenchmark(c echo.Context) error {
	id := c.Param("id")
	err := p.storage.Delete(id)
	if err != nil {
		return c.String(500, "error deleting benchmark")
	}

	return c.JSON(http.StatusNoContent, nil)
}

// func (p *benchmarkHandler) UpdateBenchmark(c echo.Context) error {
// 	id := c.Param("id")
// 	benchmark := &data.Benchmark{}
// 	c.Bind(benchmark)

// 	err := p.storage.Update(id, benchmark)
// 	if err != nil {
// 		return c.String(500, "error updating benchmark")
// 	}

// 	return c.Redirect(http.StatusSeeOther, "/benchmarks/")
// }

// func (p *benchmarkHandler) NewBenchmark(c echo.Context) error {
// 	isHTMX := c.Request().Header.Get("HX-Request")

// 	if isHTMX == "true" {
// 		return utils.Render(c, 200, benchmark.PartialCreate())
// 	}

// 	return utils.Render(c, 200, benchmark.Create())
// }
