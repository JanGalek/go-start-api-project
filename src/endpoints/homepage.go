package endpoints

import (
	"encoding/json"
	"jan-galek/api/src/libs/database"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	err, connection := database.GetConnection()

	if err != nil {
		panic("DB ERROR")
	}
	panic(connection)
	json.NewEncoder(w).Encode("ok")
}
