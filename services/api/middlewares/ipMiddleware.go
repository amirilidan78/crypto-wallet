package middlewares

import (
	"crypto-wallet/config"
	"github.com/gin-gonic/gin"
)

type IpNotAllowedResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const IpNotAllowed = "this ip is not in white list "

func IpMiddleware(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		//allowedIpsAndNetworks := config.GetArray("server.ips")
		//clientIp := c.ClientIP()
		//ip := net.ParseIP(clientIp)
		//
		//r := IpNotAllowedResponse{
		//	http.StatusForbidden,
		//	IpNotAllowed + clientIp,
		//}
		//
		//for _, ipOrNetwork := range allowedIpsAndNetworks {
		//	_, subnet, _ := net.ParseCIDR(ipOrNetwork)
		//	if subnet.Contains(ip) {

		//	}
		//}

		//c.AbortWithStatusJSON(http.StatusForbidden, r)

		c.Next()
		return
	}
}
