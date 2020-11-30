package flight

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method string
	path string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method: "GET",
		path: "/flights",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method: "GET",
		path: "/flights/:id",
		function: getOne(s),
	})

	list = append(list, &endpoint{
		method: "POST",
		path: "/flights",
		function: create(s),
	})

	list = append(list, &endpoint{
		method: "PUT",
		path: "/flights/:id",
		function: update(s),
	})

	list = append(list, &endpoint{
		method: "DELETE",
		path: "/flights/:id",
		function: delete(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"flights": s.FindAll(),
		})
	}
}

func getOne(s Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"flight": s.FindByID(id),
		})
	}
}

func create(s Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		var f Flight
		err := c.BindJSON(&f)
		if err != nil {
			panic(err)
		}
		s.Add(f)
		c.JSON(http.StatusCreated, gin.H{
			"message": "flight added",
		})
	}
}

func update(s Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		var f Flight
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		f.ID = id
		err = c.BindJSON(&f)
		if err != nil {
			panic(err)
		}
		s.Update(f)
		c.JSON(http.StatusOK, gin.H{
			"message": "ID " + i + " modified",
		})
	}
}

func delete(s Service) gin.HandlerFunc {
	return func (c *gin.Context) {
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		s.Delete(id)
		c.JSON(http.StatusOK, gin.H{
			"message": "ID " + i + " deleted",
		})
	}
}



// Register ...
func (s httpService) Register(r *gin.Engine)  {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}