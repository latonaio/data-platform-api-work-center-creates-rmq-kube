package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-work-center-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToGeneralUpdates(general dpfm_api_input_reader.General) *GeneralUpdates {
	data := general

	return &GeneralUpdates{
			BusinessPartner:      data.BusinessPartner,
			Plant:                data.Plant,
	}
}

func ConvertToProductionCapacityUpdates(general dpfm_api_input_reader.General, productionCapacity dpfm_api_input_reader.ProductionCapacity) *ProductionCapacityUpdates {
	dataGeneral := general
	data := productionCapacity

	return &ProductionCapacity{
		    BusinessPartner:              dataGeneral.BusinessPartner,
		    Plant:                        dataGeneral.Plant,
	}
}
