module webp_server_go

go 1.15

require (
	github.com/bep/gowebp v0.1.0
	github.com/gofiber/fiber/v2 v2.4.0
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.3.0
	golang.org/x/image v0.0.0-20210220032944-ac19c3e999fb
)

replace (
	github.com/bep/gowebp v0.1.0 => github.com/webp-sh/gowebp v0.1.0
	github.com/gofiber/fiber/v2 v2.4.0 => github.com/webp-sh/fiber/v2 v2.4.0
)
