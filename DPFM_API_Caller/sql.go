package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-country-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-country-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) Country(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.Country {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE country.Country = \"%s\ ", input.Country.Country),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	country.Country
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_country_country_data as country 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToCountry(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
