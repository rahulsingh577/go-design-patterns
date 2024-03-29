package creational

// A Builder design pattern tries to:
// Abstract complex creations so that object creation is separated from the object
// user
// Create an object step by step by filling its fields and creating the embedded
// objects
// Reuse the object creation algorithm between many objects

type ManufacturingDirector struct {
	builder BuildProcess
}

func (dir *ManufacturingDirector) Construct() {
	dir.builder.SetStructure().SetWheels().SetSeats()
}

func (dir *ManufacturingDirector) SetBuilder(b BuildProcess) {
	dir.builder = b
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}
func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}
func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}
func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}
