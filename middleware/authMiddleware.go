package middleware that
import(
	"fmt"
	"net/http"
	helper "jwt-auth/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error":fmt.Sprintf("No Authorization Header Provided")})			
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		
	}
}