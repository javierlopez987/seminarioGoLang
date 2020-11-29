package flight

import (
	"net/http"
	"strconv"
	"fmt"

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
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"flight": s.FindByID(id),
		})
	}
}

// Register ...
func (s httpService) Register(r *gin.Engine)  {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}