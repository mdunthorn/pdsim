package lane

type Lane struct {
    Speed, Volume, Occupancy int
}

type Detector struct {
    Name        string
    Lanes       [8]Lane
    Voltage     int
    SensorId    int
    FirmwareRev int
    Health      int
}

