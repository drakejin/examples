package main

import (
	"fmt"
	"net"

	pbhttp "github.com/drakejin/examples/protobuf-on-http-practice/back/proto"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/protobuf/proto"
)

const (
	headerContentType = "application/protobuf"
)

func main() {
	f := fiber.New()

	f.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
	}))

	f.Post("/", func() fiber.Handler {
		return func(ctx *fiber.Ctx) error {
			contentType := ctx.Get(fiber.HeaderContentType)
			if contentType != headerContentType {
				return ctx.SendStatus(fiber.StatusBadRequest)
			}

			req, err := parseRequest(ctx)
			if err != nil {
				return err
			}

			res := doSomeLogic(req)

			fmt.Println("----------- Request ----------------- ")
			fmt.Println(req)
			fmt.Println("----------- Response ----------------- ")
			fmt.Println(res)

			return sendResponse(ctx, res)
		}
	}())

	host := net.JoinHostPort("0.0.0.0", "4100")
	if err := f.Listen(host); err != nil {
		fmt.Println(err)
	}
}
func doSomeLogic(req *pbhttp.Request) *pbhttp.Response {
	res := &pbhttp.Response{}
	if req.ReqA != nil {
		res.ResA = &pbhttp.ResponseA{}
		res.ResA.Query = fmt.Sprintf("%s%s", req.ReqA.Query, "-ReqA")
	}
	if req.ReqB != nil {
		res.ResB = &pbhttp.ResponseB{}
		res.ResB.Query = fmt.Sprintf("%s%s", req.ReqB.Query, "-ReqB")
	}
	if req.ReqC != nil {
		res.ResC = &pbhttp.ResponseC{}
		res.ResC.Query = fmt.Sprintf("%s%s", req.ReqC.Query, "-ReqC")
	}
	// if req.ReqD != nil {
	// 	res.ResD = &pbhttp.ResponseD{}
	// 	res.ResD.Query = fmt.Sprintf("%s%s", req.ReqD.Query, "-ReqD")
	// }
	return res
}

func parseRequest(ctx *fiber.Ctx) (*pbhttp.Request, error) {
	req := &pbhttp.Request{}
	if err := proto.Unmarshal(ctx.Body(), req); err != nil {
		return nil, err
	}

	return req, nil
}

func sendResponse(ctx *fiber.Ctx, res *pbhttp.Response) error {
	b, err := proto.Marshal(res)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	ctx.
		Status(fiber.StatusOK).
		Set(fiber.HeaderContentType, headerContentType)
	return ctx.Send(b)
}