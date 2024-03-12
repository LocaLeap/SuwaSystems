package main

import (
	"context"
	"log"
	"suwasystem/backend/ent"

	"entgo.io/ent/dialect"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Client *ent.Client
	DBCtx  context.Context
}

func main() {
	// Create Database Client
	client, err := ent.Open(dialect.MySQL, "root:pass@tcp(localhost:3306)/wallet?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Run Database Auto Migration
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("faild creating schema resources: %v", err)
	}

	app := &App{Client: client, DBCtx: ctx}

	r := gin.Default()

	// User API
	r.POST("/login", app.login)
	r.POST("/signup", signup)
	r.GET("/logout", logout)

	// Wallet API
	r.GET("/wallet", getWallets)
	r.GET("/wallet/:waleltId", getWallet)
	r.POST("/wallet/:communityId", createWallet)
	r.DELETE("/wallet/:walletId", deleteWallet)

	// Community API
	r.GET("/community", getCommunities)
	r.GET("/community/:communityId", getCommunity)
	r.POST("/community/:communityId", createCommunity)

	// Payment API
	r.POST("/wallet/pay", walletPay)
	r.POST("/wallet/change", walletChange)

	// Transaction API
	r.GET("/transaction/", getTransactions)
	r.GET("/transaction/:transactionId", getTransaction)

	// Application Start
	r.Run()
}

func (app App) login(c *gin.Context) {
	task1, _ := app.Client.User.Create().SetType("CARD").Save(app.DBCtx)
	c.JSON(200, gin.H{
		"message": "login",
		"task1":   task1,
	})
}
func signup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "signup",
	})
}
func logout(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "logout",
	})
}

func getWallets(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getWallets",
	})
}
func getWallet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getWallet",
	})
}
func createWallet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "createWallet",
	})
}
func deleteWallet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "deleteWallet",
	})
}

func getCommunities(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getCommunities",
	})
}
func getCommunity(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getCommunity",
	})
}
func createCommunity(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "createCommunity",
	})
}

func walletPay(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "WalletPay",
	})
}
func walletChange(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "WalletChange",
	})
}

func getTransactions(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getTransactions",
	})
}
func getTransaction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "getTransaction",
	})
}
