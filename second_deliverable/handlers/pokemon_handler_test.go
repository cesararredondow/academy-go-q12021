package handlers

import (
	"testing"

	"github.com/cesararredondow/course/second_deliverable/handlers/mocks"
	"github.com/cesararredondow/course/second_deliverable/models"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/unrolled/render"
)

//mockgen -source=second_deliverable/handlers/pokemon_handeler.go -destination=second_deliverable/handlers/mocks/pokemon_handeler.go -package=mocks

func Test_Controller(t *testing.T) {
	l := &logrus.Logger{}
	r := render.New()

	tests := []struct {
		name                    string
		expectedParams          string
		expectedUsecaseResponse []*models.Pokemon
		expectUsecaseCall       bool
		expectedError           error
		wantError               bool
	}{
		{
			name:           "OK, GetPokemons",
			expectedParams: "2",
			expectedUsecaseResponse: []*models.Pokemon{
				{
					ID:   1,
					Name: "bulbasaur",
					URL:  "https://pokeapi.co/api/v2/pokemon/1/",
				},
				{
					ID:   2,
					Name: "ivysaur",
					URL:  "https://pokeapi.co/api/v2/pokemon/2/",
				},
			},
			expectUsecaseCall: true,
			wantError:         false,
			expectedError:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			u := mocks.NewMockUseCase(mockCtrl)

			if tt.expectUsecaseCall {
				u.EXPECT().GetPokemons(tt.expectedParams).Return(tt.expectedUsecaseResponse, tt.expectedError)
			}

			c := New(u, l, r)

			response, err := c.useCase.GetPokemons(tt.expectedParams)
			assert.Equal(t, response, tt.expectedUsecaseResponse)

			// Pro gamer tip: Use wantError to check if the error should be nil
			if tt.wantError {
				assert.NotNil(t, err)
				assert.Equal(t, err, tt.expectedError)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
