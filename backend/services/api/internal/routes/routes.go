package routes

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Azat201003/eduflow/backend/libs/gen/go/filager"
	"github.com/Azat201003/eduflow/backend/libs/gen/go/summary"
	"github.com/Azat201003/eduflow/backend/libs/gen/go/user"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
)

type SummarySearchBind struct {
	Title       string `query:"t"`
	Description string `query:"d"`
	AuthorId    uint64 `query:"a"`
}

type Report struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var UserClient user.UserServiceClient
var SummaryClient summary.SummaryServiceClient
var FilagerClient filager.FileManagerServiceClient

func SummaryList(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("p"))
	if err != nil {
		fmt.Println("a")
		return c.Redirect(http.StatusPermanentRedirect, "?p=1")
	}

	size_str := c.QueryParam("s")
	size := 5
	if size_str != "" {
		size, err = strconv.Atoi(size_str)
		if err != nil {
			return c.JSON(http.StatusBadRequest, Report{
				Ok:      false,
				Message: "Cannot convert size to int.",
			})
		}
	}

	stream, err := SummaryClient.GetFilteredSummaries(context.Background(), &summary.FilterRequest{
		Filter: &summary.Summary{},
		Page: &summary.Page{
			Number: uint32(page),
			Size:   uint32(size),
		},
	})

	var summaries []map[string]interface{}
	for {
		summary, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return c.JSON(http.StatusInternalServerError, Report{
				Ok:      false,
				Message: "Failed to get summaries.",
			})
		}
		summaries = append(summaries, map[string]interface{}{
			"id":          summary.Id.Id,
			"title":       summary.Title,
			"description": summary.Description,
			"author_id":   summary.AuthorId.Id,
			"file_path":   summary.FilePath,
		})
	}

	return c.JSON(http.StatusFound, Report{
		Ok:   true,
		Data: summaries,
	})
}

func MainRoute(c echo.Context) error {
	name := "world"
	if c.Get("is_user_setted").(bool) {
		authed_user := c.Get("user").(*user.User)
		name = authed_user.GetUsername()
	}
	return c.JSON(http.StatusOK, Report{
		Data: fmt.Sprintf("Hello, %s!", name),
		Ok:   true,
	})
}

func SignUp(c echo.Context) error {
	fmt.Println("saf")
	c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	username := c.FormValue("username")
	password := c.FormValue("password")
	fmt.Println(username, password)
	token, err := UserClient.Register(context.Background(), &user.Creditionals{Username: username, Password: password})

	if s, ok := status.FromError(err); !ok {
		return c.JSON(http.StatusNotFound, Report{
			Ok:      false,
			Message: s.Message(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]string{"token": token.Token})
}

func SignIn(c echo.Context) error {
	fmt.Println("af")
	c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	username := c.FormValue("username")
	password := c.FormValue("password")
	var bytes []byte
	c.Request().Body.Read(bytes)
	fmt.Println(c.FormParams())
	fmt.Println(username, password, bytes)
	token, err := UserClient.Login(context.Background(), &user.Creditionals{Username: username, Password: password})
	if s, _ := status.FromError(err); err != nil {
		return c.JSON(http.StatusNotFound, &Report{
			Ok:      false,
			Message: s.Message(),
		})
	}
	return c.JSON(http.StatusOK, &Report{
		Ok:   true,
		Data: map[string]string{"token": token.Token},
	})
}

func MyUser(c echo.Context) error {
	c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	if c.Get("is_user_setted").(bool) {
		authed_user := c.Get("user").(*user.User)
		return c.JSON(http.StatusOK, Report{
			Data: map[string]interface{}{
				"username": authed_user.Username,
				"id":       authed_user.Id.Id,
				"is_staff": authed_user.IsStaff,
			},
			Ok: true,
		})
	}
	return c.JSON(http.StatusUnauthorized, Report{
		Ok:      false,
		Message: "Cannot find user with this token",
	})
}

func GetUserById(c echo.Context) error {
	c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnsupportedMediaType, Report{
			Ok:      false,
			Message: "Cannot convert id to int.",
		})
	}
	user_obj, err := UserClient.GetUserById(context.Background(), &user.Id{Id: uint64(id)})
	if err != nil {
		return c.JSON(http.StatusNotFound, Report{
			Ok:      false,
			Message: "Cannot find user.",
		})
	}
	return c.JSON(http.StatusFound, Report{
		Ok: true,
		Data: map[string]interface{}{
			"username": user_obj.Username,
			"id":       user_obj.Id.Id,
			"is_staff": user_obj.IsStaff,
		},
	})
}

func GetSummaryById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnsupportedMediaType, Report{
			Ok:      false,
			Message: "Cannot convert id to int.",
		})
	}

	summary, err := SummaryClient.GetSummaryById(context.Background(), &summary.Id{Id: uint64(id)})
	if err != nil {
		return c.JSON(http.StatusNotFound, Report{
			Ok:      false,
			Message: "Cannot find summary.",
		})
	}
	return c.JSON(http.StatusFound, Report{
		Ok: true,
		Data: map[string]interface{}{
			"title":       summary.Title,
			"description": summary.Description,
			"author_id":   summary.AuthorId.Id,
			"file_path":   summary.FilePath,
		},
	})
}

func StartFileSending(c echo.Context) error {
	var req filager.StartWriteRequest
	err := (&echo.DefaultBinder{}).BindBody(c, &req)
	fmt.Println(req, err)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Report{
			Ok:      false,
			Message: "Cannot read request.",
		})
	}
	resp, err := FilagerClient.StartSending(context.Background(), &req)
	fmt.Println(resp, err)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Report{
			Ok:      false,
			Message: fmt.Sprintf("Cannot start sending file: %v", err),
		})
	}

	return c.JSON(http.StatusOK, Report{
		Ok:   true,
		Data: resp,
	})
}

func SendChunk(c echo.Context) error {
	var req filager.WriteChunk
	err := (&echo.DefaultBinder{}).BindBody(c, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Report{
			Ok:      false,
			Message: "Cannot read request.",
		})
	}
	_, err = FilagerClient.SendChunk(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Report{
			Ok:      false,
			Message: fmt.Sprintf("Cannot send chunk: %v", err),
		})
	}

	return c.JSON(http.StatusOK, Report{
		Ok: true,
	})
}

func CloseFileSending(c echo.Context) error {

	var req filager.EndRequest
	err := (&echo.DefaultBinder{}).BindBody(c, &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Report{
			Ok:      false,
			Message: "Cannot read request.",
		})
	}
	_, err = FilagerClient.CloseSending(context.Background(), &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Report{
			Ok:      false,
			Message: fmt.Sprintf("Cannot close sending file: %v", err),
		})
	}

	return c.JSON(http.StatusOK, Report{
		Ok: true,
	})
}
