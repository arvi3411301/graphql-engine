package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hasura/graphql-engine/cli/migrate"
	"github.com/hasura/graphql-engine/cli/migrate/cmd"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

const (
	DataAPIError  = "Data Error: "
	MigrationMode = "migration_mode"
)

type Response struct {
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
	Name       string `json:"name,omitempty"`
	StatusCode int    `json:"-"`
}

type Request struct {
	Name string        `json:"name"`
	Up   []interface{} `json:"up"`
	Down []interface{} `json:"down"`
}

func MigrateAPI(c *gin.Context) {
	migratePtr, ok := c.Get("migrate")
	if !ok {
		return
	}
	// Get File url
	sourcePtr, ok := c.Get("filedir")
	if !ok {
		return
	}

	// Get Logger
	loggerPtr, ok := c.Get("logger")
	if !ok {
		return
	}

	metadataFilePtr, ok := c.Get("metadataFile")
	if !ok {
		return
	}

	t := migratePtr.(*migrate.Migrate)
	sourceURL := sourcePtr.(*url.URL)
	logger := loggerPtr.(*logrus.Logger)
	metadataFile := metadataFilePtr.(string)

	// Switch on request method
	switch c.Request.Method {
	case "GET":
		// Rescan file system
		err := t.ReScan()
		if err != nil {
			c.JSON(http.StatusInternalServerError, &Response{Code: "internal_error", Message: err.Error()})
			return
		}
		status, err := t.GetStatus()
		if err != nil {
			c.JSON(http.StatusInternalServerError, &Response{Code: "internal_error", Message: "Something went wrong"})
			return
		}
		c.JSON(http.StatusOK, status)
	case "POST":
		var request Request

		// Bind Request body to Request struct
		if c.BindJSON(&request) != nil {
			c.JSON(http.StatusInternalServerError, &Response{Code: "internal_error", Message: "Something went wrong"})
			return
		}

		startTime := time.Now()
		timestamp := startTime.UnixNano() / int64(time.Millisecond)

		createOptions := cmd.New(timestamp, request.Name, sourceURL.Path)
		err := createOptions.SetMetaUp(request.Up)
		if err != nil {
			c.JSON(http.StatusInternalServerError, &Response{Code: "create_file_error", Message: err.Error()})
			return
		}
		err = createOptions.SetMetaDown(request.Down)
		if err != nil {
			c.JSON(http.StatusInternalServerError, &Response{Code: "create_file_error", Message: err.Error()})
			return
		}

		err = createOptions.Create()
		if err != nil {
			c.JSON(http.StatusInternalServerError, &Response{Code: "create_file_error", Message: err.Error()})
			return
		}

		defer func() {
			if err != nil {
				err = createOptions.Delete()
				if err != nil {
					logger.Debug(err)
				}
			}
		}()

		// Rescan file system
		err = t.ReScan()
		if err != nil {
			c.JSON(http.StatusInternalServerError, &Response{Code: "internal_error", Message: err.Error()})
			return
		}

		if err = t.Migrate(uint64(timestamp), "up"); err != nil {
			if strings.HasPrefix(err.Error(), DataAPIError) {
				c.JSON(http.StatusBadRequest, &Response{Code: "data_api_error", Message: strings.TrimPrefix(err.Error(), DataAPIError)})
				return
			}

			if err == migrate.ErrNoMigrationMode {
				c.JSON(http.StatusBadRequest, &Response{Code: "migration_mode_disabled", Message: err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, &Response{Code: "internal_error", Message: err.Error()})
			return
		}
		exportMetadata := t.GetMetadataState()
		if exportMetadata != nil {
			metadataByt, err := yaml.Marshal(exportMetadata)
			if err != nil {
				c.JSON(http.StatusInternalServerError, &Response{Code: "export_metadata_error", Message: err.Error()})
				return
			}
			err = ioutil.WriteFile(metadataFile, metadataByt, 0644)
			if err != nil {
				c.JSON(http.StatusInternalServerError, &Response{Code: "export_metadata_error", Message: err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, &Response{Name: fmt.Sprintf("%d_%s", timestamp, request.Name)})
	default:
		c.JSON(http.StatusMethodNotAllowed, &gin.H{"message": "Method not allowed"})
	}
}
