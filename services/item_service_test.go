package services_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/osalomon89/go-testing/mocks"
	"github.com/osalomon89/go-testing/services"
	"github.com/stretchr/testify/assert"
)

func Test_itemService_CreateItem(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		name      string
		stock     int
		repoError error
		repoTimes int
	}

	tests := []struct {
		name      string
		args      args
		wantedErr error
	}{
		{
			name:      "Should work correctly",
			wantedErr: nil,
			args: args{
				name:      "tablet",
				stock:     10,
				repoError: nil,
				repoTimes: 1,
			},
		},
		{
			name:      "Should return error when item name is empty",
			wantedErr: fmt.Errorf("item name could not be empty"),
			args: args{
				name:      "",
				stock:     10,
				repoError: nil,
				repoTimes: 0,
			},
		},
		{
			name:      "Should return error when item stock is zero",
			wantedErr: fmt.Errorf("stock could not be zero"),
			args: args{
				name:      "tablet",
				stock:     0,
				repoError: nil,
				repoTimes: 0,
			},
		},
		{
			name:      "Should return error when repository returns an error",
			wantedErr: fmt.Errorf("error in repository: %w", errors.New("the repository error")),
			args: args{
				name:      "tablet",
				stock:     10,
				repoError: errors.New("the repository error"),
				repoTimes: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			repositoryMock := mocks.NewMockItemRepository(ctrl)

			repositoryMock.EXPECT().
				SaveItem(tt.args.name, tt.args.stock).
				Return(tt.args.repoError).
				Times(tt.args.repoTimes)

			svc := services.NewItemService(repositoryMock)

			err := svc.CreateItem(tt.args.name, tt.args.stock)
			if tt.wantedErr != nil {
				assert.NotNil(err)
				assert.Equal(tt.wantedErr.Error(), err.Error(), "Error message is not the expected")
				return
			}

			assert.Nil(err)
		})
	}
}

func Test_itemService_GetItemByID(t *testing.T) { //Casos de prueba
	assert := assert.New()
	tabla := []struct {
		name      string
		id        uint
		wantedErr error
		repoTimes int
		repoError error
	}{
		{
			name:      "Esto no Funciona",
			id:        0,
			wantedErr: fmt.Errorf("item id cannot be zero"),
			repoTimes: 0,
			repoError: nil,
		},
		{
			name:      "Funciona",
			id:        1,
			wantedErr: nil,
			repoTimes: 0,
			repoError: nil,
		},
		{
			name:      "no encuentra el item",
			id:        100,
			wantedErr: fmt.Errorf("error in repository"),
			repoTimes: 1,
			repoError: errors.New("Error en el repositorio"),
		},
	}
	for _, tt := range tabla { //Logica de negocio
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish() //defer -> si algo que no controlaste me explota en el medio el finish se va ejecutar
			repositoryMock := mocks.NewMockItemRepository(ctrl)

			//Hacemos un mock
			repositoryMock.EXPECT().GetItemByID(tt.id).Return(nil).Times(tt.repoTimes)

			err := services.NewItemService(repositoryMock).GetItemByID(tt.id)

			if tt.wantedErr != nil {
				assert.NotNil(err)
				assert.Equal(tt.wantedErr.Error(), err.Error(), "Error message is not the expected")
				return
			}
			assert.Nil(err)

		})
	}

}
