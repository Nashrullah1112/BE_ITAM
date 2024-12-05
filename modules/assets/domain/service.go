package domain

import (
	"github.com/banggibima/be-itam/pkg/config"
)

type AssetService struct {
	Config                  *config.Config
	PostgresAssetRepository PostgresAssetRepository
}

func NewAssetService(
	config *config.Config,
	postgresAssetRepository PostgresAssetRepository,
) *AssetService {
	return &AssetService{
		Config:                  config,
		PostgresAssetRepository: postgresAssetRepository,
	}
}

func (s *AssetService) CountAll() (int, error) {
	return s.PostgresAssetRepository.CountAll()
}

func (s *AssetService) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Asset, error) {
	return s.PostgresAssetRepository.FindAll(offset, limit, sort, order, filters)
}

func (s *AssetService) FindByID(id int) (*Asset, error) {
	return s.PostgresAssetRepository.FindByID(id)
}

func (s *AssetService) Create(asset *Asset) error {
	return s.PostgresAssetRepository.Create(asset)
}

func (s *AssetService) Update(asset *Asset) error {
	return s.PostgresAssetRepository.Update(asset)
}

func (s *AssetService) UpdatePartial(asset *Asset) error {
	return s.PostgresAssetRepository.UpdatePartial(asset)
}

func (s *AssetService) Delete(asset *Asset) error {
	return s.PostgresAssetRepository.Delete(asset)
}
