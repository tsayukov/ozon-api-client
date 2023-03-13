package ozon

import (
	"net/http"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
)

func TestGetListOfWarehouses(t *testing.T) {
	tests := []struct {
		statusCode int
		headers    map[string]string
		response   string
	}{
		{
			http.StatusOK,
			map[string]string{"Client-Id": "my-client-id", "Api-Key": "my-api-key"},
			`{
				"result": [
				  {
					"warehouse_id": 15588127982000,
					"name": "Proffi (Панорама Групп)",
					"is_rfbs": false
				  },
				  {
					"warehouse_id": 22142605386000,
					"name": "Склад на производственной",
					"is_rfbs": true
				  },
				  {
					"warehouse_id": 22208673494000,
					"name": "Тест 37349",
					"is_rfbs": true
				  },
				  {
					"warehouse_id": 22240462819000,
					"name": "Тест12",
					"is_rfbs": true
				  }
				]
			}`,
		},
	}

	for _, test := range tests {
		c := NewMockClient(core.NewMockHttpHandler(test.statusCode, test.response, test.headers))

		resp, err := c.GetListOfWarehouses()
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("got wrong status code: got: %d, expected: %d", resp.StatusCode, test.statusCode)
		}
	}
}
