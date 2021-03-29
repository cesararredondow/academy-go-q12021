package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cesararredondow/course/third_deriverable/routes/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

//mockgen -source=third_deriverable/routes/route.go -destination=third_deriverable/routes/mocks/route.go -package=mocks
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
			endpoint:       "/api/final/",
			handlerName:    "GetPokemonsFromApi",
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

			if tc.callController {
				c.EXPECT().GetPokemonsConcurrecny(gomock.Any(), gomock.Any()).Times(1)
			}

			New(c, r)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, tc.endpoint, nil)

			r.ServeHTTP(recorder, request)
			assert.Equal(t, tc.status, recorder.Code)
			assert.Nil(t, err)
		})
	}
}
