package skydio

// BatteriesService handles communication with the battery related
// methods of the Skydio API.
//
// Batteries are identified by their unique serial.
// High level health metrics are associated with each battery, including
// minimum and maximum voltages across all cells, minimum and maximum
// temperatures across all cells, and the total flight count and flight time
// for the battery. Batteries must be claimed into an organization in order
// to be accessed at /batteries.
//
// Read the support documentation at: https://support.skydio.com/hc/en-us/articles/7621964659483
type BatteriesService service
