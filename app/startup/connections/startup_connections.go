package connections

import "wrench/app/manifest/application_settings"

func LoadConnections() error {
	app := application_settings.ApplicationSettingsStatic

	if app.Connections == nil {
		return nil
	}

	err := loadConnectionNats(app.Connections.Nats)

	return err
}
