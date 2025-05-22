package skydio

// DocksService handles communication with the dock related
// methods of the Skydio API.
//
// Every dock claimed within your organization can be identified by the dock
// WiFi/UAV Name (eg Skydio2P-xxxx). In our API, we refer to this as the
// dock_serial.
// Docks have a dock_type, dock_name, and can be associated with a Vehicle.
// The location of a dock is available when the dock's takeoff point has been
// configured within a site.
// Docks connected to Skydio Cloud will also return information about their
// current state.
type DocksService service
