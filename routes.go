package main

func initializeRoutes() {

	// Handle the index route
	router.GET("/", showIndexPage)

	// GET : Article 상세 뷰
	router.GET("/article/view/:article_id", getArticle)

}
