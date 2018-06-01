package initial

import "github.com/rubinliudongpo/airad/utils"

func init() {
	utils.InitSql()
	utils.InitTemplate()
	utils.InitCache()
	utils.InitBootStrap()
}
