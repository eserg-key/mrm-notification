package mrm_notification

type ClientConfig interface {
	getProjectTopic() string
	getNotificationTopic() string
}

type DevClientConfig struct{}

func (d *DevClientConfig) getProjectTopic() string {
	return "dev-projects"
}
func (d *DevClientConfig) getNotificationTopic() string {
	return "dev-notifications"
}

type ProdClientConfig struct{}

func (p *ProdClientConfig) getProjectTopic() string {
	return "projects"
}
func (p *ProdClientConfig) getNotificationTopic() string {
	return "notifications"
}
