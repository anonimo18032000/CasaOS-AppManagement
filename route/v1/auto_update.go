package v1

import (
	"net/http"

	"github.com/IceWhaleTech/CasaOS-AppManagement/pkg/config"
	modelCommon "github.com/IceWhaleTech/CasaOS-Common/model"
	"github.com/IceWhaleTech/CasaOS-Common/utils/common_err"
	"github.com/labstack/echo/v4"
)

type AutoUpdateSettingRequest struct {
	Enabled bool `json:"enabled"`
}

// GetAutoUpdateSetting reports whether installed apps are automatically updated when a newer
// version is available in the app store (checked hourly, see the cron job in main.go).
func GetAutoUpdateSetting(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, modelCommon.Result{
		Success: common_err.SUCCESS,
		Message: common_err.GetMsg(common_err.SUCCESS),
		Data:    map[string]bool{"enabled": config.AppInfo.AutoUpdateEnabled},
	})
}

func SetAutoUpdateSetting(ctx echo.Context) error {
	var request AutoUpdateSettingRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, modelCommon.Result{Success: common_err.CLIENT_ERROR, Message: err.Error()})
	}

	config.AppInfo.AutoUpdateEnabled = request.Enabled
	if err := config.SaveSetup(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, modelCommon.Result{Success: common_err.SERVICE_ERROR, Message: err.Error()})
	}

	return ctx.JSON(http.StatusOK, modelCommon.Result{Success: common_err.SUCCESS, Message: common_err.GetMsg(common_err.SUCCESS)})
}
