// Code generated from Pkl module `appconfig`. DO NOT EDIT.
package appconfig

type Server struct {
	// The hostname of this application.
	Host string `pkl:"host"`

	// The port to listen on.
	Port uint16 `pkl:"port"`

	// Medsenger Agent secret key.
	MedsengerAgentKey string `pkl:"medsengerAgentKey"`

	// Sets server to debug mode.
	Debug bool `pkl:"debug"`

	Ctg *Ctg `pkl:"ctg"`
}
