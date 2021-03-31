package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cesararredondow/course/second_deliverable/routes/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

//mockgen -source=second_deliverable/routes/route.go -destination=second_deliverable/routes/mocks/route.go -package=mocks
func Test_New(t *testing.T) {
	testCases := []struct {
		name           string
		endpoint       string
		handlerName    string
		status         int
		callController bool
	}{
		{
			name:           "OK, Get square",
			endpoint:       "/api/v2/pokemons",
			handlerName:    "GetPokemonsFromApi",
			status:         200,
			callController: true,
		},
		{
			name:           "OK, Get square",
			endpoint:       "/api/v2/pokemon/1",
			handlerName:    "GetPokemonFromApi",
			status:         200,
			callController: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			r := mux.NewRouter()

			c := mocks.NewMockController(mockCtrl)

			if tc.callController && tc.handlerName == "GetPokemonsFromApi" {
				c.EXPECT().GetPokemonsFromApi(gomock.Any(), gomock.Any()).Times(1)
			} else {
				c.EXPECT().GetPokemonFromApi(gomock.Any(), gomock.Any()).Times(1)
			}

			New(r, c)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, tc.endpoint, nil)

			r.ServeHTTP(recorder, request)
			assert.Equal(t, tc.status, recorder.Code)
			assert.Nil(t, err)
		})
	}
}
