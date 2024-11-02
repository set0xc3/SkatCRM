package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/set0xc3/htmx/internal/db"

	"github.com/pocketbase/dbx"
	pb "github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

func SetRecordClientData(record *models.Record, client db.Client) {
	record.Set("id2", client.Id2)
	record.Set("mark", client.Mark)
	record.Set("contractor", client.Contractor)
	record.Set("full_name", client.FullName)
	record.Set("type", client.Type)
	record.Set("phones", client.Phones)
	record.Set("email", client.Email)
	record.Set("legal_address", client.LegalAddress)
	record.Set("physical_address", client.PhysicalAddress)
	record.Set("registration_date", client.RegistrationDate)
	record.Set("ad_channel", client.AdChannel)
	record.Set("reg_data_1", client.RegData1)
	record.Set("reg_data_2", client.RegData2)
	record.Set("note", client.Note)
	record.Set("request_count", client.RequestCount)
	record.Set("birthday", client.Birthday)
	record.Set("income", client.Income)
}

func GetProducts(ctx *pb.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		query := ctx.Dao().RecordQuery("product")
		records := []models.Record{}

		if err := query.All(&records); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Not Found"})
		}

		data := []map[string]any{}
		for _, record := range records {
			data = append(data, record.PublicExport())
		}

		return c.JSON(http.StatusOK, data)
	}
}

func GetProduct(ctx *pb.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.PathParam("id")
		record, err := ctx.Dao().FindRecordById("product", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Not Found"})
		}

		data := map[string]any{}
		data = record.PublicExport()

		return c.JSON(http.StatusOK, data)
	}
}

func GetClients(ctx *pb.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		query := ctx.Dao().RecordQuery("client")
		records := []models.Record{}

		if err := query.All(&records); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Not Found"})
		}

		data := []map[string]any{}
		for _, record := range records {
			data = append(data, record.PublicExport())
		}

		return c.JSON(http.StatusOK, data)
	}
}

func GetClient(ctx *pb.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.PathParam("id")
		record, err := ctx.Dao().FindRecordById("client", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Not Found"})
		}

		data := map[string]any{}
		data = record.PublicExport()

		return c.JSON(http.StatusOK, data)
	}
}

func PostAddClient(ctx *pb.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		var client db.Client

		if err := c.Bind(&client); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		query := ctx.Dao().RecordQuery("client")
		record := models.Record{}
		if err := query.One(&record); err != nil {
			collection, err := ctx.Dao().FindCollectionByNameOrId("client")
			if err != nil {
				return err
			}

			record := models.NewRecord(collection)
			SetRecordClientData(record, client)

			if err := ctx.Dao().SaveRecord(record); err != nil {
				return c.NoContent(http.StatusNoContent)
			}
			return c.NoContent(http.StatusInternalServerError)
		}

		client2 := db.Client{}
		err := ctx.Dao().DB().
			NewQuery("SELECT id2 FROM client WHERE id2 = {:id2}").
			Bind(dbx.Params{"id2": client.Id2}).
			One(&client2)
		if err != nil {
			collection, err := ctx.Dao().FindCollectionByNameOrId("client")
			if err != nil {
				return err
			}

			record := models.NewRecord(collection)
			SetRecordClientData(record, client)

			if err := ctx.Dao().SaveRecord(record); err != nil {
				return c.NoContent(http.StatusNoContent)
			}

			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Not Found"})
		}

		return c.String(http.StatusOK, "")
	}
}

func DeleteClient(ctx *pb.PocketBase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.PathParam("id")

		// Выполняем запрос на удаление клиента
		_, err := ctx.Dao().DB().
			NewQuery("DELETE FROM client WHERE id = {:id}").
			Bind(dbx.Params{"id": id}).
			Execute()

		if err != nil {
			// Логируем ошибку для диагностики
			log.Printf("Error deleting client with id %s: %v\n", id, err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error deleting client"})
		}

		return c.NoContent(http.StatusNoContent)
	}
}
